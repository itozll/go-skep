/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/itozll/go-skep/pkg/etcd"
	"github.com/itozll/go-skep/pkg/runtime/rtinfo"
	"github.com/itozll/go-skep/pkg/translator"
	"github.com/spf13/cobra"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		var c *translator.CmdNew

		if rtinfo.Data == "" {
			c = translator.New(etcd.NewEtc)
		} else {
			// --data '{}'
			c = translator.New2(rtinfo.Data)
		}

		return c.Run()
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			rtinfo.Workspace = args[0]
		} else if rtinfo.Data == "" {
			cmd.Help()
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().AddFlagSet(rtinfo.CmdNewFlagSet())
}
