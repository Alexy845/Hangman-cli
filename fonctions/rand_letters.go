package hangman

// Choose n random letters from the slice of runes to reveal
func (w *Word) Rand_letters() {
	var numbers int
	for i := 0; i < w.N; i++ { 
		numbers = Rand_range(0, len(w.The_word)-1)
		if w.Word_runes[numbers] == 0 {
			w.Word_runes = Replace_rune(w.Word_runes, rune(w.The_word[numbers]), numbers)
		} else {
			i--
		}
	}
}
