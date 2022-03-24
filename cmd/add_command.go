/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/itozll/go-skep/pkg/command/generator"
	"github.com/itozll/go-skep/pkg/etcd"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:          "command",
	Aliases:      []string{"cmd", "c"},
	Short:        "add a command to application",
	SilenceUsage: true,

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		initd.Binder.Command = args[0]
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		var entityCommand *generator.Resource

		switch {
		case file != "":
			data := process.ReadFile(file)
			switch fileType {
			case "yaml":
				err := yaml.Unmarshal([]byte(data), entityCommand)
				rtstatus.ExitIfError(err)
			case "json":
				err := json.Unmarshal([]byte(data), &entityCommand)
				rtstatus.ExitIfError(err)
			}

		case data != "":
			// --data '{}'
			err := json.Unmarshal([]byte(data), &entityCommand)
			rtstatus.ExitIfError(err)

		default:
			entityCommand = etcd.CommandEtc
		}

		name := args[0]
		if initd.Binder.Parent != initd.DefaultCommand {
			name = initd.Binder.Parent + "_" + name
		}

		entityCommand.Append(&generator.Action{
			Path: "app/cmd",
			Parse: []string{
				"subcmd_go::" + name + ".go",
			},
		})

		return entityCommand.Worker().Exec()
	},
}

func init() {
	addCmd.AddCommand(commandCmd)
}
