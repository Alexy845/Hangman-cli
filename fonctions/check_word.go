package hangman

// Vérifie si le mot est trouvé
func (w *Word) Check_word(input string) bool {
	return w.The_word == input
}
