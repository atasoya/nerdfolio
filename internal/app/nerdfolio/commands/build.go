package commands

import (
	"flag"
	"fmt"
	"os"
)

var BuildCmd = flag.NewFlagSet("build", flag.ExitOnError)
var BuildColorScheme = BuildCmd.String("color", "tokyonight", "Color Scheme")

func HandleBuild() {
	BuildCmd.Parse(os.Args[2:])
	fmt.Println("subcommand 'build'")
	fmt.Println("  color-scheme:", *BuildColorScheme)
}
