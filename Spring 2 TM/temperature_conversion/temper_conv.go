package temperature_conversion

import "fmt"

func Tempconversion() {
	fmt.Print("Введите температуру в градусах Цельсия = ")
	var temp float64
	fmt.Scan(&temp)
	fmt.Println("Температура в градусах Фаренгейта = ", temp*9/5+32, "℉")
}
