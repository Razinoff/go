package main

import (
	"Spring_2_TM/alimony"
	"Spring_2_TM/area_of_a_square"
	"Spring_2_TM/factorial"
	"Spring_2_TM/fair_russia"
	"Spring_2_TM/i_am_so_crazy"
	"Spring_2_TM/string_po_num"
	"Spring_2_TM/temperature_conversion"
	"Spring_2_TM/wife_witch"
	"fmt"
)

func main() {
	//Вывод Hellow...
	fmt.Println("Hellow world of warcraft")

	//Вычисление площади квадрата
	area_of_a_square.Areaofsquart()

	//Предсказание погоды от канала РЕН-ТВ
	i_am_so_crazy.Weatherrand()

	//Самый гуманный суд
	fair_russia.Result()

	//Конвертация температуры
	temperature_conversion.Tempconversion()

	//Лучшая жена
	wife_witch.Result()

	//Факториал
	factorial.Factorial()

	//Вывод строки по отдельности
	string_po_num.Stringnum()

	//Не повезло, не фортануло, теперь алименты
	alimony.Alimony()

}
