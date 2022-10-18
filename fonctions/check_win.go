package hangman

// Vérifie si le mot est trouvé
func (w *Word) Check_win() bool {
	for i := 0; i < len(w.The_word)-w.Check_os(); i++ { // Pour chaque lettre du mot
		if w.Word_runes[i] != rune(w.The_word[i]) { // Si la lettre est différente de la lettre du mot
			return false // On retourne faux
		}
	}
	return true // Sinon on retourne vrai
}
