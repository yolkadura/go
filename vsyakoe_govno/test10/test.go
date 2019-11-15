package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	var (
		file   string
		temp   []byte
		result [][]byte
		k      int
		j      int
	//	p      []int
		b      int
		c	   string
	)
	
	fmt.Print("What file you want encode?: ")
	fmt.Scanf("%s", &file) //чтение ввода имени файла юзером

	bs, err := ioutil.ReadFile(file) //bs принимает байты из файла, когда file хранит только имя
	if err != nil {
		fmt.Println("Error", err)
	}
	j = len(bs) //количество символов без знаков
	p := make([]int,len(bs)) //массив для количества
	for i := 0; i < len(bs); i++ {

		switch bs[i] {
		case 32, 46, 44, 33, 63, 92, 10, 13:
			j = j - 1
			if temp != nil {
				result = append(result, temp)
			}
			temp = nil
			continue
		case 122:
			k = k + 1
		}
		temp = append(temp, bs[i])

	}

	for l := 0; l < len(bs); l++ {
		for i := 0; i < len(bs); i++ {
			if bs[i] == bs[l] {
				p[l] = p[l] + 1
			}
		}
	}

	if temp != nil {
		result = append(result, temp)
	}
	temp = nil

	fmt.Println(string(bs))
	fmt.Println(result)
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
		fmt.Println(string(bs[i]), " ", p[i])
		}
		if b < p[i] {
			b = p[i]
			c = string(bs[i])
		}
	}
	
	fmt.Println("amount of words: ", len(result), "\namount of letters", j)
	fmt.Println ("most popular letter is ", c, " was used ", b)


}
