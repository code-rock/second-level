package main

import (
	"fmt"
	"sync"
	"time"
)

type TChannel <-chan interface{}

func uniteСhannels(channels ...TChannel) TChannel {
	united := make(chan interface{})
	var wg sync.WaitGroup

	for _, channel := range channels {
		wg.Add(1)
		go func(channel TChannel) {
			for message := range channel {
				united <- message
			}
			wg.Done()
		}(channel)
	}
	wg.Wait()
	defer close(united)
	return united
}

func main() {
	sig := func(after time.Duration) TChannel {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	united := uniteСhannels(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v %v \n", time.Since(start), united)
}
