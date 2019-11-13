package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"strconv"
)

func main() {

	var (
		line string
		err error
		yyyy int
		mm int
		dd int
		h int
		m int
		s int
	)
	
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("What year?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//year
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	yyyy, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid year, %v", err)
	}

	fmt.Print("What month?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//month
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	mm, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid month, %v", err)
	}

	fmt.Print("What day?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//day
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	dd, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid day, %v", err)
	}

	fmt.Print("What hour?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//hour
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	h, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid hour, %v", err)
	}

	fmt.Print("What min?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//min
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	m, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid min, %v", err)
	}

	fmt.Print("What sec?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//sec
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	s, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid sec, %v", err)
	}

	secondsEastOfUTC := int((3 * time.Hour).Seconds())	//таймзона
	msk := time.FixedZone("MSK Time", secondsEastOfUTC)
	t := time.Date(yyyy, time.Month(mm), dd, h, m, s, 0, msk)
	
	fmt.Println(t.Unix())     // date to seconds

	fmt.Println(time.Unix(t.Unix(), 0)) //check
	
}
