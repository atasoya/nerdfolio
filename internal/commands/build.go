package commands

import (
	"flag"
	"fmt"
	"nerdfolio/internal/colors"
	"nerdfolio/internal/helpers"
	"os"
)

var BuildCmd = flag.NewFlagSet("build", flag.ExitOnError)
var BuildColorScheme = BuildCmd.String("colorScheme", "catppuccinMocha", "Color Scheme")

func HandleBuildCommand() {
	BuildCmd.Parse(os.Args[2:])
	fmt.Println("Starting the build process!")
	fmt.Println("  colorScheme:", *BuildColorScheme)

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

	var colorScheme []byte
	switch *BuildColorScheme {
	case "catppuccinMocha":
		colorScheme = []byte(colors.CatppuccinMocha)
	case "catppuccinLatte":
		colorScheme = []byte(colors.CatppuccinLatte)
	case "catppuccinFrappe":
		colorScheme = []byte(colors.CatppuccinFrappe)
	case "catppccinMacchiato":
		colorScheme = []byte(colors.CatppuccinMacchiato)
	case "tokyoNightStorm":
		colorScheme = []byte(colors.TokyoNightStorm)
	case "gruvboxDark":
		colorScheme = []byte(colors.GruvboxDark)
	default:
		colorScheme = []byte("none")

	}
	err := os.WriteFile(currentPath+"/out/nerdfolio.css", colorScheme, 0644)
	if err != nil {
		fmt.Println("There was a problem importing nerdfolio.css")
		os.Exit(1)
	}
}
