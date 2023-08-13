package helpers

// CheckError handler 
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}