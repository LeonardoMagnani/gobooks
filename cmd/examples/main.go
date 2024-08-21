package main

import (
	"fmt"
	"time"
)

func counter(n int) {
	for i := range n {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Wow, GoRountines! :o"
	}()

	msg := <-ch

	fmt.Println(msg)
}
