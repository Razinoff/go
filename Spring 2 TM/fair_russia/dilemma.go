package fair_russia

import "fmt"

func scanvar() int {
	fmt.Println("В жизни бывают моменты, когда зашел не в ту дверь. Но хуже всего, если за дверью был человек, так что выбери правильный вариант:")
	fmt.Println("1. Если изменил муж")
	fmt.Println("2. Если изменила жена")
	var variant int
	fmt.Scan(&variant)
	return variant
}

func Result() {
	p := scanvar()
	switch p {
	case 1, 2:
		fmt.Println("Забрать у мужа имущество и детей и выставить на алименты")
	default:
		fmt.Println("Рептилойды не участвуют в процессе суда")
	}
}
