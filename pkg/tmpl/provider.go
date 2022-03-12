package tmpl

import (
	"embed"
	"io/fs"
)

type Provider interface {
	ReadFile(string) []byte
	ReadDir(string) []fs.DirEntry
}

func WithFS(pathPrefix string, _f embed.FS) Provider
func WithFileSystem(pathPrefix string) Provider
