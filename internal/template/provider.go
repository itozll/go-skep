package template

import (
	"errors"
	"io/fs"
	"os"
	"strings"

	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/itozll/go-skep/pkg/tmpl"
)

type provider interface {
	ReadFile(string) ([]byte, error)
	ReadDir(string) ([]fs.DirEntry, error)
}

type provide struct {
	f    provider
	path string
}

func newProvide(path string, f provider) tmpl.Provider {
	return &provide{path: strings.TrimRight(path, "/") + "/", f: f}
}

func (p *provide) ReadFile(name string) []byte {
	data, err := p.get(p.path + name)
	rtstatus.ExitIfError(err)
	return data
}

func (p *provide) ReadDir(name string) []fs.DirEntry {
	de, err := p.f.ReadDir(p.path + name)
	rtstatus.ExitIfError(err)
	return de
}

func (p *provide) get(name string) ([]byte, error) {
	data, err := p.f.ReadFile(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// try <name>.tpl
			if !strings.HasSuffix(name, ".tpl") {
				return p.get(name + ".tpl")
			}
		}

		rtstatus.Fatal("\n  ** can not get template[%s]: %s.", name, err)
	}

	return data, nil
}
