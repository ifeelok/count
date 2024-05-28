package count

import "unicode"

// CountChineseCharacter get the number of Chinese characters in a string, output an int
func CountChineseCharacter(s string) int {
	count := 0
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			count++
		}
	}
	return count
}

// GetChineseCharacterCount get the number of Chinese characters in a string, output a map
func GetChineseCharacterCount(s string) map[string]int {
	countMap := make(map[string]int)
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			countMap[string(r)]++
		}
	}
	return countMap
}
