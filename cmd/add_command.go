/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/itozll/go-skep/internal/etcd"
	"github.com/itozll/go-skep/pkg/command/generator"
	"github.com/itozll/go-skep/pkg/flag"
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:     "command",
	Aliases: []string{"cmd", "c"},
	Short:   "add a command to application",
	Example: fmt.Sprintf(`  %s add command test [--parent root]
  %s add cmd     test [--parent root]
  %s add c       test [--parent root]
`, appName, appName, appName),
	SilenceUsage: true,

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		initd.Binder.Command = args[0]
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		var entityCommand generator.Resource

		name := args[0]
		if initd.Binder.Parent != initd.DefaultCommand {
			name = initd.Binder.Parent + "_" + name
		}

		v := initd.MapBinder()
		v["file_name"] = name

		loadConfig(&entityCommand, parse(etcd.Get("command"), v))
		return entityCommand.Worker().Exec()
	},
}

func init() {
	addCmd.AddCommand(commandCmd)

	commandCmd.Flags().AddFlag(flag.JSONData.Flag())
	commandCmd.Flags().AddFlag(flag.File.Flag())
	commandCmd.Flags().AddFlag(flag.FileType.Flag())
}
