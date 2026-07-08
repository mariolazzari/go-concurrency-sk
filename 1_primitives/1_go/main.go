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
