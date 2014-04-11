package commands

import (
	"fmt"
	"os"
)

const Version = "0.0.1"

var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "Show neocities version",
	Long:  "Show the version number of the neocities client",
}

func init() {
	CmdRunner.Use(cmdVersion)
}

func runVersion(cmd *Command, args *Args) {
	fmt.Println("neocities version", Version)

	os.Exit(0)
}
