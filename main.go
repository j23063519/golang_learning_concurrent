package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	// 注意：
	// var c chan string 會是 nil
	// var c chan string
	// c = make(chan string)
	// 範例ㄧ
	// var c chan string = make(chan string)
	// 範例二
	// c := make(chan string, 1) buffered channel
	// 範例三
	c := make(chan string)

	// error: all goroutines are asleep - deadlock!
	// c <- "John"

	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "John"

	fmt.Println("end")
}
