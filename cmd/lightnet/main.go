package main

import (
	"os"

	"github.com/oshothebig/lightnet/cmd/lightnet/cmd"
)

func main() {
	cmd.Execute(os.Stdin, os.Stdout, os.Stderr)
}
