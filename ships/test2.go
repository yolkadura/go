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
	n = 5 //длина массива
)

func main() {

	var (
		deck [n][n] string 	//дека
		p int				//поинт - рандом для оси
		a int				//строка
		b int				//столбец
	)

	rand.Seed(time.Now().UnixNano())

	p = rand.Intn(2)
	a = rand.Intn(n)
	b = rand.Intn(n)
	
	// p = 0 //0 - горизонталь, 1 - вертикаль
	// a = 3 
	// b = 4
	
	for i := 0; i < n; i++ { //заполнение волнами
		for j := 0; j < n; j++ {
			deck[i][j] = sea
		}
	}
	
	switch p {
	case 0: //горизонталь
			if b == 0 {
				for i := b; i< b+3; i++ {
					deck[a][i] = ship
				}
			} else if b == n-1 {
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
			} else if a == n-1 {
				for i := a-2; i< a+1; i++ {
					deck[i][b] = ship
				} 
			} else {
				for i := a-1; i< a+2; i++ {
					deck[i][b] = ship
				} 
			}
	}

	for j := 0; j < n; j++ { //вывод
		fmt.Println(deck[j])
	}

	fmt.Println(p)
	fmt.Println(a)
	fmt.Println(b)
}
