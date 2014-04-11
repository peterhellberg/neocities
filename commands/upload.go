package commands

import (
	"os"

	"github.com/peterhellberg/neocities/api"
	"github.com/peterhellberg/neocities/utils"
)

var cmdUpload = &Command{
	Run:   runUpload,
	Usage: "upload <filename> [<another filename>]",
	Short: "Upload files to Neocities",
	Long:  "Upload files to your Neocities account",
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
	utils.Check(err)

	files := args.Params
	response, err := api.UploadFiles(cred, files)
	utils.Check(err)

	if os.Getenv("NEOCITIES_VERBOSE") != "false" {
		response.Print()
	}

	os.Exit(0)
}

func getCredentials() (*api.Credentials, error) {
	user := os.Getenv("NEOCITIES_USER")
	pass := os.Getenv("NEOCITIES_PASS")

	return &api.Credentials{User: user, Pass: pass}, nil
}
