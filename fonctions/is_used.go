package hangman

// Vérifie si la lettre a déjà été utilisée
func (w *Word) Is_used(letter rune) bool {
	for _, l := range w.Letter_used {
		if l == letter {
			return true
		}
	}
	return false
}
