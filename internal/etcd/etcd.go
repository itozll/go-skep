package etcd

import (
	"embed"
	"io"
	"os"
	"strings"

	"github.com/itozll/go-skep/pkg/process"
	"github.com/itozll/go-skep/pkg/runtime/rtstatus"
)

//go:embed yaml/*.yaml
var f embed.FS

func Get(name string) []byte {
	data, err := f.ReadFile("yaml/" + name + ".yaml")
	rtstatus.ExitIfError(err)
	return data
}

func CopyAll(dstPath string, exclude ...string) {
	if dstPath == "" {
		rtstatus.Fatal("path must not be empty")
	}

	entries, err := f.ReadDir("yaml")
	rtstatus.ExitIfError(err)

	err = process.MkdirAll(dstPath)
	rtstatus.ExitIfError(err)

step1:
	for _, entry := range entries {
		name := entry.Name()

		for _, name2 := range exclude {
			if name2+".yaml" == name {
				continue step1
			}
		}

		fd, err := os.Create(dstPath + "/" + name)
		rtstatus.ExitIfError(err)
		defer fd.Close()

		data, err := f.ReadFile("yaml/" + name)
		rtstatus.ExitIfError(err)
		_, err = io.Copy(fd, strings.NewReader(string(data)))
		rtstatus.ExitIfError(err)
	}
}
