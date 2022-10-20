package hangman

import (
	"fmt"
	"io/ioutil"
)

// Check if there is a save
func Is_not_empty () bool {
	save, err := ioutil.ReadFile(".\\asset\\save.txt")
	if err != nil {
		fmt.Println("Peut pas lire")
		panic(err)
	}
	if save != nil{
		return true
	}
	return false
}
