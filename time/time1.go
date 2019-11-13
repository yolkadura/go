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
		sec int
	)

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("How much sec?: ")									//новый скан какой-то
	_, line, err = input.Scan(), input.Text(), input.Err()		//sec to date
	if err != nil {
		fmt.Errorf("can't scan, %v", err)
	}

	sec, err = strconv.Atoi(line)
	if err != nil {
		fmt.Errorf("invalid sec, %v", err)
	}

	fmt.Println(time.Unix(int64(sec), 0))     // seconds

}
