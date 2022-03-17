package rtinfo

import (
	"github.com/spf13/pflag"
)

var (
	FlagGoVersion *pflag.Flag
	FlagGroup     *pflag.Flag
	FlagSkipGit   *pflag.Flag
	FlagJSONData  *pflag.Flag
	FlagFile      *pflag.Flag
	FlagFileType  *pflag.Flag
	FlagParent    *pflag.Flag
)

func init() {
	flagSet := pflag.NewFlagSet("__anonymous__", pflag.ExitOnError)

	flagSet.StringVarP(&Group, "group", "g", DefaultGroup, "group name.")
	flagSet.StringVarP(&Go, "go", "", DefaultGoVersion, "the golang version used by the project.")
	flagSet.BoolVarP(&SkipGit, "skip-git", "", false, "do not initialize a git repository.")
	flagSet.StringVarP(&Data, "json", "", "", "customize project with json.")
	flagSet.StringVarP(&File, "file", "f", "", "customize project with file.")
	flagSet.StringVarP(&FileType, "file-type", "", "yaml", "file type, support json/yaml")
	flagSet.StringVarP(&Parent, "parent", "p", "root", `variable name of parent command for this command`)

	FlagGoVersion = flagSet.Lookup("go")
	FlagSkipGit = flagSet.Lookup("skip-git")
	FlagGroup = flagSet.Lookup("group")
	FlagJSONData = flagSet.Lookup("json")
	FlagFile = flagSet.Lookup("file")
	FlagFileType = flagSet.Lookup("file-type")
	FlagParent = flagSet.Lookup("parent")
}
