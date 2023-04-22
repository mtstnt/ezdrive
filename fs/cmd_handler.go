package fs

type commandHandlerFn func(args []string) error

var commandHandlers = map[string]commandHandlerFn{
	"cd":  cdCommandHandler,
	"ls":  lsCommandHandler,
	"get": getCommandHandler,
	"use": useCommandHandler,
}

// The command format is like this:
// [command] [arg1] [arg2] [arg3]
func processCommand(tokens []string) error {
	// No token, do nothing.
	if len(tokens) == 0 {
		return nil
	}

	command := tokens[0]
	args := tokens[1:]

	return commandHandlers[command](args)
}

func cdCommandHandler(args []string) error {
	return nil
}

func lsCommandHandler(args []string) error {
	return nil
}

func getCommandHandler(args []string) error {
	return nil
}

func useCommandHandler(args []string) error {
	return nil
}
