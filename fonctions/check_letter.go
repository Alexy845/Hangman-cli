package hangman

import (
	"fmt"
	"strconv"
	"strings"
)

// Vérifie si la lettre est dans le mot
func (w *Word) Check_letter(letter rune) {
	if strings.ContainsRune(w.The_word, letter) { 
		for i := 0; i < len(w.The_word); i++ {
			if rune(w.The_word[i]) == letter { 
				w.Word_runes = Replace_rune(w.Word_runes, letter, i) // Si oui on remplace la lettre dans le slice
			}
		}
	} else { // Si la lettre n'est pas dans le mot
		w.Attempts-- // On perd une vie
		fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
		Affiche(9 - w.Attempts)
	}
}
