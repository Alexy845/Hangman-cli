package hangman

import (
	"io/ioutil"
	"log"
)

// Check if there is a save
func Is_not_empty() bool {
	save, err := ioutil.ReadFile(".\\asset\\save.txt")
	if err != nil {
		log.Fatal(err)
	}
	if save != nil {
		return true
	}
	return false
}
