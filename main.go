package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// thread 1
func mainl() {
	ch := make(chan int)
	qtsWorkers := 1000

	//inicializa workers
	for i := 0; i < qtsWorkers; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 1000; i++ {
		ch <- i
	}
}
