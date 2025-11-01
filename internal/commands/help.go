package commands

import (
	"flag"
	"fmt"
)

var HelpCmd = flag.NewFlagSet("help", flag.ExitOnError)

func HandleHelpCommand() {
	ShowHelp()
}

func ShowHelp() {
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Reset = "\033[0m"

	fmt.Println(Blue+"\nnerdfolio"+Reset, "1.0")

	fmt.Println(Blue+"USAGE:", Reset)
	fmt.Println("	nerdfolio [OPTIONS]\n")

	fmt.Println(Blue+"OPTIONS:", Reset)
	fmt.Println(Magenta, "	build <COLOR_SCHEME>", Reset)
	fmt.Println("		Build your source code")

	fmt.Println(Magenta, "	create", Reset)
	fmt.Println("		Create initial nerdfolio project scaffolding")

	fmt.Println(Magenta, "	open", Reset)
	fmt.Println("		Open out/index.html in browser")

	fmt.Println(Magenta, "	help", Reset)
	fmt.Println("		Print help information")
}
