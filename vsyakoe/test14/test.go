package main

import (
	"fmt"
)

func defangIPaddr(address string) string {
	var (
		b [][]byte
		temp []byte
		c string
	)
    a := []byte(address)
    for i := range a {
        if a[i] == 46 {
				if temp != nil {
					b = append(b, temp)
				}
				temp = nil
				continue
        }
		temp = append(temp, a[i])
	}
	if temp != nil {
		b = append(b, temp)
	}

	c = string(b[0])+"[.]"+string(b[1])+"[.]"+string(b[2])+"[.]"+string(b[3])
    return c
}

func main() {
	address := "255.100.50.0"
    x := defangIPaddr(address)
    fmt.Print(x)
}
