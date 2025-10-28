package commands

import (
	"flag"
	"fmt"
	"nerdfolio/internal/flags"
	"nerdfolio/internal/helpers"
	"os"
	"time"
)

var BuildCmd = flag.NewFlagSet("build", flag.ExitOnError)
var BuildColorSchemeFlag = BuildCmd.String("colorScheme", "catppuccinMocha", "Color Scheme")

func HandleBuildCommand() {
	start := time.Now()
	BuildCmd.Parse(os.Args[2:])
	fmt.Println("Starting the build process!")

	currentPath, _ := os.Getwd()
	outputDirectory := currentPath + "/out"

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		err := os.Mkdir("out", 0755)
		if err != nil {
			fmt.Println("There was a problem creating /out directry")
			os.Exit(1)
		}
	}

	helpers.Copy(currentPath+"/index.html", outputDirectory+"/index.html")

	flags.HandleColorSchemeFlag(BuildColorSchemeFlag, currentPath)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Build completed in:", elapsed.Milliseconds(), "ms")

}
