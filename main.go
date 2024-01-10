package main

import (
	"fmt"
)

// // 範例二
// func greet(read <-chan string) {
// 	fmt.Println("Hello " + <-read + "!")
// }

// // 範例三
// func square(c chan int) {
// 	fmt.Println("[square] wait for testNum")
// 	num := <-c
// 	fmt.Println("[square] sent square to squateChan (blocking)")
// 	c <- num * num // blocking
// }

// func cube(c chan int) {
// 	fmt.Println("[cube] wait for testNum (blocking)")
// 	num := <-c
// 	fmt.Println("[cube] sent square to cubeChan")
// 	c <- num * num * num // blocking
// }

// 範例四
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("start")

	// // 範例一
	// // Unidirectional Channels 單向通道
	// // channel 除了能寫入(write)及讀取(read)，也能有 只能寫入(write-only)或只能讀取(receive-only) 的 channel
	// read := make(<-chan int)  // receive only channel
	// write := make(chan<- int) // write only channel

	// fmt.Printf("read: %T\n", read)
	// fmt.Printf("write: %T\n", write)

	// // 範例二
	// // 透過 unidirectional channels 增加型別的安全性 (type-safety)
	// // 如果希望在某 goroutine 中只能從 channel 讀取資料，但在 main goroutine 中可以對 channel 進行寫入和讀取時，可以透過 go 提供的語法將 bi-directional channel 轉成 unidirectional channel
	// c := make(chan string)

	// go greet(c)

	// c <- "John"

	// // 範例三
	// // Multiple Channel
	// // 當前的 goroutine 阻塞時，就會切換到其他 goroutine
	// squareChan := make(chan int)
	// cubeChan := make(chan int)

	// go square(squareChan)
	// go cube(cubeChan)

	// testNum := 3

	// fmt.Println("[main] sent testNum to squareChan (blocking)")
	// squareChan <- testNum

	// fmt.Println("[main] resuming")

	// fmt.Println("[main] sent testNum to cubeChan")
	// cubeChan <- testNum

	// fmt.Println("[main] resuming")

	// fmt.Println("[main] reading from channels (blocking)")
	// squareVal, cubeVal := <-squareChan, <-cubeChan

	// fmt.Println(squareVal, cubeVal)

	// 範例四
	// first-class channel
	// 在 golang 中 channel 是 first-class values，和其他型別一樣，可以被當作是 struct 中的值、function 的參數、回傳的值
	// 下述例子，make(chan chan string)，表示這個 channel 可以傳送和接收另一個(可以傳送和接收 string) 的 channel

	// a channel of data type channel of data type string
	// 建立一個 channel 可以讀寫另一個（可以讀寫 string）的 channel
	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greet(c)
	c <- "John"

	fmt.Println("end")
}
