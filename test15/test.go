package main

import (
	"fmt"
)
const (
	offset = 97
)
var (
	words = []string {"gin", "zen", "gig", "omg"}
	alphabet = []string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	x [][]byte
	k string
	y []byte
	l [][]byte
	n [][]byte
	temp []byte
	s int = 0
	z [][]byte
)

func qqq(a []byte,b []byte) bool {

	if len(a) != len(b) {return false}

	for i := range a {
		if a[i] != b[i] {return false}
	}

	return true

}

func main() {
	for i := range words { //перевод массива текста в массив байт
		k = words[i]
		y = []byte(k)
		x = append(x, y)
	}

	for i := range alphabet { //перевод азбуки в байты
		k = alphabet[i]
		y = []byte(k)
		l = append(l, y)
	}

	for i := range x {
		for j := range x[i]{
			qq := l[x[i][j]-offset]
			temp = append(temp, qq...)
			if j == (len(x[i])-1) {
				n = append(n, temp)
				temp = nil
			} else {continue}
					
		} 

	}


	for i := 0; i < len(n)-1; i++ {
		if qqq(n[i], n[i+1]) == true {
			s++
		}
	}


	

	fmt.Println("x", x)
	fmt.Println("l",l)
	fmt.Println("n",n)
	fmt.Println("s",s)

	for i := range n {
		fmt.Println("n",n[i])
	}


}
