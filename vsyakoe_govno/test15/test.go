package main

import (
	"fmt"
)
const (
	offset = 97
)
var (
	// words = []string {"gin", "zen", "gig", "msg"}
	// words = []string {"gin", "zen", "gig", "omg"}
	// words = []string {"zocd","gjkl","hzqk","hzgq","gjkl"}
	// words = []string {"zocd","gjkl","hzqk","hzgq","gjkl","gig","omg"}
	words = []string {"gin","gin"}
	// words = []string {"yxmine","yxzd","eljys","uiaopi","pwlk"}
	alphabet = []string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	x [][]byte
	k string
	y []byte
	l [][]byte
	n [][]byte
	temp []byte
	s int = 0
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


	for i := range n {
		fmt.Println("n",n[i])
	}

	// switch len(n) {
	// case 0, 1 : 
	// 	s = len(n)
	// default:
	// 	for i := 0; i < len(n)-1; i++ {
	// 		for j := i+1; j < len(n); j++ {
	// 			if qqq(n[i], n[j]) == true {
	// 				n = append(n[:i], n[i+1:]...)
					
	// 			}
	// 		}
	// 	}
	// 	s = len(n)
	// }

	switch len(n) {
		case 0, 1 : 
			s = len(n)
		default:
			for i := 0; i < len(n); i++ {
				for j := i+1; j < len(n); j++ {
					if qqq(n[i], n[j]) == true {
						n[i] = nil
						
					}
				}
			}
			s = len(n)
			for i := range n {
				if n[i] == nil {
					s--
				}
			}
	}
	


	fmt.Println("s",s)

	for i := range n {
		fmt.Println("n",n[i])
	}


	// fmt.Println("x", x)
	// fmt.Println("l",l)
	// fmt.Println("n",n)
	


}
