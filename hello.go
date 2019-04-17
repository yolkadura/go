package main

import "fmt"

var (
	a int = 1
	b int = 2
	
)

func main() {
	
	p := &a         
	fmt.Println(*p) 
	*p = 21         // set i through the pointer
	fmt.Println(a)  // see the new value of i

	p = &b         // point to j
	*p = *p * 2   // divide j through the pointer
	fmt.Println(b) // see the new value of j
}