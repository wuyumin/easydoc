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
