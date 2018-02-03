package utils

import (
	"fmt"
	"os"
	"strings"
)

// CheckErr handle the error
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ExistsOrMkdir is to determine the existence of the folder, Or make.
func ExistsOrMkdir(dir string) error {
	//exist?
	if _, err := os.Stat(dir); err == nil {
		return nil
	}
	//make
	err := os.MkdirAll(dir, os.ModePerm)
	return err
}

// IsExternalLink is to determine external link.
func IsExternalLink(path string) bool {
	return strings.HasPrefix(path, "https:") || strings.HasPrefix(path, "http:") || strings.HasPrefix(path, "ftp:")
}

// If is ternary operator.
func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}
