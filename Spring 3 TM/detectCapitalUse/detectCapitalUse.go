package detectcapitaluse

import "strings"

func DetectCapitalUse(word string) bool {
	if word == strings.ToUpper(word) || word[1:] == strings.ToLower(word[1:]) {
		return true
	} else {
		return false
	}
}
