package hangman

// Vérifie si le mot est trouvé
func (w *Word) Check_word(input string) bool {
	for i := 0; i < len(input)-1; i++ { // Pour chaque lettre du mot entré
		if w.Word_runes[i] != rune(input[i]) { // Si la lettre n'est pas dans le mot
			return false // On retourne faux
		}
	}
	return true // Si toutes les lettres sont dans le mot, on retourne vrai
}
