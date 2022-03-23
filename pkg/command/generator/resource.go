package generator

import (
	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Resource struct {
	Before []string `json:"before,omitempty" yaml:"before"`
	After  []string `json:"after,omitempty" yaml:"after"`

	Provider string                 `json:"provider,omitempty" yaml:"provider"`
	Binder   map[string]interface{} `json:"binder,omitempty" yaml:"binder"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions"`
}

func (rc *Resource) Worker() *command.Worker {
	c := &command.Worker{
		Binder: rc.Binder,
		P:      tmpl.GetTemplateProvider(rc.Provider),
		Before: process.Command(rc.Before),
		After:  process.Command(rc.After),
	}

	for _, ac := range rc.Actions {
		c.Handlers = append(c.Handlers, ac)
	}

	return c
}
