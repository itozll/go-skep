package translator

import (
	"encoding/json"
	"os"

	"github.com/itozll/go-skep/pkg/etcd"
	"github.com/itozll/go-skep/pkg/generator"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtinfo"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

type CmdNew struct {
	etc *etcd.New
}

func New(etc *etcd.New) *CmdNew {
	return &CmdNew{etc: etc}
}

func New2(str string) *CmdNew {
	var etc = etcd.New{}
	err := json.Unmarshal([]byte(str), &etc)
	rtstatus.ExitIfError(err)

	if !rtinfo.FlagSkipGit.Changed && etc.SkipGit {
		rtinfo.SkipGit = true
	}

	if !rtinfo.FlagGroup.Changed && etc.Group != "" {
		rtinfo.Group = etc.Group
	}

	if !rtinfo.FlagGoVersion.Changed && etc.GoVersion != "" {
		rtinfo.GoVersion = etc.GoVersion
	}

	if rtinfo.Workspace == "" {
		rtinfo.Workspace = etc.Workspace
	}

	return New(&etc)
}

func (c *CmdNew) Run() error {
	return c.Translate().Exec()
}

func (c *CmdNew) Translate() *generator.Command {
	rtinfo.Init(rtinfo.Workspace)
	rtinfo.Binder()["command"] = "server"

	command := generator.New(&c.etc.Resource)

	command.Path = rtinfo.Project + "/"
	before := command.Before
	command.Before = func() error {
		_, err := os.Stat(rtinfo.Project)
		if err == nil {
			rtstatus.Fatal("'" + rtinfo.Project + "' exists.")
		}

		if !os.IsNotExist(err) {
			rtstatus.Fatal(err.Error())
		}

		return before()
	}

	command.After = func() error {
		process.Chdir(rtinfo.Project)

		cmds := [][]string{{"go", "mod", "tidy"}}
		if !rtinfo.SkipGit {
			// 初始化 git 库
			cmds = append(cmds,
				[]string{"git", "init", "-b", "main"},
				[]string{"git", "remote", "add", "origin", "git@" + rtinfo.Github + ".git"},
				[]string{"git", "add", "."},
				[]string{"git", "commit", "-m", "Init Commit"},
			)
		}
		cmds = append(cmds, c.etc.After)
		return process.Run(cmds...)
	}
	return command
}
