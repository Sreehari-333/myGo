package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {

	res, err := divide(10, 10)

	if err != nil {
		fmt.Println("Error caught : ", err)
	} else {
		fmt.Println("Result : ", res)
	}

}
