package helpers

import (
	"fmt"
	"os"
)

func Copy(src string, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("There was a problem reading file to copy", src)
		os.Exit(1)
	}
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Println("There was a problem writhing file to copy", dst)
		os.Exit(1)
	}
}
