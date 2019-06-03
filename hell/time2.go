package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		yyyy int
		mm int
		dd int
		h int
		m int
		s int
	)
    
  fmt.Print("What year?: ")
	fmt.Scanf("%d", &yyyy)
	fmt.Print("What month?: ")
	fmt.Scanf("%d", &mm)
	fmt.Print("What day?: ")
	fmt.Scanf("%d", &dd)
	fmt.Print("What hour?: ")
	fmt.Scanf("%d", &h)
	fmt.Print("What minute?: ")
	fmt.Scanf("%d", &m)
	fmt.Print("What sec?: ")
  fmt.Scanf("%d", &s)



	//fmt.Println(time.Unix(1527789600, 0))     // seconds to date

	secondsEastOfUTC := int((3 * time.Hour).Seconds())	
	msk := time.FixedZone("MSK Time", secondsEastOfUTC)
	//t := time.Date(2018, 6, 1, 0, 0, 0, 0, msk)
	t := time.Date(yyyy, mm, dd, h, m, s, 0, msk)
	fmt.Println(t.Unix())     // date to seconds

	fmt.Println(time.Unix(t.Unix(), 0)) //check
	
}
