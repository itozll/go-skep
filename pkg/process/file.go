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
