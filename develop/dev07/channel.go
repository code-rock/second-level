package main

import (
	"fmt"
	"time"
)

type TChannel <-chan bool

func uniteСhannels(channels ...TChannel) TChannel {
	united := make(chan bool)
	fmt.Println("start")
	for _, channel := range channels {
		fmt.Println("channel before", channel)
		go func(channel TChannel) {
			fmt.Println("channel", channel)
			for message := range channel {
				united <- message
			}
		}(channel)
	}

	go func() {
		for {
			select {
			case <-united:
				close(united)
				fmt.Println("close")

				return
			}
		}

	}()

	return united
}

func main() {
	sig := func(after time.Duration) TChannel {
		c := make(chan bool)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	united := uniteСhannels(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(10*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Println("united", united)

}
