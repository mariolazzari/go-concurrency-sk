# The Go concurrency survival kit

***Don't communicate by sharing memory: share memory by communicating***

## Primitives

### Goroutines

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go myFunc()
	go anotherFunc()

	// bad
	time.Sleep(time.Second)
}

func myFunc() {
	fmt.Println("myFunc")
}

func anotherFunc() {
	fmt.Println("anotherFunc")
}
```

### Waitgroups

```go
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
```

### Channels

```go
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
```

### Select

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go great(ch, "Mario")
	go great(ch, "Mariarosa")

	select {
	case msg := <-ch:
		fmt.Println("Ciao", msg)
	default:
		fmt.Println("Ciao")
	}

}

func great(ch chan string, name string) {
	ch <- name
}
```

## Patterns

### For-Select loop

```go

```
