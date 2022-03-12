/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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
	Example: `  new --group mygroup myrepos
  new mygroup/myrepos
`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := translator.New(etcd.NewEtc)
		return c.Run()
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		rtinfo.Workspace = args[0]
		// rtinfo.Init(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&rtinfo.OGroup, "group", "g", "", "group name (default "+rtinfo.DefaultGroup+")")
	newCmd.Flags().StringVarP(&rtinfo.OGoVersion, "go-version", "", rtinfo.DefaultGoVersion, "the golang version used by the project.")
	newCmd.Flags().StringVarP(&rtinfo.OSkipGit, "skip-git", "", "false", "do not initialize a git repository.")
}
