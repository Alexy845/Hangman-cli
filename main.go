package main

import (
	"encoding/json"
	"fmt"
	hangman "hangman/fonctions"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	switch len(os.Args) { // Vérifie le nombre d'arguments passés au programme (doit être égal à 2)
	case 1: // Si il n'y a pas d'arguments
		fmt.Print("Please enter dictionary file")
		return // On sort du programme
	case 2: // Si il y a un argument
		if os.Args[1] == "--startWith" {
			if hangman.Is_not_empty() {
				// var jsonS = `{"the_word": "", "word_runes": [90],"n": 0, "attempts": 10, "hangmanPositions": [], "letter_used": []}`
				var jsonS, err = ioutil.ReadFile("./asset/save.txt")
				if err != nil {
					panic(err)
				}
				var object_save hangman.Word
				err = json.Unmarshal([]byte(jsonS), &object_save)
				if err != nil {
					panic(err)
				}
				object_save.Play()
			} else {
				fmt.Println("No save found ")
			}
		} else {
			dico := os.Args[1] // Récupère le nom du fichier dictionnaire
			if _, err := os.Stat(os.Args[1]); err != nil {
				log.Fatal(err) // Vérifie que le fichier existe
			}
			word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10, Letter_used: []rune{}, Alpha_letter: "abcdefghijklmnopqrstuvwxyz"} // Création de l'objet Word
			word.N = len(word.The_word)/2 - 1                                                                                                                                    // Nombre de lettres à trouver
			word.Word_runes = make([]rune, len(word.The_word))                                                                                                                   // Création d'un slice de rune de la taille du mot à trouver (Linux)
			word.Rand_letters()                                                                                                                                                  // Rempli le tableau de runes avec des lettres aléatoires
			word.Play()
		}
	default: // Si il y a plus d'un argument
		fmt.Print("Too many arguments !")
	}
}
