package asciidrawer

func isLetter(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func isNumber(s rune) bool {
	return s >= '0' && s <= '9'
}
