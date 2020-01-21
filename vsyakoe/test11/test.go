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
	)
	
	fmt.Print("What file you want encode?: ")
	fmt.Scanf("%s", &file) //чтение ввода имени файла юзером

	bs, err := ioutil.ReadFile(file) //bs принимает байты из файла, когда file хранит только имя
	if err != nil {
		fmt.Println("Error", err)
	}
	j = len(bs) //количество символов без знаков
	m := make(map[string]int) //мапа
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

	for i := 0; i < len(bs); i++ {
		m[string(bs[i])] += 1
	}

	fmt.Println(string(bs))
	fmt.Println(bs)
	fmt.Println(result)
	fmt.Println("count letters ", m)

}
