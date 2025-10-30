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

var currentPath string

func HandleBuildCommand() {
	start := time.Now()
	BuildCmd.Parse(os.Args[2:])
	fmt.Println("Starting the build process!")

	currentPath, _ = os.Getwd()
	outputDirectory := currentPath + "/out"

	createOutDirectoryIfNotExsists(outputDirectory)

	htmlFilesMap := createHtmlFilesMap()
	htmlFilesMap, err := replaceTemplates(htmlFilesMap, currentPath)
	if err != nil {
		fmt.Println("Error replacing templates:", err)
		os.Exit(1)
	}

	writeHtmlFilesToOutDirectory(htmlFilesMap)

	flags.HandleColorSchemeFlag(BuildColorSchemeFlag, currentPath)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Build completed in:", elapsed.Milliseconds(), "ms")

}

func replaceTemplates(htmlFilesMap map[string]string, currentPath string) (map[string]string, error) {
	templateDir := filepath.Join(currentPath, "templates")
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, fmt.Errorf("reading templates: %w", err)
	}

	templateMap := make(map[string]string)
	for _, file := range entries {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		data, err := os.ReadFile(filepath.Join(templateDir, fileName))
		if err != nil {
			return nil, fmt.Errorf("reading template %s: %w", fileName, err)
		}
		templateMap[baseName] = string(data)
	}

	for htmlFile, content := range htmlFilesMap {
		for name, tmpl := range templateMap {
			re := regexp.MustCompile(fmt.Sprintf(`\{%%\s*%s\s*%%\}`, regexp.QuoteMeta(name)))
			content = re.ReplaceAllString(content, tmpl)
		}
		htmlFilesMap[htmlFile] = content
	}

	return htmlFilesMap, nil
}

func createHtmlFilesMap() map[string]string {
	htmlFilesMap := make(map[string]string)
	entries, _ := os.ReadDir(currentPath)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()

		if !strings.HasSuffix(strings.ToLower(fileName), ".html") {
			continue
		}

		filePath := filepath.Join(currentPath, fileName)

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", fileName, err)
			os.Exit(1)
		}

		htmlFilesMap[fileName] = string(data)

	}

	return htmlFilesMap

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

func writeHtmlFilesToOutDirectory(htmlFilesMap map[string]string) {
	for htmlFile, htmlFileContent := range htmlFilesMap {

		err := os.WriteFile(currentPath+"/out/"+htmlFile, []byte(htmlFileContent), 0644)
		if err != nil {
			fmt.Println("error while writing index.html content")
			os.Exit(1)
		}
	}
}
