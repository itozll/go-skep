package template

import (
	"embed"
	_ "unsafe"

	"github.com/itozll/go-skep/pkg/tmpl"
)

//go:embed templates/*
var f embed.FS

func builtin() tmpl.Provider {
	return newProvide("templates", f)
}

//go:linkname withFS github.com/itozll/go-skep/pkg/tmpl.WithFS
func withFS(path string, _f embed.FS) tmpl.Provider {
	return newProvide(path, _f)
}

func init() {
	tmpl.AddTemplateProvider("base", builtin())
}
