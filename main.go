package main

import (
	"fmt"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	fmt.Println("start")

	go printHello()

	// answer
	// time.Sleep(10 * time.Millisecond)
	fmt.Println("end")
}
