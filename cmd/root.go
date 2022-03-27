/*
Copyright © 2022 Joiky

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"

	_ "github.com/itozll/go-skep/internal/template"
	"github.com/itozll/go-skep/pkg/flag"
	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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

	// rootCmd.PersistentFlags().BoolVarP(&initd.Verbose, "verbose", "V", false, "add more details to output logging")
	rootCmd.PersistentFlags().AddFlag(flag.Verbose.Flag())

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parse(plain []byte, v map[string]any) []byte {
	var buff bytes.Buffer
	tpl, err := template.New("").Parse(string(plain))
	rtstatus.ExitIfError(err)

	err = tpl.Execute(&buff, v)
	rtstatus.ExitIfError(err)

	return buff.Bytes()
}

func loadConfig(p any, dfl []byte) {
	var dat []byte
	var fileType string

	switch {
	case flag.File.Value() != "":
		dat = process.ReadFile(flag.File.Value())
		fileType = flag.FileType.Value()

	case flag.JSONData.Value() != "":
		dat = []byte(flag.JSONData.Value())
		fileType = "json"

	default:
		dat = dfl
		fileType = "yaml"
	}

	unmarshal(dat, fileType, p)
}

func unmarshal(data []byte, typ string, dst interface{}) {
	var err error
	switch typ {
	default:
		panic("no support type")

	case "yaml", "yml":
		err = yaml.Unmarshal(data, dst)
	case "json":
		err = json.Unmarshal(data, dst)
	}

	rtstatus.ExitIfError(err)
}
