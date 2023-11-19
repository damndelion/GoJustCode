package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			ch1 <- num
		}
		close(ch1)
	}()
	go func() {
		for _, num := range []int{4, 5, 6} {
			ch2 <- num
		}
		close(ch2)
	}()
	go func() {
		for _, num := range []int{10, 20, 30} {
			ch3 <- num
		}
		close(ch3)
	}()

	mergedCh := joinChannels(ch1, ch2, ch3)
	for num := range mergedCh {
		fmt.Println(num)
	}
}

func joinChannels(ch ...<-chan int) <-chan int {

	mergedCh := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(len(ch))
	go func() {
		for _, num := range ch {
			go func(num <-chan int) {
				defer wg.Done()
				for val := range num {
					mergedCh <- val
				}
			}(num)

		}
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}
