package main

import (
	"fmt"
	"math/rand"
	"time"
)
const (
	ship = "o"
	sea = "~"
	shot = "x"
)

func main() {

	var (
		deck [5][5] string //дека
		p int			//поинт - рандом для оси
		a int			//строка
		b int			//столбец
	)

	rand.Seed(time.Now().UnixNano())

	p = rand.Intn(2)
	a = rand.Intn(5)
	b = rand.Intn(5)
	
	// p = 0 //0 - горизонталь, 1 - вертикаль
	// a = 3 
	// b = 4
	
	for i := 0; i < 5; i++ { //заполнение волнами
		for j := 0; j < 5; j++ {
			deck[i][j] = sea
		}
	}
	
	switch p {
	case 0: //горизонталь
			if b == 0 {
				for i := b; i< b+3; i++ {
					deck[a][i] = ship
				}
			} else if b == 4 {
				for i := b-2; i< b+1; i++ {
					deck[a][i] = ship
				} 
			} else {
				for i := b-1; i< b+2; i++ {
					deck[a][i] = ship
				} 
			}
			
	case 1: //вертикаль
			if a == 0 {
				for i := a; i< a+3; i++ {
					deck[i][b] = ship
				}
			} else if a == 4 {
				for i := a-2; i< a+1; i++ {
					deck[i][b] = ship
				} 
			} else {
				for i := a-1; i< a+2; i++ {
					deck[i][b] = ship
				} 
			}
	}

	for j := 0; j < 5; j++ { //вывод
		fmt.Println(deck[j])
	}

	fmt.Println(p)
	fmt.Println(a)
	fmt.Println(b)
}
