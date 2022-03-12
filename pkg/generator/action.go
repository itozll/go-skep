package generator

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type Action struct {
	Before func() error
	After  func() error

	P tmpl.Provider

	Binder map[string]interface{}

	To       string
	Template []string
	Copy     []string

	Script []Script

	Actions []*Action
}

// run sed ...
type Script struct {
	File  string
	Flag  string
	After bool
}

func (ac *Action) Exec(path string, p tmpl.Provider, binder map[string]interface{}) error {
	if ac.Before != nil {
		if err := ac.Before(); err != nil {
			return err
		}
	}

	if ac.P != nil {
		ac.P = p
	}

	if ac.P == nil {
		fmt.Fprintf(os.Stderr, "need a template provider\n")
		os.Exit(1)
	}

	if len(ac.To) > 0 {
		path += ac.To + "/"
	}
	if len(ac.Template) > 0 {
		if ac.Binder == nil {
			ac.Binder = binder
		} else {
			for key, val := range binder {
				if _, ok := ac.Binder[key]; !ok {
					ac.Binder[key] = val
				}
			}
		}

		if err := ac.parseAndCopy(path, ac.Template, true); err != nil {
			return err
		}
	}

	if len(ac.Copy) > 0 {
		if err := ac.parseAndCopy(path, ac.Copy, false); err != nil {
			return err
		}
	}

	for _, v := range ac.Actions {
		if err := v.Exec(path, ac.P, ac.Binder); err != nil {
			return err
		}
	}

	if ac.After != nil {
		return ac.After()
	}

	return nil
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

		data := ac.P.ReadFile(srcName)
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
