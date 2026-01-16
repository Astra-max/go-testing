package util

func SkipSixth(str string) string {
	content := collectAll(str)
	runeStr := []rune(content)

	if len(runeStr) < 5 {
		return "Invalid Input"
	}

	for i := 5; i < len(runeStr); i += 6 {
		runeStr[i] = ' '
	}
	return string(runeStr)
}

func collectAll(str string) string {
	results := ""

	for _, char := range str {
		if !IsSpace(char) {
			results += string(char)
		}
	}
	return results
}

func IsSpace(c rune) bool {
	return c == ' '
}
