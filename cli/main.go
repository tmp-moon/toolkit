package main

import (
	"os"

	"github.com/tmp-moon/toolkit/cli/cli"
)

func main() {
	err := cli.Execute()
	if err != nil {
		os.Exit(1)
	}
}
