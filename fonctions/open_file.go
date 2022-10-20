package hangman

import (
	"log"
	"os"
	"strings"
)

// Open the file and return an array of strings cut by line breaks
func Open_file(f string, sep string) []string {
	file, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(file), sep)
	return str
}
