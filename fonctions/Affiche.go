package hangman

import (
	"fmt"
)

// Créer l'affichage du pendu
func Affiche(index int) {
	hangman := Open_file("asset/hangman.txt", "\n") // Création d'un tableau de string avec les lignes du fichier
	min := 0                                        // Index de début
	max := 8                                        // Index de fin
	for i := 0; i < index; i++ {                    // Boucle pour afficher les lignes du pendu
		min += 8 // Incrémentation de l'index de début
		max += 8 // Incrémentation de l'index de fin
	}
	for i := min; i < max; i++ { // Boucle pour afficher les lignes du pendu
		fmt.Println(hangman[i]) // Affiche la ligne
	}
}

// Converti le slice de runes en string
func AfficheRune(runes []rune) string {
	str := ""                 // Création d'une string vide
	for _, r := range runes { // Boucle pour convertir le slice de runes en string
		if r == 0 { // Si le caractère est égal à 0
			str += "_" // On ajoute un underscore
		} else { // Sinon
			str += string(r) // On ajoute le caractère
		}
	}
	return str // Retourne la string
}

func Print_lose() {
	fmt.Printf("  ___    ___ ________  ___  ___          ___       ________  ________  _______           ___       \n" +
		" |\\  \\  /  /|\\   __  \\|\\  \\|\\  \\        |\\  \\     |\\   __  \\|\\   ____\\|\\  ___ \\         |\\  \\      \n" +
		" \\ \\  \\/  / | \\  \\|\\  \\ \\  \\\\\\  \\       \\ \\  \\    \\ \\  \\|\\  \\ \\  \\___|\\ \\   __/|        \\ \\  \\     \n" +
		"  \\ \\    / / \\ \\  \\\\\\  \\ \\  \\\\\\  \\       \\ \\  \\    \\ \\  \\\\\\  \\ \\_____  \\ \\  \\_|/__       \\ \\  \\    \n" +
		"   \\/  /  /   \\ \\  \\\\\\  \\ \\  \\\\\\  \\       \\ \\  \\____\\ \\  \\\\\\  \\|____|\\  \\ \\  \\_|\\ \\       \\ \\__\\   \n" +
		" __/  / /      \\ \\_______\\ \\_______\\       \\ \\_______\\ \\_______\\____\\_\\  \\ \\_______\\       \\|__|   \n" +
		"|\\___/ /        \\|_______|\\|_______|        \\|_______|\\|_______|\\_________\\|_______|           ___ \n" +
		"\\|___|/                                                        \\|_________|                   |\\__\\\n" +
		"                                                                                              \\|__|\n")
}

func Print_win() {
	fmt.Printf("  ___    ___ ________  ___  ___          ___       __   ___  ________           ___       \n" +
		" |\\  \\  /  /|\\   __  \\|\\  \\|\\  \\        |\\  \\     |\\  \\|\\  \\|\\   ___  \\        |\\  \\      \n" +
		" \\ \\  \\/  / | \\  \\|\\  \\ \\  \\\\\\  \\       \\ \\  \\    \\ \\  \\ \\  \\ \\  \\\\ \\  \\       \\ \\  \\     \n" +
		"  \\ \\    / / \\ \\  \\\\\\  \\ \\  \\\\\\  \\       \\ \\  \\  __\\ \\  \\ \\  \\ \\  \\\\ \\  \\       \\ \\  \\    \n" +
		"   \\/  /  /   \\ \\  \\\\\\  \\ \\  \\\\\\  \\       \\ \\  \\|\\__\\_\\  \\ \\  \\ \\  \\\\ \\  \\       \\ \\__\\   \n" +
		" __/  / /      \\ \\_______\\ \\_______\\       \\ \\____________\\ \\__\\ \\__\\\\ \\__\\       \\|__|   \n" +
		"|\\___/ /        \\|_______|\\|_______|        \\|____________|\\|__|\\|__| \\|__|           ___ \n" +
		"\\|___|/                                                                              |\\__\\\n" +
		"                                                                                     \\|__|\n")
}
