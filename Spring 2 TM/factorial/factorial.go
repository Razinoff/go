package factorial

import "fmt"

func Factorial() {
	fmt.Print("Введите число: ")
	var num int
	fmt.Scan(&num)
	num1 := 1
	for num > 0 {
		num1 *= num
		num--
	}
	fmt.Println(num1)

}
