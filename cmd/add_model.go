/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:          "model",
	Aliases:      []string{"m"},
	Short:        "add a model to application",
	SilenceUsage: true,

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		initd.Binder.Model = args[0]
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
		log.Printf("%+v\n", initd.Binder)
	},
}

func init() {
	addCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
