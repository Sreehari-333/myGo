package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter)
}
