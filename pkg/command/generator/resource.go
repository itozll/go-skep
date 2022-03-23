package generator

import (
	"errors"

	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Resourcer interface {
	GetResource() *Resource
}

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

type Base struct {
	Before []string `json:"before,omitempty" yaml:"before"`
	After  []string `json:"after,omitempty" yaml:"after"`

	before func() error
	after  func() error

	Template string `json:"template,omitempty" yaml:"template"`
	p        tmpl.Provider

	Binder map[string]interface{} `json:"binder,omitempty" yaml:"binder"`
	Path   string                 `json:"path,omitempty" yaml:"path"`
}

func (b *Base) Dir() string                       { return b.Path }
func (b *Base) Provider() tmpl.Provider           { return b.p }
func (b *Base) MapBinder() map[string]interface{} { return b.Binder }

func (b *Base) init() {
	b.before = process.Command(b.Before)
	b.after = process.Command(b.After)
}

func (b *Base) Exec(worker command.WorkerHandler) error {
	if len(b.Template) == 0 {
		b.p = worker.Provider()
	} else {
		b.p = tmpl.GetTemplateProvider(b.Template)
	}

	if b.p == nil {
		return errors.New("need the template provider")
	}

	if b.Binder == nil {
		b.Binder = worker.MapBinder()
	} else {
		for key, val := range worker.MapBinder() {
			if _, ok := b.Binder[key]; !ok {
				b.Binder[key] = val
			}
		}
	}

	if len(b.Path) > 0 {
		b.Path = worker.Dir() + b.Path + "/"
	} else {
		b.Path = worker.Dir()
	}

	return nil
}
