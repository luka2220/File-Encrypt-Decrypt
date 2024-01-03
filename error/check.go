package error

// Displays error if it exists
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
