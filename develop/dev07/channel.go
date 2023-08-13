package main

import (
	"fmt"
	"sync"
	"time"
)

func JoinChan(inputs ...<-chan interface{}) <-chan interface{} {
	output := make(chan interface{})

	var wg sync.WaitGroup
	for _, input := range inputs {
		wg.Add(1)

		go func(input <-chan interface{}) {
			defer wg.Done()

			for msg := range input {
				output <- msg
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func printChan(ch <-chan interface{}) {
	fmt.Println("start")
	for msg := range ch {
		fmt.Printf("new message: %v\n", msg)
	}
	fmt.Println("stop")
}

func delayChan(delay time.Duration) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		time.Sleep(delay)
		ch <- time.Now()
	}()

	return ch
}

func main() {
	ch := JoinChan(
		delayChan(1*time.Second),
		delayChan(5*time.Second),
		delayChan(10*time.Second),
	)

	printChan(ch)
}
