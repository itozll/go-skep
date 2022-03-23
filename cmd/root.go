/*
Copyright © 2022 Joiky

*/
package cmd

import (
	"os"
	"path/filepath"

	_ "github.com/itozll/go-skep/internal/template"
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/spf13/cobra"
)

var appName = filepath.Base(os.Args[0])

var rootCmd = &cobra.Command{
	Use:     appName,
	Short:   "scaffolding for go projects",
	Version: initd.Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&initd.Verbose, "verbose", "V", false, "add more details to output logging")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
