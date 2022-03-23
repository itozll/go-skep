/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:          "add <command>",
	Short:        "add a command to application",
	SilenceUsage: true,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		data := process.ReadFile(initd.ConfigFile)
		return yaml.Unmarshal(data, initd.Binder)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().String("parent", "root", `variable name of parent command for this command`)

}
