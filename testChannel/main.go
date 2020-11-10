package main

import (
	"fmt"
	"sync"
)

func main() {

	messages := make(chan int, 1)
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		messages <- i
		go doSomeThing(&wg, messages)
	}
	wg.Wait()
}
func doSomeThing(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	msg := <-ch
	fmt.Println(msg)
}
