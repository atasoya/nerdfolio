package commands

import (
	"flag"
	"fmt"
	"os"
)

var HelpCmd = flag.NewFlagSet("help", flag.ExitOnError)

func HandleHelpCommand() {
	HelpCmd.Parse(os.Args[2:])
	fmt.Println("help command")
}
