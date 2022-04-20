package main

import (
	"fmt"
	"sync"
)

func main() {
	Solution()
}

func Solution() {
	wg := sync.WaitGroup{}

	ch := make(chan int)

	fmt.Println(ch)

	wg.Add(2)
	go func() {
		i := 1
		for ; i <= 10; i += 2 {
			fmt.Printf("curr_a: %d\n", i)
			ch <- 1
			<-ch
		}
		wg.Done()
	}()
	go func() {
		i := 2
		for ; i <= 10; i += 2 {
			<-ch
			fmt.Printf("curr_b: %d\n", i)
			ch <- 1
		}
		wg.Done()
	}()
	wg.Wait()

	close(ch)
}
