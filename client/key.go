package client

import (
	"fmt"

	"github.com/peterhellberg/neocities/api"
)

var cmdKey = &Command{
	Run:   runKey,
	Usage: "key",
	Short: "Neocities API Key",
	Long:  "Retrieve an API Key for your Neocities user",
}

func init() {
	CmdRunner.Use(cmdKey)
}

func runKey(cmd *Command, args *Args) {
	cred := getCredentials()

	kr, err := api.Key(cred)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(kr.APIKey)
}
