package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			i := <-ch1
			fmt.Println("func", i)
			ch2 <- i
		}
	}()

	n := 1

LOOP:
	for {
		select {
		case ch1 <- n:
			fmt.Println("case1", n)
			n++
		case n := <-ch2:
			fmt.Println("case2", n)
		default:
			// fmt.Println("default", n)
			if n > 3 {
				break LOOP
			}
		}
	}
}
