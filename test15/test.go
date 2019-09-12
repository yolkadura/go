package main

import (
	"fmt"
	"reflect"
)
const (
	offset = 97
)
var (
	words = []string {"gin", "zen", "gig", "msg"}
	alphabet = []string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	x [][]byte
	k string
	y []byte
	l [][]byte
	n [][][]byte
	temp [][]byte
	s int = 0

)
// func uniqueMorseRepresentations(words []string) int {
    
// }

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
			temp = append(temp, l[x[i][j]-offset])
			if j == (len(x[i])-1) {
				n = append(n, temp)
				temp = nil
			} else {continue}
					
		} 

	}

	// for i := range n {
	// 	for j := i+1; j<len(n); j++ {
	// 		if n[i] == n[j] {
	// 			s++
	// 		}
	// 	}
	// }

	for i := 0; i < len(n)-1; i++ {
		if reflect.DeepEqual(n[i], n[i+1]) == true {
			s++
		}
	}

fmt.Println("x", x)
fmt.Println("l",l)
// fmt.Println("temp",temp)
fmt.Println("n",n)
fmt.Println("n",n[0])
fmt.Println("n",n[1])
fmt.Println("n",n[2])
fmt.Println("n",n[3])
// fmt.Println("n",n[0][0])
// fmt.Println("n",n[0][0][0])
fmt.Println("s",s)
fmt.Println("DeepEqual",reflect.DeepEqual(n[0], n[3]))
}
