package rtinfo

import (
	"github.com/spf13/pflag"
)

var (
	FlagGoVersion *pflag.Flag
	FlagGroup     *pflag.Flag
	FlagSkipGit   *pflag.Flag
	FlagJSONData  *pflag.Flag
)

func init() {
	flagSet := pflag.NewFlagSet("__anonymous__", pflag.ExitOnError)

	flagSet.StringVarP(&Group, "group", "g", DefaultGroup, "group name.")
	flagSet.StringVarP(&GoVersion, "go-version", "", DefaultGoVersion, "the golang version used by the project.")
	flagSet.BoolVarP(&SkipGit, "skip-git", "", false, "do not initialize a git repository.")
	flagSet.StringVarP(&Data, "json", "", "", "customize project with json.")

	FlagGoVersion = flagSet.Lookup("go-version")
	FlagSkipGit = flagSet.Lookup("skip-git")
	FlagGroup = flagSet.Lookup("group")
	FlagJSONData = flagSet.Lookup("json")
}

func CmdNewFlagSet() *pflag.FlagSet {
	flagSet := pflag.NewFlagSet("new", pflag.ExitOnError)

	flagSet.AddFlag(FlagGoVersion)
	flagSet.AddFlag(FlagSkipGit)
	flagSet.AddFlag(FlagGroup)
	flagSet.AddFlag(FlagJSONData)

	return flagSet
}
