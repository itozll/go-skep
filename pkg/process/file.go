package process

import (
	"os"

	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

func ReadFile(name string) []byte {
	_, err := os.Stat(name)
	rtstatus.ExitIfError(err)

	data, err := os.ReadFile(name)
	rtstatus.ExitIfError(err)

	return data
}

func MkdirAll(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil &&
		(os.IsNotExist(err) || os.IsPermission(err)) {
		rtstatus.Error("%s (%s)", err, path)
		return err
	}

	return nil
}
