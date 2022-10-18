package main

import (
	"fmt"
	hangman "hangman/fonctions"
	"log"
	"os"
	"runtime"
)

func main() {
	switch len(os.Args) { // Vérifie le nombre d'arguments passés au programme (doit être égal à 2)
	case 1: // Si il n'y a pas d'arguments
		fmt.Print("Please enter dictionary file")
		return // On sort du programme
	case 2: // Si il y a un argument
		dico := os.Args[1] // Récupère le nom du fichier dictionnaire
		if _, err := os.Stat(os.Args[1]); err != nil {
			log.Fatal(err) // Vérifie que le fichier existe
		}
		word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10} // Création de l'objet Word
		word.N = len(word.The_word)/2 - 1                                                                 // Nombre de lettres à trouver
		os := runtime.GOOS
		switch os {
		case "windows":
			word.Word_runes = make([]rune, len(word.The_word)-1) // Création d'un slice de rune de la taille du mot à trouver (Windows)
		case "linux":
			word.Word_runes = make([]rune, len(word.The_word)) // Création d'un slice de rune de la taille du mot à trouver (Linux)
		default:
			word.Word_runes = make([]rune, len(word.The_word)) // Création d'un slice de rune de la taille du mot à trouver (Autre OS)
		}
		word.Rand_letters() // Rempli le tableau de runes avec des lettres aléatoires
		word.Play()         // Début du jeu
	default: // Si il y a plus d'un argument
		fmt.Print("Too many arguments !")
	}
}
