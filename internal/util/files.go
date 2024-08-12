package util

import (
	"fmt"
	"os"
)

func ReadFileString(fp string) string {
	fBytes, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return ""
	}
	return string(fBytes)
}

func ReadFileBytes(fp string) []byte {
	fBytes, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return []byte{}
	}
	return fBytes
}
