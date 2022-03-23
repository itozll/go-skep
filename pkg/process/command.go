package process

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/itozll/go-skep/pkg/runtime/initd"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

func Chdir(path string) {
	rtstatus.ExitIfError(os.Chdir(path))
}

func Run(a ...[]string) error {
	for _, cmd := range a {
		if err := Command(cmd)(); err != nil {
			return err
		}
	}

	return nil
}

func Command(a []string) func() error {
	if len(a) == 0 {
		return func() error { return nil }
	}

	var name string
	var arg []string

	name = a[0]
	if len(a) > 1 {
		arg = a[1:]
	}

	return func() (err error) {
		var buffer bytes.Buffer

		if initd.Verbose {
			for idx, arg := range a {
				if idx != 0 {
					buffer.WriteByte(' ')
				}

				if strings.Contains(arg, " ") {
					buffer.WriteByte('"')
					buffer.WriteString(arg)
					buffer.WriteByte('"')
				} else {
					buffer.WriteString(arg)
				}
			}

			rtstatus.Info("Run", "%s", buffer.String())
		}

		var stdout bytes.Buffer

		cmd := exec.Command(name, arg...)
		cmd.Stderr = &stdout
		if initd.Verbose {
			cmd.Stdout = &stdout
		}

		if err = cmd.Run(); err != nil {
			if !initd.Verbose {
				rtstatus.Info("Run", "%s", buffer.String())
			}

			rtstatus.Printf(os.Stdout, stdout.String())
			os.Exit(1)
		}

		if initd.Verbose {
			str := stdout.String()
			if len(str) > 0 {
				rtstatus.Printf(os.Stdout, stdout.String())
			}
		}

		return nil
	}
}
