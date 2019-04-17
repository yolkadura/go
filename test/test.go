package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    var text string
    
    fmt.Print("What file you want open?: ")
    fmt.Scanf("%s", &text) //чтение ввода имени файла юзером

    bs, err := ioutil.ReadFile(text) //чтение файла
    if err != nil {
        //return
        fmt.Println("Error", err)
    } 
    
    str := string(bs) //перевод файла в строку

    fmt.Println(str) //вывод строки
    fmt.Println(bs) //вывод файла в байтах
    fmt.Printf("bs is of type %T\n", bs) //тип bs
    

}
