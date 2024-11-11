package lengthoflastword

import "strings"

func LengthOfLastWord(s string) int {
	str := strings.Fields(s)
	return len(str[len(str)-1])
}
