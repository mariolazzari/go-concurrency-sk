package main

import "fmt"

func main() {
	results := make(chan int)

	go doWork(results, 1)
	go doWork(results, 2)

	for range 2 {
		fmt.Println(<-results)
	}
}

func doWork(ch chan int, n int) {
	ch <- n * 2
}
