package generator

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/flag"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Action struct {
	Before []string `json:"before,omitempty" yaml:"before,omitempty"`
	After  []string `json:"after,omitempty" yaml:"after,omitempty"`

	before func() error
	after  func() error

	Binder map[string]interface{} `json:"binder,omitempty" yaml:"binder,omitempty"`
	Path   string                 `json:"path,omitempty" yaml:"path,omitempty"`

	Template string `json:"template,omitempty" yaml:"template,omitempty"`
	p        tmpl.Provider

	Parse []string `json:"parse,omitempty" yaml:"parse,omitempty"`
	Copy  []string `json:"copy,omitempty" yaml:"copy,omitempty"`

	Scripts [][]string `json:"scripts,omitempty" yaml:"scripts,omitempty"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions,omitempty"`
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
		if err := process.MkdirAll(path); err != nil {
			return err
		}
	}

	for _, name := range list {
		if len(name) == 0 {
			return errors.New("source name must not be empty")
		}

		dstName, srcName := splitName(name)
		dstPath := path + dstName

		if strings.Contains(dstName, "/") {
			if err := process.MkdirAll(filepath.Dir(dstPath)); err != nil {
				return err
			}
		}

		dstFd, err := os.Create(dstPath)
		rtstatus.ExitIfError(err)
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

		if flag.Verbose.Value() {
			rtstatus.Info("Create", dstPath)
		}
	}

	return nil
}

func splitName(name string) (dst, src string) {
	names := strings.Split(name, ":")
	src = strings.TrimSpace(names[0])
	dst = src

	if len(names) > 1 {
		dst = strings.TrimSpace(names[1])
	}

	return
}
