package client

import (
	"os"

	"github.com/peterhellberg/neocities/api"
)

var cmdUpload = &Command{
	Run:   runUpload,
	Usage: "upload <filename> [<another filename>]",
	Short: "Upload files to Neocities",
	Long:  "Upload files to your Neocities website",
}

func init() {
	CmdRunner.Use(cmdUpload)
}

func runUpload(cmd *Command, args *Args) {
	if args.IsParamsEmpty() {
		cmd.PrintUsage()
		os.Exit(0)
	}

	cred, err := getCredentials()
	check(err)

	files := args.Params
	response, err := api.UploadFiles(cred, files)
	if err != nil {
		response.Print()
		os.Exit(1)
	}

	if os.Getenv("NEOCITIES_VERBOSE") == "true" {
		response.Print()
	}

	os.Exit(0)
}
