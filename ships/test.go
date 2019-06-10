package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var (
		deck [5][5] string //дека
		p int			//поинт - центр корабля
		a int			//поинт - рандом для оси
	)
	rand.Seed(time.Now().UnixNano())
	p = rand.Intn(3) + 1
	a = rand.Intn(2)
	//p = 1
	//a = 1
	

	for i := 0; i < 5; i++ { //заполнение волнами
		for j := 0; j < 5; j++ {
			deck[i][j] = "~"
		}
	}
	
	switch a {
	case 0:
		for i := p-1; i < p+2; i++ {
			deck[i][p] = "o"
		}
	case 1:
		for i := p-1; i < p+2; i++ {
			deck[p][i] = "o"
		}
	}


	for j := 0; j < 5; j++ { //вывод
		fmt.Println(deck[j])
	}

	fmt.Println(p)
	fmt.Println(a)
}
