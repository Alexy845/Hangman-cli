package hangman

import (
	"fmt"
	"strconv"
	"strings"
)

// Vérifie si la lettre est dans le mot
func (w *Word) Check_letter(letter rune) {
	if strings.ContainsRune(w.The_word, letter) { // Si la lettre est dans le mot
		for i := 0; i < len(w.The_word)-1; i++ { // Pour chaque lettre du mot
			if rune(w.The_word[i]) == letter { // Si la lettre est égale à la lettre entrée
				w.Word_runes = Replace_rune(w.Word_runes, letter, i) // On remplace la lettre dans le slice
			}
		}
	} else { // Si la lettre n'est pas dans le mot
		w.Attempts-- // On perd une vie
		fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
		Affiche(9 - w.Attempts) // On affiche le pendu
	}
}
