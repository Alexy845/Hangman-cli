package hangman

// Convertit une lettre avec accent en lettre sans accent
func Convert_letter_with_accent_to_letter(letter rune) rune {
	switch letter {
	case 'à':
		return 'a'
	case 'â':
		return 'a'
	case 'ä':
		return 'a'
	case 'é':
		return 'e'
	case 'è':
		return 'e'
	case 'ê':
		return 'e'
	case 'ë':
		return 'e'
	case 'î':
		return 'i'
	case 'ï':
		return 'i'
	case 'ô':
		return 'o'
	case 'ö':
		return 'o'
	case 'ù':
		return 'u'
	case 'û':
		return 'u'
	case 'ü':
		return 'u'
	case 'ç':
		return 'c'
	default:
		return letter
	}
}

// Convertit un string contenant des accents en string sans accent
func Convert_string_with_accent_to_string(s string) string {
	var result string
	for _, letter := range s {
		result += string(Convert_letter_with_accent_to_letter(letter))
	}
	return result
}
