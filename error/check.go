package error

import (
	"log"
)

// Displays error if it exists
func Check(err error) {
	if err != nil {
		log.Fatalf("An Error occured\n%v", err)
	}
}
