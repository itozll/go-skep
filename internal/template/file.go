package template

import (
	"io/fs"
	"os"
	_ "unsafe"

	"github.com/itozll/go-skep/pkg/tmpl"
)

var _fs = FileSystem{}

type FileSystem struct{}

func (FileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (FileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

//go:linkname withFileSystem github.com/itozll/go-skep/pkg/tmpl.WithFileSystem
func withFileSystem(path string) tmpl.Provider {
	return newProvide(path, &_fs)
}
