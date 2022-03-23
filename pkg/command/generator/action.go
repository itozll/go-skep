package generator

import (
	"errors"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Action struct {
	Before []string `json:"before,omitempty" yaml:"before"`
	After  []string `json:"after,omitempty" yaml:"after"`

	before func() error
	after  func() error

	Binder map[string]interface{} `json:"binder,omitempty" yaml:"binder"`
	Path   string                 `json:"path,omitempty" yaml:"path"`

	Template string `json:"template,omitempty" yaml:"template"`
	p        tmpl.Provider

	Parse []string `json:"parse,omitempty" yaml:"parse"`
	Copy  []string `json:"copy,omitempty" yaml:"copy"`

	Scripts [][]string `json:"scripts,omitempty" yaml:"scripts"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions"`
}

func (ac *Action) Dir() string                       { return ac.Path }
func (ac *Action) Provider() tmpl.Provider           { return ac.p }
func (ac *Action) MapBinder() map[string]interface{} { return ac.Binder }

func (ac *Action) init() {
	ac.before = process.Command(ac.Before)
	ac.after = process.Command(ac.After)
}

func (ac *Action) exec(worker command.WorkerHandler) error {
	if len(ac.Template) == 0 {
		ac.p = worker.Provider()
	} else {
		ac.p = tmpl.GetTemplateProvider(ac.Template)
	}

	if ac.Binder == nil {
		ac.Binder = worker.MapBinder()
	} else {
		for key, val := range worker.MapBinder() {
			if _, ok := ac.Binder[key]; !ok {
				ac.Binder[key] = val
			}
		}
	}

	if len(ac.Path) > 0 {
		ac.Path = worker.Dir() + ac.Path + "/"
	} else {
		ac.Path = worker.Dir()
	}

	return nil
}

func (ac *Action) Exec(worker command.WorkerHandler) (err error) {
	ac.init()

	if err = ac.before(); err != nil {
		return
	}

	if err = ac.exec(worker); err != nil {
		return
	}

	if len(ac.Parse) > 0 {
		if ac.p == nil {
			return errors.New("need the template provider")
		}

		if err = ac.parseAndCopy(ac.Path, ac.Parse, true); err != nil {
			return
		}
	}

	if len(ac.Copy) > 0 {
		if err = ac.parseAndCopy(ac.Path, ac.Copy, false); err != nil {
			return
		}
	}

	if len(ac.Scripts) > 0 {
		if err = process.Run(ac.Scripts...); err != nil {
			return
		}
	}

	for _, v := range ac.Actions {
		if err = v.Exec(ac); err != nil {
			return
		}
	}

	return ac.after()
}

func (ac *Action) parseAndCopy(path string, list []string, isTmpl bool) error {
	if len(path) > 0 {
		if err := os.MkdirAll(path, os.ModePerm); err != nil && os.IsNotExist(err) {
			rtstatus.Error("%s (%s)", err, path)
			return err
		}
	}

	for _, name := range list {
		if len(name) == 0 {
			return errors.New("source name must not be empty")
		}

		dstName, srcName := splitName(name)
		dstPath := path + dstName

		dstFd, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFd.Close()

		data := ac.p.ReadFile(srcName)
		if !isTmpl {
			_, err = io.Copy(dstFd, strings.NewReader(string(data)))
			if err != nil {
				return err
			}
		} else {
			tpl, err := template.New(dstPath).Parse(string(data))
			if err != nil {
				return err
			}

			err = tpl.Execute(dstFd, ac.Binder)
			if err != nil {
				return err
			}
		}

		rtstatus.Info("Create", dstPath)
	}

	return nil
}

func splitName(name string) (dst, src string) {
	names := strings.Split(name, ":")
	src = names[0]
	dst = src

	l := len(names)

	// name
	if l == 1 {
		return
	}

	if l > 2 {
		// name::new-name
		if len(names[2]) > 0 {
			dst = names[2]
			return
		}
	}

	// name:<suffix> -> model:go change to model.go
	if len(names[1]) > 0 {
		dst += names[1]
	}

	return
}
