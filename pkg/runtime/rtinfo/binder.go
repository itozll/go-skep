package rtinfo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

var m = map[string]interface{}{
	"app_name":   filepath.Base(os.Args[0]),
	"parent_cmd": "root",
}

func Binder() map[string]interface{} { return m }

func Init(workspace string) map[string]interface{} {
	if len(workspace) == 0 {
		rtstatus.Fatal("repos_name must not be empty")
	}

	Workspace = workspace
	groupName, reposName := split(workspace)
	if len(groupName) == 0 {
		rtstatus.Fatal("group_name must not be empty (%s)", reposName)
	}

	m["group_name"], m["app_name"] = groupName, reposName

	m["snake_app_name"] = strcase.ToSnake(reposName)
	m["kebab_app_name"] = strcase.ToKebab(reposName)
	m["camel_app_name"] = strcase.ToCamel(reposName)

	Github = "github.com/" + groupName + "/" + reposName
	Project = reposName
	Group = groupName

	m["github"] = Github
	m["go_version"] = GoVersion
	return m
}

func split(reposName string) (string, string) {
	list := strings.Split(reposName, "/")
	if len(list) > 2 {
		rtstatus.Fatal("error repos_name: %s", reposName)
	}

	if len(list) == 1 {
		return Group, list[0]
	}

	return list[0], list[1]
}
