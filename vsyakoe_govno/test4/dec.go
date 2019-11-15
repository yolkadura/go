package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	var text string
    
    fmt.Print("What file you want decode?: ")
    fmt.Scanf("%s", &text) //чтение ввода имени файла юзером

    bs, err := ioutil.ReadFile(text) //чтение файла
    if err != nil {
        fmt.Println("Error", err)
    } 


    for i := 0; i < len(bs); i++ {
		
		bs[i] = bs[i] - 4
	
    }

	err2 := ioutil.WriteFile("test.txt.dec", bs, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

}
