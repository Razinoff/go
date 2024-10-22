package string_po_num

import "fmt"

func Stringnum() {
	fmt.Print("Введите строку: ")
	var stroka string
	fmt.Scan(&stroka)

	runes := []rune(stroka)

	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}
}
