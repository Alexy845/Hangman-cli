package hangman

import (
	"bufio"
)

// Check if the word is right 
func (w *Word)Check_word() bool {
	for i := 0; i < len(input.Text())-1; i++ {
		if w.Word_runes[i] != rune(input.Text()[i]) {
			return false
		}
	}
	return true
}
