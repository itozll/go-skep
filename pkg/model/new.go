package model

import (
	"os"

	"github.com/itozll/go-skep/pkg/generator"
	"github.com/itozll/go-skep/pkg/model/entity"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtinfo"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"gopkg.in/yaml.v2"
)

type New struct {
	eni *entity.New
}

func NewNew(eni *entity.New) *New {
	return &New{eni: eni}
}

func (m *New) Command() *generator.Command {
	rtinfo.Init(rtinfo.Workspace)
	rtinfo.Binder()["command"] = "server"

	command := generator.NewCommand(&m.eni.Resource)

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
		cmds = append(cmds, m.eni.After)

		delete(rtinfo.Binder(), "command")
		data, err := yaml.Marshal(rtinfo.Binder())
		if err == nil {
			os.WriteFile(".skeprc.yml", data, 0644)
		}

		return process.Run(cmds...)
	}
	return command
}
