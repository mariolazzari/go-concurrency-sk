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
