package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yuin/goldmark"
	"io"
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
	outputDirectory := filepath.Join(currentPath, "out")
	createOutDirectoryIfNotExsists(outputDirectory)

	err := os.CopyFS(filepath.Join(outputDirectory, "public"), os.DirFS(filepath.Join(currentPath, "public")))
	if err != nil {
		fmt.Println("Problem ocurred while copying public directory")
		os.Exit(1)
	}

	htmlFilesMap := createHtmlFilesMap()

	htmlFilesMap, err = replaceTemplates(htmlFilesMap, currentPath)
	if err != nil {
		fmt.Println("Problem ocurred while replacing templates")
		os.Exit(1)
	}

	rawJson, flatJson := createJsonDataMaps()

	htmlFilesMap, err = replaceLoops(rawJson, htmlFilesMap)
	if err != nil {
		fmt.Println("Error ocurred while replacing loops")
		os.Exit(1)
	}

	htmlFilesMap, err = replaceJsonData(flatJson, htmlFilesMap)
	if err != nil {
		fmt.Println("Error occured while replacing json data")
		os.Exit(1)
	}

	writeHtmlFilesToOutDirectory(htmlFilesMap)
	flags.HandleColorSchemeFlag(BuildColorSchemeFlag, currentPath)

	fmt.Printf("Build completed in %d ms\n", time.Since(start).Milliseconds())
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

func replaceJsonData(jsonDataMap map[string]any, htmlFilesMap map[string]string) (map[string]string, error) {
	for htmlFile, content := range htmlFilesMap {
		for name, data := range jsonDataMap {
			re := regexp.MustCompile(`\{\{\s*` + regexp.QuoteMeta(name) + `\s*\}\}`)
			content = re.ReplaceAllString(content, fmt.Sprintf("%v", data))
		}
		htmlFilesMap[htmlFile] = content
	}
	return htmlFilesMap, nil
}

func createHtmlFilesMap() map[string]string {
	htmlFilesMap := make(map[string]string)
	entries, _ := os.ReadDir(currentPath)

	for _, entry := range entries {
		if !entry.IsDir() {
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

			htmlFilesMap[filepath.Join("out", fileName)] = string(data)
			continue
		}

		if entry.IsDir() && entry.Name() == "blogs" {
			blogPath := filepath.Join(currentPath, "blogs")
			mdEntries, err := os.ReadDir(blogPath)
			if err != nil {
				fmt.Println("Error reading blogs directory:", err)
				os.Exit(1)
			}

			for _, mdFile := range mdEntries {
				if mdFile.IsDir() || !strings.HasSuffix(strings.ToLower(mdFile.Name()), ".md") {
					continue
				}

				mdFilePath := filepath.Join(blogPath, mdFile.Name())
				mdData, err := os.ReadFile(mdFilePath)
				if err != nil {
					fmt.Println("Error reading markdown file:", mdFile.Name(), err)
					os.Exit(1)
				}

				var htmlOutput strings.Builder
				if err := goldmark.Convert(mdData, &htmlOutput); err != nil {
					fmt.Println("Error converting markdown:", mdFile.Name(), err)
					os.Exit(1)
				}

				htmlFileName := strings.TrimSuffix(mdFile.Name(), ".md") + ".html"
				htmlFilesMap[filepath.Join("out", "blogs", htmlFileName)] = htmlOutput.String()
			}
		}
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
	for relPath, htmlFileContent := range htmlFilesMap {
		outputPath := filepath.Join(currentPath, relPath)

		dir := filepath.Dir(outputPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", dir, err)
			os.Exit(1)
		}

		if err := os.WriteFile(outputPath, []byte(htmlFileContent), 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", outputPath, err)
			os.Exit(1)
		}
	}
}

func createJsonDataMaps() (map[string]any, map[string]any) {
	jsonPath := filepath.Join(currentPath, "data.json")
	file, err := os.Open(jsonPath)
	if err != nil {
		return map[string]any{}, map[string]any{}
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return map[string]any{}, map[string]any{}
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return map[string]any{}, map[string]any{}
	}

	flat := make(map[string]any)
	flatten("", raw, flat)

	return raw, flat
}

func flatten(prefix string, m map[string]any, result map[string]any) {
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch val := v.(type) {
		case map[string]any:
			flatten(key, val, result)
		case []any:
			values := make([]string, len(val))
			for i, x := range val {
				values[i] = fmt.Sprintf("%v", x)
			}
			result[key] = strings.Join(values, ", ")
		default:
			result[key] = val
		}
	}
}

func replaceLoops(jsonDataMap map[string]any, htmlFilesMap map[string]string) (map[string]string, error) {
	loopRegex := regexp.MustCompile(`\{\{#each\s+([\w\.]+)\}\}([\s\S]*?)\{\{\/each\}\}`)

	for file, content := range htmlFilesMap {
		matches := loopRegex.FindAllStringSubmatch(content, -1)
		for _, match := range matches {
			keyPath := match[1]
			innerBlock := match[2]

			value, ok := getNestedValue(jsonDataMap, keyPath)
			if !ok {
				continue
			}

			var replacement strings.Builder

			switch v := value.(type) {
			case []any:
				for _, item := range v {
					block := innerBlock
					switch val := item.(type) {
					case map[string]any:
						for subKey, subVal := range val {
							re := regexp.MustCompile(`\{\{\s*` + regexp.QuoteMeta(subKey) + `\s*\}\}`)
							block = re.ReplaceAllString(block, fmt.Sprintf("%v", subVal))
						}
					default:
						re := regexp.MustCompile(`\{\{\s*this\s*\}\}`)
						block = re.ReplaceAllString(block, fmt.Sprintf("%v", val))
					}
					replacement.WriteString(block)
				}
			case map[string]any:
				for subKey, subVal := range v {
					block := innerBlock
					block = strings.ReplaceAll(block, "{{ key }}", subKey)
					block = strings.ReplaceAll(block, "{{ value }}", fmt.Sprintf("%v", subVal))
					replacement.WriteString(block)
				}
			}

			content = strings.ReplaceAll(content, match[0], replacement.String())
		}
		htmlFilesMap[file] = content
	}

	return htmlFilesMap, nil
}

func getNestedValue(data map[string]any, path string) (any, bool) {
	parts := strings.Split(path, ".")
	var current any = data
	for _, part := range parts {
		if m, ok := current.(map[string]any); ok {
			current, ok = m[part]
			if !ok {
				return nil, false
			}
		} else {
			return nil, false
		}
	}
	return current, true
}
