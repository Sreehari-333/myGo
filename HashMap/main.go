package main

import "fmt"

func main() {

	myMap := make(map[int]string)
	names := []string{"Nisa", "Niza", "Nizzz"}

	for i, val := range names {
		myMap[i+1] = val
	}

	for key, val := range myMap {
		fmt.Printf("%d -> %s \n", key, val)
	}

	fmt.Println(myMap)

}
