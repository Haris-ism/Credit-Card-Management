package main

import (
	"fmt"
)

func Generator() chan int {

	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 0           // stops underlying goroutine
	fmt.Println(<-number) // error, no one is sending anymore
	// …
}
