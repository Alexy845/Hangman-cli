package hangman

import (
	"fmt"
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
	str := [][]string{}       // Création d'une string vide
	for _, r := range runes { // Boucle pour convertir le slice de runes en string
		if r == 0 { // Si le caractère est égal à 0
			str = append(str, strings.Split(Print_letter(0), "\n")) // On ajoute un underscore
		} else { // Sinon
			str = append(str, strings.Split(Print_letter(strings.Index(w.Alpha_letter, string(r))+1), "\n")) // On ajoute le caractère
		}
	}
	fmt.Println(str[0][0])
	str1 := ""
	for i := 0; i < len(w.Word_runes); i++ {
		for j := 0; j < len(str[i]); j++ {
			for _, s := range str {
				str1 += s[j]
			}
			fmt.Println(str1)
			str1 = ""
		}
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

func Print_letter(index int) string {
	file := Open_file("asset/standard.txt", "SEP")
	return file[index]
}
