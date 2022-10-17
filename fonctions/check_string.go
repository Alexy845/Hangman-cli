package hangman

import (
	"strings"
	"fmt"
	"strconv"
)

// Check if the letter is in the word
func (w *Word)Check_string(letter string) {
	if strings.Contains(w.The_word, letter) {
		for i := 0; i < len(w.The_word)-1; i++ {
			if string(w.The_word[i]) == letter {
				w.Word_runes = Replace_string(w.Word_runes, letter, i)
			}
		}
	} else {
		w.Attempts--
		fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
		Affiche(9-w.Attempts)
	}
}