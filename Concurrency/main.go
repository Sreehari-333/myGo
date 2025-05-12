package main

import (
	"fmt"
	"time"
)

func printEven() {

	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}

func printOdd() {

	time.Sleep(1 * time.Second)

	for i := 0; i <= 20; i++ {

		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func main() {

	go printEven()
	go printOdd()

	time.Sleep(2 * time.Second)
}
