package hangman

import (
	"log"
	"os"
	"strings"

)

// Open a file and split words with a separator
func Open_file(f string, sep string) []string {
	file, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(file), sep)
	return str
}