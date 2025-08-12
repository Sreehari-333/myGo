package main

import (
	"fmt"
	"time"
)

func main() {

	data := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println("Data Added :", i)
			data <- i
			time.Sleep(time.Second * 1)
		}

		close(data)
	}()

	for val := range data {
		fmt.Println("Data Used :", val)
		time.Sleep(time.Second * 2)
	}

}
