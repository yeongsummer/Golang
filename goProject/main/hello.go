package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			fmt.Println(i)
			worker(i)
		}()
	}
	wg.Wait()
	fmt.Println("Finish")
}
