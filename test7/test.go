package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	var (
		file string
		temp []byte
		result [][]byte
	)

    fmt.Print("What file you want encode?: ")
    fmt.Scanf("%s", &file) //чтение ввода имени файла юзером
	
    bs, err := ioutil.ReadFile(file) //bs принимает байты из файла, когда file хранит только имя
    if err != nil {
        fmt.Println("Error", err)
	} 

	fmt.Println(bs)

	for i := 0; i < len(bs); i++ {

		switch bs[i] {
		case 32, 46, 44, 33, 63, 92, 10, 13 : 
				if temp != nil {result = append(result, temp)} 
				temp = nil 
				continue;
		default : 
				temp = append(temp, bs[i])
		
		}

	}
	
	if temp != nil {result = append(result, temp)}
	temp = nil
		
	fmt.Println(result)
	fmt.Println("amount of words: ", len(result))

}
