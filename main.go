package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func send(s chan int) {
	for i := 0; i < 10; i++ {
		s <- i
	}

}
func receive(s chan int) {
	for i := 0; i < 10; i++ {
		c := <-s
		fmt.Println(c)
		wg.Done()
	}
}
func main() {

	wg.Add(10)
	sender := make(chan int)
	go send(sender)
	go receive(sender)
	wg.Wait()
}
