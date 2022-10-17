package hangman

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Create the Hangman Affiche
func Affiche(index int) {
	file, err := os.ReadFile("asset/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	hangman := []string(strings.Split(string(file), "\n"))
	min := 0
	max := 8
	for i := 0; i < index; i++ {
		min += 8
		max += 8
	}
	for i := min; i < max; i++ {
		fmt.Println(hangman[i])
	}
}

//  Affiche the slice of rune in a string
func AfficheRune(runes []rune) string {
	str := ""
	for _, r := range runes {
		if r == 0 {
			str += "_"
		} else {
			str += string(r)
		}
	}
	return str
}
