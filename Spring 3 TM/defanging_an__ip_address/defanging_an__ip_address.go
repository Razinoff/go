package defanginganipaddress

func DefangIPaddr(address string) string {
	result := []rune{}
	for _, v := range address {
		if v == '.' {
			result = append(result, '[', '.', ']')
		} else {
			result = append(result, v)
		}
	}
	return string(result)
}

// func DefangIPaddr(address string) string {
//  return strings.ReplaceAll(address, ".", "[.]")
// }
