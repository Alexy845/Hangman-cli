package hangman

// Check if the word is right
func (w *Word) Check_word(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if w.Word_runes[i] != rune(input[i]) {
			return false
		}
	}
	return true
}
