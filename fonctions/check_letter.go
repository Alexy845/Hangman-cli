package hangman

import (
	"strings"
	"fmt"
	"strconv"
)

// Check if the letter is in the word
func (w *Word)Check_letter(letter rune) {
	if strings.ContainsRune(w.The_word, letter) {
		for i := 0; i < len(w.The_word)-1; i++ {
			if rune(w.The_word[i]) == letter {
				w.Word_runes = Replace_rune(w.Word_runes, letter, i)
			}
		}
	} else {
		w.Attempts--
		fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
		Affiche(9-w.Attempts)
	}
}