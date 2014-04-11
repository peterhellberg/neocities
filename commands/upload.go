package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/peterhellberg/neocities/api"
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
	check(err)

	files := args.Params
	response, err := api.UploadFiles(cred, files)
	check(err)

	if os.Getenv("NEOCITIES_VERBOSE") != "false" {
		response.Print()
	}

	os.Exit(0)
}

func getCredentials() (*api.Credentials, error) {
	user, err := getenv("NEOCITIES_USER")
	check(err)

	pass, err := getenv("NEOCITIES_PASS")
	check(err)

	return &api.Credentials{User: user, Pass: pass}, nil
}

func getenv(variable string) (string, error) {
	value := os.Getenv(variable)

	if value == "" {
		return value, errors.New("Missing environment variable " + variable)
	}

	return value, nil
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)

		os.Exit(1)
	}
}
