package internal

import (
	"github.com/Ar1ste1a/Potfile/internal/password"
	"github.com/Ar1ste1a/Potfile/internal/util"
	"os"
	"path"
	"strings"
)

var location = ".local/share/hashcat/hashcat.potfile"

func init() {
	// Initialize the Potfile location
	home := os.Getenv("~")
	if home == "" {
		home = os.Getenv("HOME")
	}
	location = path.Join(home, location)
}

func RawString() string {
	return util.ReadFileString(location)
}

func RawBytes() []byte {
	return util.ReadFileBytes(location)
}

func ParsedMap() []map[string]string {
	fullContent := util.ReadFileString(location)
	lines := strings.Split(fullContent, "\n")

	var cleanedLines []map[string]string
	for _, line := range lines {
		if line != "" {
			hash, pw := password.SplitPassword(line)
			if password.IsHexPassword(line) {
				pw = password.HexToPassword(pw)
			}
			cleanedLines = append(cleanedLines, map[string]string{"hash": hash, "password": pw})
		}
	}
	return cleanedLines
}

func Passwords() []string {
	fullContent := util.ReadFileString(location)
	lines := strings.Split(fullContent, "\n")

	var cleanedLines []string
	for _, line := range lines {
		if line != "" {
			_, pw := password.SplitPassword(line)
			if password.IsHexPassword(line) {
				pw = password.HexToPassword(pw)
			}
			cleanedLines = append(cleanedLines, pw)
		}
	}
	return cleanedLines
}

func Location() string {
	return location
}

func SetLocation(newLocation string) {
	location = newLocation
}

func Count() int {
	fullContent := util.ReadFileString(location)
	lines := strings.Split(fullContent, "\n")
	return len(lines)
}
