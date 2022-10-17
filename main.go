package main

import (
	"fmt"
	hangman "hangman/fonctions"
	"log"
	"os"
)

func main() {

	switch len(os.Args) {
	case 1:
		fmt.Print("Please enter dictionary file")
		return
	case 2:
		dico := os.Args[1]
		if _, err := os.Stat(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10}
		word.N = len(word.The_word)/2 - 1
		word.Word_runes = make([]rune, len(word.The_word)-1)
		word.Rand_letters()
		word.Play()
	default:
		fmt.Print("Too many arguments !")
	}
}
