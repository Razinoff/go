package wife_witch

import (
	"fmt"
	"math/rand"
)

// Варианты пила
func variant() {
	fmt.Println("Жена пилит:")
	num := rand.Intn(5)
	switch num {
	case 0:
		fmt.Println("Найди вторую работу")
	case 1:
		fmt.Println("Выгуляй шпица-Масю")
	case 2:
		fmt.Println("Купи мне сумочку")
	case 3:
		fmt.Println("На выходные к нам приедет моя мама")
	case 4:
		fmt.Println("Приготовь покушать")
	case 5:
		fmt.Println("Убери дом")
	}
}

func Result() {
	var answer string
	//Функция для функции проверки ответа
	otvet := func() {
		fmt.Println("Ты скуф?")
		fmt.Scan(&answer)
	}

	otvet()
	//Функция для проверки ответа(если бред)
	erro := func() {
		for answer != "Да" && answer != "да" && answer != "Нет" && answer != "нет" {
			fmt.Println("Не пиши бред, напиши нормально")
			otvet()
		}
	}

	//Функция для проверки ответа(если да)
	terpila := func() {
		for answer == "Да" || answer == "да" {
			variant()
			fmt.Println("Терпеть?")
			fmt.Scan(&answer)
		}
	}

	//Финальная проверка ответа и пил
	for {
		if answer == "Нет" || answer == "нет" {
			fmt.Println("Красава, держи альтушку в подарок!")
			break
		} else if answer == "Да" || answer == "да" {
			terpila()
		} else {
			erro()
		}
	}

}
