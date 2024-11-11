package findwords

import "strings"

func findWords(words []string) []string {
	Topline := "qwertyuiop"
	Middleline := "asdfghjkl"
	Lowerline := "zxcvbnm"
	var result []string
	for _, value := range words {
		if strings.Contains(Topline, strings.ToLower(value)) || strings.Contains(Middleline, strings.ToLower(value)) || strings.Contains(Lowerline, strings.ToLower(value)) {
			result = append(result, value)
		}
	}
	return result
}
