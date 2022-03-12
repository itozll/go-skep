package rtstatus

import (
	"fmt"
	"os"
)

func ExitIfError(err error) {
	if err == nil {
		return
	}

	Fatal(err.Error())
}

func Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, " ** \033[31;1mError\033[0m: "+format+"\n", args...)
}

func Info(note string, format string, args ...interface{}) {
	fmt.Printf(" \033[32m"+note+"\033[0m  "+format+"\n", args...)
}

func Fatal(format string, args ...interface{}) {
	Error(format, args...)
	os.Exit(1)
}
