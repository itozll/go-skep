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

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.Help()
			os.Exit(1)
		}
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		var entityNew generator.New
		loadConfig(&entityNew, etcd.Get("new"))

		if flag.Local.Value() && !flag.SkipGit.Changed() {
			// ignore git command if --local is set
			flag.SkipGit.Set(true)
		}

		return entityNew.Worker(cmd, args).Exec()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().AddFlag(flag.SkipGit.Flag())
	newCmd.Flags().AddFlag(flag.Group.Flag())
	newCmd.Flags().AddFlag(flag.Go.Flag())
	newCmd.Flags().AddFlag(flag.JSONData.Flag())
	newCmd.Flags().AddFlag(flag.File.Flag())
	newCmd.Flags().AddFlag(flag.FileType.Flag())
	newCmd.Flags().AddFlag(flag.Local.Flag())
	newCmd.Flags().AddFlag(flag.IncludeNew.Flag())
}
