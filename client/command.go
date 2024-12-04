package client

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"

	"github.com/peterhellberg/neocities/api"
)

// CmdRunner is the default command runner
var CmdRunner = NewRunner()

// Command contains a function to run,
// a flagset, and usage instructions
type Command struct {
	Run  func(cmd *Command, args *Args)
	Flag flag.FlagSet

	Key   string
	Usage string
	Short string
	Long  string
}

// Name returns the name of the command.
// It falls back on the first word in Usage
// if no Name was provided.
func (c *Command) Name() string {
	if c.Key != "" {
		return c.Key
	}

	return strings.Split(c.Usage, " ")[0]
}

func (c *Command) parseArguments(args *Args) (err error) {
	c.Flag.SetInterspersed(true)
	c.Flag.Usage = c.PrintUsage

	if err = c.Flag.Parse(args.Params); err == nil {
		args.Params = c.Flag.Args()
	}

	return
}

// Call calls the command with provided args
func (c *Command) Call(args *Args) (err error) {
	err = c.parseArguments(args)
	if err != nil {
		return
	}

	c.Run(c, args)

	return
}

// PrintUsage prints usage instructions for the command
func (c *Command) PrintUsage() {
	if c.Runnable() {
		fmt.Printf("usage: %s\n\n", c.FormattedUsage())
	}

	fmt.Println(strings.Trim(c.Long, "\n"))
}

// FormattedUsage returns the formatted usage for runnable command
func (c *Command) FormattedUsage() string {
	return fmt.Sprintf("neocities %s", c.Usage)
}

// Runnable checks if the command has a Run function
func (c *Command) Runnable() bool {
	return c.Run != nil
}

func getCredentials() api.Credentials {
	if key := os.Getenv("NEOCITIES_KEY"); key != "" {
		return api.Credentials{Key: key}
	}

	user := os.Getenv("NEOCITIES_USER")
	pass := os.Getenv("NEOCITIES_PASS")

	if user == "" {
		fmt.Println("Error: Missing environment variable NEOCITIES_USER or NEOCITIES_KEY")

		os.Exit(0)
	}

	if pass == "" {
		fmt.Println("Error: Missing environment variable NEOCITIES_PASS")

		os.Exit(0)
	}

	return api.Credentials{
		User: user,
		Pass: pass,
	}
}
