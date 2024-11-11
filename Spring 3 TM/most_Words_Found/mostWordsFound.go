package mostwordsfound

import "strings"

func MostWordsFound(sentences []string) int {
	result := 0
	for _, value := range sentences {
		if result < len(strings.Fields(value)) {
			result = len(strings.Fields(value))
		}
	}
	return result
}
