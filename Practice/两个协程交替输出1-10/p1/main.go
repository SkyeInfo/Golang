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
	ch1 := make(chan int)
	//ch2 := make(chan string)
	ch3 := ch

	if ch == ch1 {
		fmt.Println("ch == ch1")
	}

	if ch3 == ch {
		fmt.Println("ch3 == ch")
	}

	fmt.Println(ch)

	wg.Add(2)
	go func() {
		i := 1
		for ; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Printf("curr_a: %d\n", i)
			}

			ch <- 1
		}
		wg.Done()
	}()
	go func() {
		i := 1
		for ; i <= 10; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Printf("curr_b: %d\n", i)
			}
		}
		wg.Done()
	}()
	wg.Wait()

	close(ch)
}
