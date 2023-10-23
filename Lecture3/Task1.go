package main

import (
	"fmt"
)

func channel_num(ch chan int) {

	ch <- 239

}

func main() {
	ch := make(chan int)

	go channel_num(ch)

	i := <-ch
	j := <-ch
	fmt.Println("Value of Channel i,j =", i, j)
}
