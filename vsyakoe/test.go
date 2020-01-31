package main

import (
	"fmt"
)

func main()  {

var x = 0
x = pointer(x)
fmt.Println("x =",x)

}

func pointer(x int) int {
	x = 2
	return x
}
