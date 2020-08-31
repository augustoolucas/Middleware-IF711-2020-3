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

func producer(ch chan<- int, n int, done chan<- bool) {
	for i := 0; i < countConsumers; i++ {
		ch <- n
		fmt.Println("Produced: ", n)
	}
	done <- true
	//close(ch)
}

func consumer(id int, ch <-chan int, done chan<- bool, wg *sync.WaitGroup) {
	fmt.Println("Consumer ID:", id)
	for {
		select {
		case n := <-ch:
			fmt.Println("Consumer", id, "received:", n)
			wg.Done()
			break
		default:
			fmt.Println("Consumer", id, "waiting producer")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//countConsumers = rand.Intn(MaxConsumers) + 1
	countConsumers = 5
	fmt.Println(countConsumers, " consumers")

	ch := make(chan int, countConsumers)
	done := make(chan bool)

	rand.Seed(time.Now().UnixNano())

	go producer(ch, rand.Intn(100), done)
	var wg sync.WaitGroup
	wg.Add(countConsumers)
	for i := 0; i < countConsumers; i++ {
		go consumer(i, ch, done, &wg)
	}
	wg.Wait()
	<-done
}
