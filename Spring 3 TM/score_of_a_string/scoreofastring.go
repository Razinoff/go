package scoreofastring

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ScoreOfString(s string) int {
	runesc := []rune(s)
	var strtoint int = int(runesc[0])
	var result int
	for i := 1; i < len(runesc); i++ {
		result += absDiffInt(strtoint, int(runesc[i]))
		strtoint = int(runesc[i])

	}
	return result
}
