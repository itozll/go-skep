/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/itozll/go-skep/pkg/etcd"
	"github.com/itozll/go-skep/pkg/model"
	"github.com/itozll/go-skep/pkg/runtime/rtinfo"
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
		var m *model.New

		switch {
		case rtinfo.File != "":
			m = model.NewNewWithFile(rtinfo.File, rtinfo.FileType)

		case rtinfo.Data != "":
			// --data '{}'
			m = model.NewNewWithJSON([]byte(rtinfo.Data))

		default:
			m = model.NewNew(etcd.NewEtc)
		}

		return m.Command().Exec()
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		rtinfo.Workspace = args[0]
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().AddFlagSet(rtinfo.CmdNewFlagSet)
}
