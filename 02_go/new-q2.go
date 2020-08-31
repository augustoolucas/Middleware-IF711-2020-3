package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	//MaxConsumers: maximo numero de consumidores
	MaxConsumers   = 10
	countConsumers = 0
)

func producer(ch chan<- int, n int) {
	for i := 0; i < countConsumers; i++ {
		ch <- n
		fmt.Println("Produced: ", n)
	}
}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	fmt.Println("Consumer ID:", id)
	for {
		select {
		case n := <-ch:
			fmt.Println("Consumer", id, "received:", n)
			wg.Done()
			return
		default:
			fmt.Println("Consumer", id, "waiting producer")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	countConsumers = rand.Intn(MaxConsumers) + 1
	fmt.Println(countConsumers, " consumers")
	time.Sleep(time.Second)

	ch := make(chan int, countConsumers)

	rand.Seed(time.Now().UnixNano())

	go producer(ch, rand.Intn(100))
	var wg sync.WaitGroup
	wg.Add(countConsumers)
	for i := 0; i < countConsumers; i++ {
		go consumer(i, ch, &wg)
	}
	wg.Wait()
}
