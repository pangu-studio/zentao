package main

import (
	"os"

	"github.com/awesome-skill/template/cmd/myskill/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
