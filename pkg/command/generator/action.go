package generator

import (
	"errors"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/itozll/go-skep/pkg/command"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

type Action struct {
	Base

	Parse []string `json:"parse,omitempty" yaml:"parse"`
	Copy  []string `json:"copy,omitempty" yaml:"copy"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions"`
}

func (ac *Action) Exec(worker command.WorkerHandler) (err error) {
	ac.init()

	if ac.before != nil {
		if err = ac.before(); err != nil {
			return err
		}
	}

	if err = ac.Base.Exec(worker); err != nil {
		return err
	}

	if len(ac.Parse) > 0 {
		if err := ac.parseAndCopy(ac.Path, ac.Parse, true); err != nil {
			return err
		}
	}

	if len(ac.Copy) > 0 {
		if err := ac.parseAndCopy(ac.Path, ac.Copy, false); err != nil {
			return err
		}
	}

	for _, v := range ac.Actions {
		if err := v.Exec(ac); err != nil {
			return err
		}
	}

	if ac.after != nil {
		if err = ac.after(); err != nil {
			return err
		}
	}

	return
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