package commands

import (
	"flag"
	"fmt"
	"os"
)

var HelpCmd = flag.NewFlagSet("help", flag.ExitOnError)

func HandleHelpCommand() {
	ShowHelp()
}

func ShowHelp() {
	HelpCmd.Parse(os.Args[2:])

	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Reset = "\033[0m"

	fmt.Println(Blue+"\nnerdfolio"+Reset, "0.1")

	fmt.Println(Blue+"USAGE:", Reset)
	fmt.Println("	nerdfolio [OPTIONS]\n")

	fmt.Println(Blue+"OPTIONS:", Reset)
	fmt.Println(Magenta, "	build <COLOR_SCHEME>", Reset)
	fmt.Println("		Build your source code")

	fmt.Println(Magenta, "	help", Reset)
	fmt.Println("		Print help information")
}
