package main

import (
	"github.com/itozll/go-skep/cmd"
)

//go:generate ./generator.sh

func main() {
	cmd.Execute()
}
