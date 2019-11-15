package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	var text string
    
    fmt.Print("What file you want open?: ")
    fmt.Scanf("%s", &text) //чтение ввода имени файла юзером

    bs, err := ioutil.ReadFile(text) //чтение файла
    if err != nil {
        fmt.Println("Error", err)
    } 
    
    var i = len(bs) - 1 
	var ii = 0
	x := make([]byte, len(bs))

    for i >= 0 {
		
		x[i] = bs[ii]
		fmt.Print(x[i])

		i = i - 1
		ii = ii + 1
    }

	message := x
	err2 := ioutil.WriteFile("test.txt.rev", message, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

}
