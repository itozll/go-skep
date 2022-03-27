package flag

import (
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/spf13/pflag"
)

var (
	Verbose  = newFlag((*pflag.FlagSet).BoolVarP, "verbose", "V", false, "add more details to output logging")
	SkipGit  = newFlag((*pflag.FlagSet).BoolVarP, "skip-git", "", false, "do not initialize a git repository")
	Group    = newFlag((*pflag.FlagSet).StringVarP, "group", "g", initd.DefaultGroup, "group name")
	Go       = newFlag((*pflag.FlagSet).StringVarP, "go", "", initd.DefaultGoVersion, "the golang version used by the project.")
	JSONData = newFlag((*pflag.FlagSet).StringVarP, "json", "", "", "customize project with json.")
	File     = newFlag((*pflag.FlagSet).StringVarP, "file", "f", "", "customize project with file.")
	FileType = newFlag((*pflag.FlagSet).StringVarP, "file-type", "", "yaml", "file type, support json/yaml")

	Parent = newFlag((*pflag.FlagSet).StringVarP, "parent", "p", "root", `variable name of parent command for this command`)
)

type flag[T any] struct {
	v *T
	f *pflag.Flag
}

type handlerNewFlag[T any] func(set *pflag.FlagSet, p *T, name, shorthand string, value T, usage string)

func newFlag[T any](pf handlerNewFlag[T], name, shorthand string, value T, usage string) *flag[T] {
	_f := &flag[T]{
		v: new(T),
	}

	_flag := pflag.NewFlagSet("__hidden__", pflag.ExitOnError)
	pf(_flag, _f.v, name, shorthand, value, usage)
	_f.f = _flag.Lookup(name)

	return _f
}

func (f *flag[T]) Value() T          { return *f.v }
func (f *flag[T]) Flag() *pflag.Flag { return f.f }
func (f *flag[T]) Changed() bool     { return f.f.Changed }
func (f *flag[T]) IfChanged(v *T) {
	if f.f.Changed {
		*v = *f.v
	}
}
