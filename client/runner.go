package client

import (
	"os"

	flag "github.com/ogier/pflag"
)

// Runner contains a map of commands
type Runner struct {
	commands map[string]*Command
}

// NewRunner creates a new Runner
func NewRunner() *Runner {
	return &Runner{
		commands: make(map[string]*Command),
	}
}

// All returns all commands
func (r *Runner) All() map[string]*Command {
	return r.commands
}

// Use adds the command to the runner
func (r *Runner) Use(command *Command) {
	r.commands[command.Name()] = command
}

// Lookup looks up a command with the given name
func (r *Runner) Lookup(name string) *Command {
	return r.commands[name]
}

// Execute executes a command
func (r *Runner) Execute() ExecError {
	args := newArgs(os.Args[1:])

	if args.Command == "" {
		printUsage()

		return newExecError(nil)
	}

	cmd := r.Lookup(args.Command)

	if cmd != nil && cmd.Runnable() {
		return r.Call(cmd, args)
	}

	return newExecError(nil)
}

// Call calls a command with the given args
func (r *Runner) Call(cmd *Command, args *Args) ExecError {
	err := cmd.Call(args)

	if err != nil {
		if err == flag.ErrHelp {
			err = nil
		}

		return newExecError(err)
	}

	return newExecError(nil)
}
