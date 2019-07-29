package main

import (
	"fmt"
	"math/rand"
	"time"
)
const (
	ship = "o"
	sea = "~"
	//shot = "x"
	n = 5 //длина массива
)

func ship3 (deck*[n][n] string, a, b int) { //кораблик 3
if deck[a-1][b] == "~" && deck[a-2][b] == "~" {
	deck[a][b] = "x"
	deck[a-1][b] = "x"
	deck[a-2][b] = "x"
}
if deck[a][b+1] == "~" && deck[a][b+2] == "~" {
	deck[a][b] = "x"
	deck[a][b+1] = "x"
	deck[a][b+2] = "x"
}
if deck[a+1][b] == "~" && deck[a+2][b] == "~" {
	deck[a][b] = "x"
	deck[a+1][b] = "x"
	deck[a+2][b] = "x"
}
if deck[a][b-1] == "~" && deck[a][b-2] == "~" {
	deck[a][b] = "x"
	deck[a][b-1] = "x"
	deck[a][b-2] = "x"
}
}

func main() {

	var (
		deck [n][n] string 	//дека
		a int				//строка
		b int				//столбец
	)

	rand.Seed(time.Now().UnixNano())

	a = rand.Intn(n)
	b = rand.Intn(n)
	
	for i := 0; i < n; i++ { //заполнение волнами
		for j := 0; j < n; j++ {
			deck[i][j] = sea
		}
	}
	
	ship3 (&deck,a,b)

	for j := 0; j < n; j++ { //вывод
		fmt.Println(deck[j])
	}

	fmt.Println("a",a)
	fmt.Println("b",b)
}
