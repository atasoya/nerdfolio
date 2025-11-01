package flags

import (
	"fmt"
	"nerdfolio/internal/colors"
	"nerdfolio/internal/helpers"
	"os"
)

func HandleColorSchemeFlag(buildColorScheme *string, currentPath string) {

	var colorScheme []byte
	custom := false
	switch *buildColorScheme {
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
	case "custom":
		custom = true
	default:
		fmt.Println(*buildColorScheme, "is not a supported color scheme")
		os.Exit(1)
	}

	if custom {
		helpers.Copy(currentPath+"/custom.css", currentPath+"/out/custom.css")
		fmt.Println("  » Copied custom.css file to out/ directory")
	} else {
		err := os.WriteFile(currentPath+"/out/nerdfolio.css", colorScheme, 0644)
		if err != nil {
			fmt.Println("There was a problem importing nerdfolio.css")
			os.Exit(1)
		}
		fmt.Println("  » Copied nerdfolio.css file to out/ directory")

	}

}
