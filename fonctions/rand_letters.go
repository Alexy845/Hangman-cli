package hangman

import (
	"strings"
)

// Choose n random letters from the slice of runes to reveal
func (w *Word) Rand_letters() {

	var numbers int
	for i := 0; i < w.N; i++ {
		numbers = Rand_range(0, len(w.The_word)-1)
		if w.Word_runes[numbers] == 0 {
			w.cl(rune(w.The_word[numbers]))
		} else {
			i--
		}
	}
}

func (w *Word) cl(letter rune) {
	if strings.ContainsRune(w.The_word, letter) {
		for i := 0; i < len(w.The_word); i++ {
			if rune(w.The_word[i]) == letter {
				w.Word_runes = Replace_rune(w.Word_runes, letter, i)
				w.N--
			}
		}
	}
}
