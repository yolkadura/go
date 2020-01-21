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
    
    var i = 0
    var str string
    for i < len(bs) {
        str = string(bs[i])
        //defer fmt.Print(bs[i]," ")
        defer fmt.Print(str," ")
        i += 1
    }

}
