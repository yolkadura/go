package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	var (
		file string
		k int = 1
	)

    fmt.Print("What file you want encode?: ")
    fmt.Scanf("%s", &file) //чтение ввода имени файла юзером

    bs, err := ioutil.ReadFile(file) //bs принимает байты из файла, когда file хранит только имя
    if err != nil {
        fmt.Println("Error", err)
	} 

	for i := 0; i < len(bs); i++ {
		//fmt.Printf("byte %d is\n", bs[i])
		if bs[i] == 32 {
			k = k+1
		}

		}

		fmt.Printf("%d", k)

}
