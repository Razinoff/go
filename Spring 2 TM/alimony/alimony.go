package alimony

import (
	"fmt"
)

func Alimony() {

	var num_det int
	var dolg1 string
	var dolg float64
	var zp float64

	//Функция нужная для проверки количества детей
	proverka_num := func() int {
		fmt.Print("Введите количество детей:")
		fmt.Scan(&num_det)
		return num_det
	}
	//Функция нужная для проверки суммы задолжности
	proverka_dolg := func() float64 {
		fmt.Print("Сумма задолжности:")
		fmt.Scan(&dolg)
		return dolg
	}
	//Функция нужная для проверки зарплаты
	proverka_zp := func() float64 {
		fmt.Print("Введите зароботную плату с учетом НДФЛ:")
		fmt.Scan(&zp)
		return zp
	}

	//Проверка количества детей
	proverka_num()
	for num_det < 0 {
		fmt.Println("Отрицательное количество детей? Это рофлс? Напиши нормально!")
		proverka_num()
	}

	fmt.Print("Есть ли задолженность по платежам:")
	fmt.Scan(&dolg1)

	//Проверка положительности задолжности
	if dolg1 == "да" || dolg1 == "Да" {
		proverka_dolg()
		for dolg < 0 {
			fmt.Println("Отрицательное долг? Дети, где нафиг мои деньги? Это рофл? Напиши нормально!")
			proverka_dolg()
		}
	}

	//Проверка положительности зарплаты
	proverka_zp()
	for zp < 0 {
		fmt.Println("Отрицательное зарплата? Соболезную( Смени работу и попробуй снова)!")
		proverka_zp()
	}

	//Вычисление алиментов
	final := func() float64 {
		switch num_det {
		case (0):
			return 0.0
		case (1):
			return zp / 4
		case (2):
			return zp / 3
		default:
			return zp / 2
		}
	}

	//Вычисление конечной суммы алиментов
	if dolg1 == "да" || dolg1 == "Да" {
		if final()+dolg > zp*0.7 {
			fmt.Println("Ваш выплаты по элементам в этом месяце:", zp*0.7)
		} else {
			fmt.Println("Ваш выплаты по элементам в этом месяце:", final()+dolg)
		}
	} else {
		fmt.Println("Ваш выплаты по элементам в этом месяце:", final())
	}
}
