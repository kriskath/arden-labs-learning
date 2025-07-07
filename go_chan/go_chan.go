package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := range 3 {
		go func() {
			fmt.Println("goroutine:", i)
		}()
	}

	// Bad to do
	time.Sleep(10 * time.Millisecond)

	ch := make(chan int)
	go func() {
		ch <- 7 // send
	}()
	v := <-ch // receive
	fmt.Println(v)

	fmt.Println(sleepSort([]int{20, 30, 10})) // [10 20 30]
}

/*
	Algorithim

- For every value "n" in values, spin a goroutine that
  - sleeps for "n" milliseconds
  - sends "n" over channel

- collect all values from the channel to a slice and return it
*/
func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n // send
		}()
	}

	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
}

/* Channel Semantics
- send/receive to/from a channel will block until opposite operation(*)
	- guarantee of delivery
*/
