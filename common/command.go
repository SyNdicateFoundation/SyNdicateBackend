package command

import (
	"SyNdicateBackend/common/logger"
	"bufio"
	"github.com/muesli/termenv"
	"os"
)

type Command struct {
	Name        string
	Description string
	Handler     func(t *termenv.Output)
}

var (
	scanner  = bufio.NewScanner(os.Stdin)
	Commands = map[string]Command{
		"clear": {
			Name:        "clear",
			Description: "Clear the screen",
			Handler: func(t *termenv.Output) {
				t.ClearScreen()
			},
		},
		"help": {
			Name:        "help",
			Description: "Show this help message",
			Handler: func(t *termenv.Output) {

			},
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the TELNET session",
			Handler: func(t *termenv.Output) {
				logger.Logger.Info("Bye!")
				os.Exit(1)
			},
		},
	}
)

func readInput() string {
	scanner.Scan()
	return scanner.Text()
}
func HandleCommand(commandName string, terminal *termenv.Output) {

	if command, ok := Commands[commandName]; ok {
		command.Handler(terminal)

	} else {
		logger.Logger.Error("Unknown Command")
	}
}

func Loop(terminal *termenv.Output) {
	for {
		HandleCommand(readInput(), terminal)
	}
}
