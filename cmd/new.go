/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/itozll/go-skep/pkg/command/generator"
	"github.com/itozll/go-skep/pkg/etcd"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/initd"
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

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.Help()
			os.Exit(1)
		}
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		var entityNew *generator.New

		switch {
		case file != "":
			data := process.ReadFile(file)
			switch fileType {
			case "yaml":
				err := yaml.Unmarshal([]byte(data), entityNew)
				rtstatus.ExitIfError(err)
			case "json":
				err := json.Unmarshal([]byte(data), &entityNew)
				rtstatus.ExitIfError(err)
			}

		case data != "":
			// --data '{}'
			err := json.Unmarshal([]byte(data), &entityNew)
			rtstatus.ExitIfError(err)

		default:
			entityNew = etcd.NewEtc
		}

		return entityNew.Worker(cmd, args).Exec()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().String("group", initd.DefaultGroup, "group name.")
	newCmd.Flags().String("go", initd.DefaultGoVersion, "the golang version used by the project.")
	newCmd.Flags().Bool("skip-git", false, "do not initialize a git repository.")
	newCmd.Flags().StringVarP(&data, "json", "", "", "customize project with json.")
	newCmd.Flags().StringVarP(&file, "file", "f", "", "customize project with file.")
	newCmd.Flags().StringVarP(&fileType, "file-type", "", "yaml", "file type, support json/yaml")
	// newCmd.Flags().String("parent", "root", `variable name of parent command for this command`)
}
