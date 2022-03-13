package generator

import (
	"github.com/itozll/go-skep/pkg/model/entity"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtinfo"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Command struct {
	Before func() error
	After  func() error

	Path string

	P               tmpl.Provider
	SkipGitModified bool

	Binder map[string]interface{}

	Actions []*Action
}

func NewCommand(etc *entity.Resource) *Command {
	c := &Command{
		Binder: etc.Binder,
		P:      tmpl.GetTemplateProvider(etc.Provider),
		Before: process.Command(etc.Before),
		After:  process.Command(etc.After),
	}

	for _, action := range etc.Actions {
		c.Actions = append(c.Actions, &Action{
			Before:   process.Command(action.Before),
			After:    process.Command(action.After),
			Binder:   action.Binder,
			P:        tmpl.GetTemplateProvider(action.Provider),
			To:       action.To,
			Template: action.Template,
			Copy:     action.Copy,
		})
	}

	return c
}

func (cmd *Command) Exec() (err error) {
	if cmd.Binder == nil {
		cmd.Binder = rtinfo.Binder()
	} else {
		for key, val := range rtinfo.Binder() {
			if _, ok := cmd.Binder[key]; !ok {
				cmd.Binder[key] = val
			}
		}
	}

	if cmd.Before != nil {
		if err = cmd.Before(); err != nil {
			return
		}
	}

	for _, action := range cmd.Actions {
		rtstatus.ExitIfError(action.Exec(cmd.Path, cmd.P, cmd.Binder))
	}

	if cmd.After != nil {
		return cmd.After()
	}

	return nil
}
