package commands

import (
	"flag"
	"os"
)

type Credentials struct {
	User string
	Pass string
}

type Runner struct {
	commands map[string]*Command
}

func NewRunner() *Runner {
	return &Runner{commands: make(map[string]*Command)}
}

func (r *Runner) All() map[string]*Command {
	return r.commands
}

func (r *Runner) Use(command *Command) {
	r.commands[command.Name()] = command
}

func (r *Runner) Lookup(name string) *Command {
	return r.commands[name]
}

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
