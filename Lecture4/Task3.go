package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	rwm := sync.RWMutex{}

	for i := 0; i < 100; i++ {
		rwm.Lock()
		m[i] = i
		rwm.Unlock()
	}
	for i := 0; i < 100; i++ {
		rwm.RLock()
		fmt.Println(m[i])
		rwm.RUnlock()
	}

	fmt.Println(len(m))
}
