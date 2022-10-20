package hangman

// Arrange in alphabetical order
func AlphaSort(str []rune,) {
    for x := range str {
        y := x + 1
        for y = range str {
            if str[x] < str[y] {
                str[x], str[y] = str[y], str[x]
            }
        }
    }
}
