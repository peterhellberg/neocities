package client

import (
	"os"

	"github.com/peterhellberg/neocities/api"
)

var cmdDelete = &Command{
	Run:   runDelete,
	Usage: "delete <filename> [<another filename>]",
	Short: "Delete files from Neocities",
	Long:  "Delete files from your Neocities website",
}

func init() {
	CmdRunner.Use(cmdDelete)
}

func runDelete(cmd *Command, args *Args) {
	if args.IsParamsEmpty() {
		cmd.PrintUsage()

		os.Exit(0)
	}

	cred := getCredentials()

	files := args.Params

	response, err := api.DeleteFiles(cred, files)
	if err != nil {
		response.Print()

		os.Exit(1)
	}

	if os.Getenv("NEOCITIES_VERBOSE") == "true" {
		response.Print()
	}

	os.Exit(0)
}
