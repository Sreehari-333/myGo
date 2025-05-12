package main

import "fmt"

func changeValue(num *int) {

	*num = 20
}
func main() {

	num := 10

	changeValue(&num)

	fmt.Println(num)

}
