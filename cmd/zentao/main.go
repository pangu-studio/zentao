package main

import (
	"os"

	"github.com/awesome-skill/zentao/cmd/zentao/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
