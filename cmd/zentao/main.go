package main

import (
	"os"

	"github.com/pangu-studio/zentao/cmd/zentao/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
