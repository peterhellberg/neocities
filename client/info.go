package client

import (
	"fmt"
	"os"

	"github.com/peterhellberg/neocities/api"
)

var cmdInfo = &Command{
	Run:   runInfo,
	Usage: "info [sitename]",
	Short: "Info about Neocities websites",
	Long:  "Info about your Neocities website, or somebody elses",
}

func init() {
	CmdRunner.Use(cmdInfo)
}

func runInfo(cmd *Command, args *Args) {
	var (
		site string
		cred api.Credentials
	)

	if args.IsParamsEmpty() {
		cred = getCredentials()
		site = cred.User
	} else {
		site = args.FirstParam()
	}

	response, err := api.SiteInfo(cred, site)
	if err != nil {
		response.Print()

		os.Exit(1)
	}

	if len(response.Body) > 0 {
		fmt.Print(string(response.Body))
	}

	os.Exit(0)
}
