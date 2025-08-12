package main

import "fmt"

func main() {

	students := make(map[string]int)

	students["Sreehari"] = 100
	students["Roshith"] = 99
	students["Bijoy"] = 95

	for key, value := range students {
		fmt.Printf("%s - %d \n", key, value)
	}

}
