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
	switch len(os.Args) {
	case 1: // If there isn't args
		fmt.Print("Please enter dictionary file")
		return
	case 2: // If there is one arg = normal game 
		dico := os.Args[1]
		if _, err := os.Stat(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10, Letter_used: []rune{}, Alpha_letter: "abcdefghijklmnopqrstuvwxyz"} // Création de l'objet Word
		word.N = len(word.The_word)/2 - 1                                                                                                                                    // Nombre de lettres à trouver
		word.Word_runes = make([]rune, len(word.The_word))                                                                                                                   // Création d'un slice de rune de la taille du mot à trouver (Linux)
		word.Rand_letters()                                                                                                                                                  // Rempli le tableau de runes avec des lettres aléatoires
		word.Play()
	case 3 : // If there is 2 args = game from save.txt
		if os.Args[1] == "--startWith" {
			if os.Args[2] == "save.txt" {
				if hangman.Is_not_empty() {
					var jsonS, err = ioutil.ReadFile(".\\asset\\" + os.Args[2])
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
			}
		}
	default: // If there is 3 args or more
		fmt.Print("Too many arguments !")
	}
}
