package generator

import (
	"os"

	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type New struct {
	Resource `json:",inline" yaml:",inline"`

	Workspace string `json:"workspace,omitempty" yaml:"workspace"`

	SkipGit   bool   `json:"skip_git,omitempty" yaml:"skip_git"`
	Group     string `json:"group,omitempty" yaml:"group"`
	GoVersion string `json:"go_version,omitempty" yaml:"go_version"`
}

func (rc *New) Worker(cmd *cobra.Command, args []string) *command.Worker {
	rc.Parse(cmd, args)
	initd.Setup(rc.Workspace)

	binder := initd.Binder
	binder.Command = "server"

	worker := rc.Resource.Worker()
	worker.Path = binder.Project + "/"

	before := worker.Before
	worker.Before = func() error {
		_, err := os.Stat(binder.Project)
		if err == nil {
			rtstatus.Fatal("'" + binder.Project + "' exists.")
		}

		if !os.IsNotExist(err) {
			rtstatus.Fatal(err.Error())
		}

		return before()
	}

	worker.After = func() error {
		process.Chdir(binder.Project)

		cmds := [][]string{{"go", "mod", "tidy"}}
		if !rc.SkipGit {
			// 初始化 git 库
			cmds = append(cmds,
				[]string{"git", "init", "-b", "main"},
				[]string{"git", "remote", "add", "origin", "git@" + binder.Github + ".git"},
				[]string{"git", "add", "."},
				[]string{"git", "commit", "-m", "Init Commit"},
			)
		}
		cmds = append(cmds, rc.After)

		delete(worker.Binder, "command")
		data, err := yaml.Marshal(worker.Binder)
		if err != nil {
			return err
		}

		os.WriteFile(initd.ConfigFile, data, 0644)

		return process.Run(cmds...)
	}

	return worker
}

func (rc *New) Parse(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()

	if flags.Changed("group") {
		rc.Group, _ = cmd.Flags().GetString("group")
	}

	if flags.Changed("skip-git") {
		rc.SkipGit, _ = cmd.Flags().GetBool("skip-git")
	}

	if flags.Changed("go") {
		rc.GoVersion, _ = cmd.Flags().GetString("go")
	}

	if len(args) == 1 {
		rc.Workspace = args[0]
	}

	if len(rc.Group) != 0 {
		initd.Binder.Group = rc.Group
	}

	if len(rc.GoVersion) != 0 {
		initd.Binder.GoVersion = rc.GoVersion
	}
}
