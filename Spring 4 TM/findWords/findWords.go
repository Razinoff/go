package findwords

import "strings"

func findWords(words []string) []string {
	Topline := "qwertyuiop"
	Middleline := "asdfghjkl"
	Lowerline := "zxcvbnm"
	var result []string
	var flag bool
	for _, value := range words {
		valuelow := strings.ToLower(value)
		valuelowarr := strings.Split(valuelow, "")
		if strings.Contains(Topline, valuelowarr[0]) {
			flag = findWord(valuelowarr, Topline)
		} else if strings.Contains(Middleline, valuelowarr[0]) {
			flag = findWord(valuelowarr, Middleline)
		} else if strings.Contains(Lowerline, valuelowarr[0]) {
			flag = findWord(valuelowarr, Lowerline)
		} else {
			flag = false
		}
		if flag {
			result = append(result, value)
		}

	}

	return result
}
func findWord(strword []string, strline string) bool {
	for _, value := range strword {
		if strings.Contains(strline, value) {
		} else {
			return false
		}
	}
	return true
}
