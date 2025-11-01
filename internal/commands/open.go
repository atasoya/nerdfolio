package commands

import (
	"fmt"
	"os/exec"
	"runtime"
)

func HandleOpenCommand() {
	filePath := "out/index.html"
	err := openInBrowser(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Opened", filePath, "in browser")
	}
}

func openInBrowser(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
