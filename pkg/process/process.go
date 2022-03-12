package process

import (
	"os"
	"os/exec"

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

	return func() error {
		cmd := exec.Command(name, arg...)
		cmd.Stderr = os.Stderr
		rtstatus.ExitIfError(cmd.Run())

		return nil
	}
}

func Commands(a [][]string) func() error {
	return func() error {
		for _, cmd := range a {
			if err := Command(cmd)(); err != nil {
				return err
			}
		}

		return nil
	}
}
