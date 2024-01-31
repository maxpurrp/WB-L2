package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

var wg = sync.WaitGroup{}

func or(channels ...<-chan interface{}) <-chan interface{} {
	// create a new channel to merge channels
	singleChan := make(chan interface{})

	// use WaitGroup to synchronize the goroutines
	// and ensure they all finish before closing the merged channel
	wg.Add(len(channels))

	// iterate over each input channel
	for _, channel := range channels {

		go func(ch <-chan interface{}) {
			// copy values from the input channel to the merged channel
			for v := range ch {
				singleChan <- v
			}

			fmt.Println("Done channel closed")
			wg.Done()
		}(channel)

	}
	// waiting all goroutines to finish cope values before closing channel
	wg.Wait()
	close(singleChan)

	return singleChan
}

// function that returns a channel which
// sends a single value after a specified duration
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(4*time.Second),
		sig(2*time.Second),
		sig(6*time.Second),
		sig(4*time.Second),
		sig(7*time.Second),
		sig(5*time.Second),
		sig(6*time.Second),
		sig(2*time.Second),
	)
	wg.Wait()
	fmt.Printf("Done after %v \n", time.Since(start))
}
