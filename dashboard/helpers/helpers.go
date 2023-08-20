package helpers

// CheckError handler 
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckError with custom message
func CheckErrorWithMessage(err error, message string) {
	if err != nil {
		panic(message)
	}
}