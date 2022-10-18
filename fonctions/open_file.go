package hangman

import (
	"log"
	"os"
	"strings"
)

// Ouvre le fichie et retourne un tableau de string couper par les sauts de ligne
func Open_file(f string, sep string) []string {
	file, err := os.ReadFile(f) // Ouvre le fichier
	if err != nil {             // Si il y a une erreur
		log.Fatal(err) // On arrête le programme
	}
	str := strings.Split(string(file), sep) // Création d'un tableau de string avec les lignes du fichier
	return str                              // Retourne le tableau de string
}
