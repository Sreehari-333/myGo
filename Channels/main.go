package main

import "fmt"

func passHundred(ch chan string) {

	ch <- "First Message"
	ch <- "Second Messaage"
	ch <- "Third Message"

}
func main() {

	ch := make(chan string)

	go passHundred(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
