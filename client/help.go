package client

import (
	"fmt"
	"os"
)

var cmdHelp = &Command{
	Run:   runHelp,
	Usage: "help [command]",
	Short: "Show help",
	Long:  "Show usage instructions for a command",
}

func init() {
	CmdRunner.Use(cmdHelp)
}

func runHelp(cmd *Command, args *Args) {
	if args.IsParamsEmpty() {
		printUsage()

		os.Exit(0)
	}

	for _, cmd := range CmdRunner.All() {
		if cmd.Name() == args.FirstParam() {
			cmd.PrintUsage()

			os.Exit(0)
		}
	}
}

var helpText = `usage: neocities <command> [<args>]

Commands:
   upload    Upload files to Neocities
   delete    Delete files from Neocities
   version   Show neocities client version

Help for a specific command:
   help [command]

Environment setup:

   export NEOCITIES_USER=<username>
   export NEOCITIES_PASS=<password>

`

func printUsage() {
	fmt.Print(helpText)
}
