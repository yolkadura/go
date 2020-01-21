package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	var file string
	var key string




    
    fmt.Print("What file you want encode?: ")
    fmt.Scanf("%s", &file) //чтение ввода имени файла юзером

    bs, err := ioutil.ReadFile(file) //bs принимает байты из файла, когда file хранит только имя
    if err != nil {
        fmt.Println("Error", err)
	} 

	ss := make([]byte, len(bs)) //создала срез чтоб потом в него пихнуть ключ длиной в строку

	fmt.Print("Enter key word: ")
	fmt.Scan(&key)  //ввод ключа в формате стринг
	bw := []byte(key)


for i := 0; i < len(bs); i++ { //магия по которой пока длина не совпадет с bs выводим bw

	//print(bw[i % len(bw)], " ")

	 ss[i] = bw[i % len(bw)]
	 bs[i] = bs[i] - ss[i]

	}

	err2 := ioutil.WriteFile("test.txt.dec", bs, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

}
