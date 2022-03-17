/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/itozll/go-skep/pkg/etcd"
	"github.com/itozll/go-skep/pkg/model"
	"github.com/itozll/go-skep/pkg/model/entity"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtinfo"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new [flags] <workspace>",
	Aliases: []string{"n"},
	Short:   "create an go workspace",
	Example: fmt.Sprintf(`  %s new --group mygroup myrepos
  %s new mygroup/myrepos
`, appName, appName),
	SilenceUsage: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		rtinfo.Workspace = args[0]
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		var entityNew *entity.New

		switch {
		case rtinfo.File != "":
			data := process.ReadFile(rtinfo.File)
			switch rtinfo.FileType {
			case "yaml":
				err := yaml.Unmarshal([]byte(data), entityNew)
				rtstatus.ExitIfError(err)
			case "json":
				err := json.Unmarshal([]byte(data), &entityNew)
				rtstatus.ExitIfError(err)
			}

		case rtinfo.Data != "":
			// --data '{}'
			err := json.Unmarshal([]byte(rtinfo.Data), &entityNew)
			rtstatus.ExitIfError(err)

		default:
			entityNew = etcd.NewEtc
		}

		entityNew.Parse()
		return model.NewNew(entityNew).Command().Exec()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().AddFlag(rtinfo.FlagGoVersion)
	newCmd.Flags().AddFlag(rtinfo.FlagSkipGit)
	newCmd.Flags().AddFlag(rtinfo.FlagGroup)
	newCmd.Flags().AddFlag(rtinfo.FlagJSONData)
	newCmd.Flags().AddFlag(rtinfo.FlagFile)
	newCmd.Flags().AddFlag(rtinfo.FlagFileType)
}
