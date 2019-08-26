package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	// "math/rand"
	// "time"
)
const (
	ship = "o"
	sea = "~"
	text string
	//shot = "x"
	n = 5 //длина массива
)

func coord (text string) {
	fmt.Print(text)									
	_, line, err = input.Scan(), input.Text(), input.Err()		
	if err != nil {
		fmt.Errorf("can't find point, %v", err)
	}

	x, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid point, %v", err)
	}
	return x
}

func main() {

	var (
		deck [n][n] string 	//дека
		x int
		y int
		line string
		err error

	)
	
	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ { //заполнение волнами
		for j := 0; j < n; j++ {
			deck[i][j] = sea
		}
	}

	for j := 0; j < n; j++ { //вывод
		fmt.Println(deck[j])
	}


	coord("Куда поставим кораблик? Координата Х: ")

	fmt.Print("Куда поставим кораблик? Координата Y: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//month
	if err != nil {
		fmt.Errorf("can't find point, %v", err)
	}

	y, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid point, %v", err)
	}


	fmt.Print(x,y)
}
