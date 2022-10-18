package hangman

// Vérfie le système d'exploitation et adapte le programme en conséquence
func (w Word) Check_os() int {
	switch w.OS { // Selon le système d'exploitation
	case "windows": // Si le système d'exploitation est windows
		return 1 // On retourne 1
	case "linux": // Si le système d'exploitation est linux
		return 0 // On retourne 0
	default: // Si le système d'exploitation est autre
		return 0 // On retourne 0
	}
}
