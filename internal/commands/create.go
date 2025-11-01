package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func HandleCreateCommand() {
	currentPath, _ := os.Getwd()

	isEmpty, err := IsEmpty(currentPath)
	if err != nil {
		fmt.Println("Error checking directory:", err)
		return
	}

	if !isEmpty {
		fmt.Println("Current directory is not empty!")
		return
	}

	fmt.Println("Creating project structure...")

	dirs := []string{
		"public",
		"blogs",
		"templates",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(currentPath, dir), 0755); err != nil {
			fmt.Println("Error creating directory:", dir, err)
			return
		}
	}

	indexContent := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{ title }}</title>
  <style>
    body {
      margin: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
      font-family: system-ui, sans-serif;
    }
    main {
      text-align: center;
    }
  </style>
	<link rel="stylesheet" href="nerdfolio.css">
</head>
<body>
  <main>
    {% madeby %}
    <h1>{{ title }}</h1>
	<a href = "./blogs/my-cool-blog.html"> Check out cool blog written in .md </a>
  </main>
</body>
</html>`

	indexPath := filepath.Join(currentPath, "index.html")
	if err := os.WriteFile(indexPath, []byte(indexContent), 0644); err != nil {
		fmt.Println("Error writing index.html:", err)
		return
	}

	blogContent := `{% blogHeader %}
# My Cool Blog
Welcome to my first blog post!

{% madeby %}`

	blogPath := filepath.Join(currentPath, "blogs", "my-cool-blog.md")
	if err := os.WriteFile(blogPath, []byte(blogContent), 0644); err != nil {
		fmt.Println("Error writing blog file:", err)
		return
	}

	data := map[string]any{
		"title": "Welcome to Nerdfolio!",
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	dataPath := filepath.Join(currentPath, "data.json")
	if err := os.WriteFile(dataPath, jsonData, 0644); err != nil {
		fmt.Println("Error writing data.json:", err)
		return
	}

	madeByContent := `<footer style="font-size:0.9rem; margin-top:1rem;">
<a href="https://github.com/atasoya/nerdfolio" target="_blank">powered by nerdfolio btw.</a>
</footer>`

	madeByPath := filepath.Join(currentPath, "templates", "madeby.html")
	if err := os.WriteFile(madeByPath, []byte(madeByContent), 0644); err != nil {
		fmt.Println("Error writing madeby.html:", err)
		return
	}

	blogHeaderContent := `<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Blog - {{ title }}</title>
<link rel="stylesheet" href="../nerdfolio.css" />
`

	blogHeaderPath := filepath.Join(currentPath, "templates", "blogHeader.html")
	if err := os.WriteFile(blogHeaderPath, []byte(blogHeaderContent), 0644); err != nil {
		fmt.Println("Error writing madeby.html:", err)
		return
	}

	fmt.Println("Project created successfully!")
}

func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
