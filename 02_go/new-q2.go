package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	MAX_CONSUMERS   = 10
	count_consumers = 0
)

func producer(ch chan<- int, n int, done chan<- bool) {
	for i := 0; i < count_consumers; i++ {
		ch <- n
		fmt.Println("Produced: ", n)
	}
	done <- true
	close(ch)
}

func consumer(id int, ch <-chan int, done chan<- bool) {
	fmt.Println("Consumer ID:", id)
	for {
		select {
		case n := <-ch:
			fmt.Println("Consumer", id, "received:", n)
			done <- true
			break
		default:
			fmt.Println("Consumer", id, "waiting producer")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	count_consumers = rand.Intn(MAX_CONSUMERS) + 1
	count_consumers = 5
	fmt.Println(count_consumers, " consumers")

	ch := make(chan int, count_consumers)
	done := make(chan bool)

	rand.Seed(time.Now().UnixNano())

	go producer(ch, rand.Intn(100), done)

	for i := 0; i <= count_consumers; i++ {
		go consumer(i, ch, done)
	}
	<-done
}
