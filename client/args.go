package client

// Args contains a command and its params
type Args struct {
	Command string
	Params  []string
}

// newArgs creates a new Args using a list of strings
func newArgs(args []string) *Args {
	var (
		command string
		params  []string
	)

	if len(args) == 0 {
		params = []string{}
	} else {
		command = args[0]
		params = args[1:]
	}

	return &Args{
		Command: command,
		Params:  params,
	}
}

// IsParamsEmpty checks if the params size is 0
func (a *Args) IsParamsEmpty() bool {
	return a.ParamsSize() == 0
}

// ParamsSize returns the number of params
func (a *Args) ParamsSize() int {
	return len(a.Params)
}

// FirstParam returns the first param.
// It returns empty string if there are no params.
func (a *Args) FirstParam() string {
	if a.ParamsSize() == 0 {
		return ""
	}

	return a.Params[0]
}
