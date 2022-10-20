package hangman

import (
	"log"
	"os"
	"strings"
)

// Ouvre le fichie et retourne un tableau de string couper par les sauts de ligne
func Open_file(f string, sep string) []string {
	file, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(file), sep) // Création d'un tableau de string avec les lignes du fichier
	return str
}
