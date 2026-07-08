package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		doWork()
	})

	// wg.Add(1)

	// go func() {
	// defer wg.Done()
	// doWork()
	// }()

	wg.Go(func() {
		doWork2()
	})

	wg.Wait()
}

func doWork() {
	fmt.Println("doWork")
}

func doWork2() {
	fmt.Println("doWork2")
}
