package i_am_so_crazy

import (
	"fmt"
	"math/rand"
)

func Weatherrand() {
	num := rand.Intn(11)
	if num > 5 {
		fmt.Println("Сегодня рептилойды просят взять зонтик")
	} else {
		fmt.Println("Сегодня рептилойды молчат, дождя не будет")
	}
}
