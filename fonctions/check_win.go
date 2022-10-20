package hangman

// Vérifie si le mot est trouvé
func (w *Word) Check_win() bool {
	for i := 0; i < len(w.The_word); i++ {
		if w.Word_runes[i] != rune(w.The_word[i]) {
			return false
		}
	}
	return true
}
