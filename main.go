package main

import (
	"fmt"
	hangman "hangman/fonctions"
	"log"
	"os"
)

func main() {
	switch len(os.Args) { // Vérifie le nombre d'arguments passés au programme (doit être égal à 2)
	case 1:
		fmt.Print("Please enter dictionary file")
		return
	case 2:
		dico := os.Args[1] // Récupère le nom du fichier dictionnaire
		if _, err := os.Stat(os.Args[1]); err != nil {
			log.Fatal(err) // Vérifie que le fichier existe
		}
		word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10} // Création de l'objet Word
		word.N = len(word.The_word)/2 - 1                                                                 // Nombre de lettres à trouver
		word.Word_runes = make([]rune, len(word.The_word)-1)                                              // Création d'un slice de rune de la taille du mot à trouver
		word.Rand_letters()                                                                               // Rempli le tableau de runes avec des lettres aléatoires
		word.Play()                                                                                       // Début du jeu
	default:
		fmt.Print("Too many arguments !")
	}
}
