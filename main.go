/*

A Neocities API client written in Go

Installation

Just go get the command:

    go get -u github.com/peterhellberg/neocities

Usage

First you need to export some environment variables:

	export NEOCITIES_USER=<username>
	export NEOCITIES_PASS=<password>

Then you can run the command:

    neocities upload <filename> [<another filename>]

*/
package main

import (
	"os"

	"github.com/peterhellberg/neocities/commands"
)

func main() {
	err := commands.CmdRunner.Execute()

	os.Exit(err.ExitCode)
}
