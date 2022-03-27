package generator

import (
	"strings"

	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Resource struct {
	Before []string `json:"before,omitempty" yaml:"before,omitempty"`
	After  []string `json:"after,omitempty" yaml:"after,omitempty"`

	Template     string                 `json:"template,omitempty" yaml:"template,omitempty"`
	TemplatePath string                 `json:"template_path,omitempty" yaml:"template_path,omitempty"`
	Binder       map[string]interface{} `json:"binder,omitempty" yaml:"binder,omitempty"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions,omitempty"`
}

func (rc *Resource) Append(ac *Action) {
	rc.Actions = append(rc.Actions, ac)
}

func (rc *Resource) Worker() *command.Worker {
	c := &command.Worker{
		Binder: rc.Binder,
		Before: process.Commands(rc.Before),
		After:  process.Commands(rc.After),
	}

	if len(rc.TemplatePath) > 0 {
		c.P = tmpl.WithFileSystem(strings.TrimRight(rc.TemplatePath, "/") + "/")
	} else {
		c.P = tmpl.GetTemplateProvider(rc.Template)
	}

	for _, ac := range rc.Actions {
		c.Handlers = append(c.Handlers, ac)
	}

	return c
}
