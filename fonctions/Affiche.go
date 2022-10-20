package hangman

import (
	"fmt"
	"os"
	"strings"
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
func (w Word) AfficheRune(runes []rune) {
	str := make([]string, 9)  // Création d'une string vide
	for _, r := range runes { // Boucle pour convertir le slice de runes en string
		if r == 0 { // Si le caractère est égal à 0
			for i, line := range Print_letter('_') {
				str[i] = str[i] + line
			}
		} else { // Sinon
			for i, line := range Print_letter(int(r)) {
				str[i] = str[i] + line
			}
		}
	}
	for _, line := range str {
		println(line)
	}
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
func Print_letter(index int) []string {
	current_min := 9 * (index - 32)
	current_max := current_min + 9
	content, _ := os.ReadFile("asset/standard.txt")
	strcontent := strings.Split(string(content), "\n")
	ascii := []string{}
	for i := current_min; i < current_max; i++ {
		line := ""
		for _, r := range strcontent[i] {
			if r >= 32 && r < 127 {
				line += string(r)
			}
		}
		ascii = append(ascii, line)
	}
	return ascii
}
