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
	//p = rand.Intn(5)
	//a = rand.Intn(3)
	p = 0
	a = 0
	

	for i := 0; i < 5; i++ { //заполнение волнами
		for j := 0; j < 5; j++ {
			deck[i][j] = "~"
		}
	}
	
	switch p {
	case 0:
		switch a {
		case 0:
			deck[0][1] = "o"
			deck[0][0] = "o"
			deck[1][0] = "o"
		case 1:
			for i := p; i < p+3; i++ {
				deck[p][i] = "o"
			}
		case 2:
			deck[4][0] = "o"
			deck[3][0] = "o"
			deck[4][1] = "o"
		}
	case 1, 2, 3:
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
		
	case 4:
		switch a {
		case 0:
			deck[0][4] = "o"
			deck[0][3] = "o"
			deck[1][4] = "o"
		case 1:
			for i := p-1; i < p+2; i++ {
				deck[p][i] = "o"
			}
		}
	}

	for j := 0; j < 5; j++ { //вывод
		fmt.Println(deck[j])
	}

	fmt.Println(p)
	fmt.Println(a)
}
