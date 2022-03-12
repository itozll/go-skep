package template

import (
	"io/fs"
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
	data, err := p.f.ReadFile(p.path + name)
	rtstatus.ExitIfError(err)
	return data
}

func (p *provide) ReadDir(name string) []fs.DirEntry {
	de, err := p.f.ReadDir(p.path + name)
	rtstatus.ExitIfError(err)
	return de
}
