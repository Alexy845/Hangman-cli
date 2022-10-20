package hangman

// Check if the word is correct
func (w *Word) Check_word(input string) bool {
	return w.The_word == input
}
