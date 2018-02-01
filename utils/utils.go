package utils

import (
	"fmt"
	"os"
	"strings"
)

// Handle the error
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Is directory exist? Or make.
func ExistsOrMkdir(dir string) error {
	//exist?
	if _, err := os.Stat(dir); err == nil {
		return nil
	}
	//make
	err := os.MkdirAll(dir, os.ModePerm)
	return err
}

// Is External Link
func IsExternalLink(path string) bool {
	return strings.HasPrefix(path, "https:") || strings.HasPrefix(path, "http:") || strings.HasPrefix(path, "ftp:")
}
