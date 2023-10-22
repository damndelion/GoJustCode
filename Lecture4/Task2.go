package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			mux.Lock()
			defer mux.Unlock()
			m[i] = i
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}
