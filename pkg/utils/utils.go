package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// SplitLines takes a multiline string and splits it on newlines
// currently we are also stripping \r's which may have adverse effects for
// windows users (but no issues have been raised yet)
func SplitLines(multilineString string) []string {
	multilineString = strings.Replace(multilineString, "\r", "", -1)
	if multilineString == "" || multilineString == "\n" {
		return make([]string, 0)
	}
	lines := strings.Split(multilineString, "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}

// WithPadding pads a string as much as you want
func WithPadding(str string, padding int) string {
	if padding-len(str) < 0 {
		return str
	}
	return str + strings.Repeat(" ", padding-len(str))
}

// ColoredString takes a string and a colour attribute and returns a colored
// string with that attribute
func ColoredString(str string, colorAttribute color.Attribute) string {
	colour := color.New(colorAttribute)
	return ColoredStringDirect(str, colour)
}

// ColoredStringDirect used for aggregating a few color attributes rather than
// just sending a single one
func ColoredStringDirect(str string, colour *color.Color) string {
	return colour.SprintFunc()(fmt.Sprint(str))
}

// GetCurrentRepoName gets the repo's base name
func GetCurrentRepoName() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return filepath.Base(pwd)
}

// TrimTrailingNewline - Trims the trailing newline
// TODO: replace with `chomp` after refactor
func TrimTrailingNewline(str string) string {
	if strings.HasSuffix(str, "\n") {
		return str[:len(str)-1]
	}
	return str
}
