package commands

import (
	"flag"
	"fmt"
	"nerdfolio/internal/flags"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

	createOutDirectoryIfNotExsists(outputDirectory)

	indexHtmlContentData, _ := os.ReadFile(currentPath + "/index.html")
	indexHtmlContent := string(indexHtmlContentData)

	indexHtmlContent = replaceTemplates(indexHtmlContent, currentPath)

	err := os.WriteFile(currentPath+"/out/index.html", []byte(indexHtmlContent), 0644)
	if err != nil {
		fmt.Println("error while writing index.html content")
		os.Exit(1)
	}

	flags.HandleColorSchemeFlag(BuildColorSchemeFlag, currentPath)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Build completed in:", elapsed.Milliseconds(), "ms")

}

func replaceTemplates(indexHtmlContent string, currentPath string) string {
	templateMap := make(map[string]string)
	files, err := os.ReadDir("templates/")
	if err != nil {
		fmt.Println("Error while reading templates")
		os.Exit(1)
	}

	for _, file := range files {
		fileName := file.Name()
		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		data, err := os.ReadFile(currentPath + "/templates/" + fileName)
		if err != nil {
			fmt.Println("Error while reading html files")
			os.Exit(1)
		}
		templateMap[baseName] = string(data)
		re := regexp.MustCompile(`\{%\s*` + baseName + `\s*%\}`)
		indexHtmlContent = re.ReplaceAllString(indexHtmlContent, templateMap[baseName])
	}
	return indexHtmlContent
}

func createOutDirectoryIfNotExsists(outputDirectory string) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		err := os.Mkdir("out", 0755)
		if err != nil {
			fmt.Println("There was a problem creating /out directry")
			os.Exit(1)
		}
	}
}
