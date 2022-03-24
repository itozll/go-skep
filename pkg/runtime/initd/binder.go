package initd

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

var (
	Binder = &bind{
		Generator: filepath.Base(os.Args[0]),
		Parent:    DefaultCommand,
		GoVersion: DefaultGoVersion,
	}

	mapBinder map[string]interface{}
	once      sync.Once
)

type bind struct {
	Generator string `json:"generator,omitempty" yaml:"generator"`
	AppName   string `json:"app_name,omitempty" yaml:"app_name"`
	Parent    string `json:"parent,omitempty" yaml:"parent"`

	Workspace string `json:"workspace,omitempty" yaml:"workspace"`

	// for add command <Command>
	Command string `json:"command,omitempty" yaml:"-"`

	// for add model <Model>
	Model string `json:"model,omitempty" yaml:"-"`

	Github       string `json:"github,omitempty" yaml:"github"`
	Group        string `json:"group,omitempty" yaml:"group"`
	Project      string `json:"project,omitempty" yaml:"project"`
	SnakeProject string `json:"snake_project,omitempty" yaml:"snake_project"`
	KebabProject string `json:"kebab_project,omitempty" yaml:"kebab_project"`
	CamelProject string `json:"camel_project,omitempty" yaml:"camel_project"`

	GoVersion string `json:"go_version,omitempty" yaml:"go_version"`
}

func MapBinder() map[string]interface{} {
	once.Do(func() {
		data, _ := json.Marshal(Binder)

		err := json.Unmarshal(data, &mapBinder)
		rtstatus.ExitIfError(err)
	})

	return mapBinder
}

func Setup(workspace string) error {
	return Binder.Setup(workspace)
}

func (b *bind) Setup(workspace string) error {
	if len(workspace) == 0 {
		return errors.New("repos name must not be empty")
	}

	fields := strings.Split(workspace, "/")

	if len(fields) > 2 {
		return errors.New("repos name cannot have more '/'")
	}

	var group string
	if len(fields) == 1 {
		Binder.Project = strings.TrimSpace(fields[0])
		group = DefaultGroup
	} else {
		Binder.Project = strings.TrimSpace(fields[1])
		group = strings.TrimSpace(fields[0])
	}

	if len(Binder.Group) == 0 {
		Binder.Group = group
	}

	if len(Binder.Group) == 0 {
		return errors.New("group name must not be empty")
	}

	Binder.SnakeProject = strcase.ToSnake(Binder.Project)
	Binder.KebabProject = strcase.ToKebab(Binder.Project)
	Binder.CamelProject = strcase.ToCamel(Binder.Project)
	Binder.AppName = Binder.Project
	Binder.Workspace = Binder.Group + "/" + Binder.Project

	Binder.Github = "github.com/" + Binder.Group + "/" + Binder.Project
	return nil
}
