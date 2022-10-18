package hangman

// Vérifie si le mot est trouvé
func (w *Word) Check_word(input string) bool {
	if w.The_word != input { // Si la l'entrée n'est pas égal au mot
		return false // On retourne faux
	}
	return true // Sinon on retourne vrai
}
