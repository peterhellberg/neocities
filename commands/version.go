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
	Long:  "Show the version number of the Neocities API client",
}

func init() {
	CmdRunner.Use(cmdVersion)
}

func runVersion(cmd *Command, args *Args) {
	neocitiesVersion := fmt.Sprintf("neocities version %s", Version)

	fmt.Println(neocitiesVersion)

	os.Exit(0)
}
