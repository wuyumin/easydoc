package utils

import (
	"fmt"
	"os"
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
