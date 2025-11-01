package main

import (
	"fmt"
	"nerdfolio/internal/commands"
	"os"
)

func main() {

	arguments := os.Args
	// If 'nerdfolio' called without any arguments
	if len(arguments) < 2 {
		fmt.Println("expected 'build' or 'help' subcommands")
		os.Exit(1)
	}

	// First index is path and other indexes are arguments
	switch arguments[1] {
	case "build":
		commands.HandleBuildCommand()

	case "open":
		commands.HandleOpenCommand()

	case "create":
		commands.HandleCreateCommand()

	case "help":
		commands.HandleHelpCommand()

	default:
		commands.HandleHelpCommand()
	}
}
