package utils

// Handle the error
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
