package util

func SkipSixth(str string) string {
	content := collectAll(str)
	runeStr := []rune(content)
	results := []rune{}

	if len(runeStr) < 5 {
		return "Invalid Input"
	}

	for i, char := range runeStr {
		if (i+1)%6 == 0 && i < len(runeStr)-1 {
			results = append(results, ' ')
			continue
		}
		results = append(results, char)
	}
	return string(results)
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
