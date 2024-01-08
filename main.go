package main

import "fmt"

func main() {
	// // 案例一
	// // 只有"寫入"資料卻沒"讀取"該 channel -> deadlock
	// all goroutines are asleep - deadlock!
	// fmt.Println("start")

	// c := make(chan string)

	// c <- "Join"

	// fmt.Println("end")

	// // 案例二
	// // 只有"讀取"資料卻沒"寫入"該 channel -> 沒事情
	// fmt.Println("start")

	// c := make(chan string)

	// go func(c chan string) {
	// 	fmt.Println("Hi " + <-c)
	// }(c)

	// fmt.Println("end")

	// =============================
	// // buffered channel
	// // 案例ㄧ
	// // 由於寫入 channel 的值並沒超出 buffered channel 的 size，因此 main goroutine 並不會阻塞，所以 print 這個 goroutine 並不會有機會印出log
	// fmt.Println("start")

	// c := make(chan string, 3)

	// go func(c chan string) {
	// 	for i := 0; i <= 3; i++ {
	// 		fmt.Println("Hi " + <-c)
	// 	}
	// }(c)

	// c <- "1"
	// c <- "2"
	// c <- "3"
	// // 若是使用 c <- "4"，因超出 buffered channel 的 size，也就是溢出 (overflow)，因此在這會被阻塞，阻塞後便有機會執行 print goroutine ，一旦 print goroutine 開始讀取 channel 的值後，它就會把該 buffer 中所有的值全部讀完
	// // c <- "4"
	// // 可若是寫入超出 buffered channel 的 size 多個，雖然會被阻塞，但是他會印出的 log 會是隨機的(3~5)
	// // c <- "5"

	// fmt.Println("end")

	// // 案例二
	// // buffered channel 即使寫值後，不用等待值被讀取，主程式就會結束
	// fmt.Println("start")

	// c := make(chan bool, 1)

	// go func() {
	// 	fmt.Printf("Hi %v\n", <-c)
	// }()

	// c <- true

	// fmt.Println("end")

	// // 案例三
	// // buffered channel 也有 length 及 capacity
	// // length: 指的是在 channel buffer 中佔了多少數量
	// // capacity: 指的是在 buffer 中實際的 size
	// fmt.Println("start")

	// c := make(chan int, 3)

	// c <- 1

	// fmt.Printf("c: len: %v, cap: %v\n", len(c), cap(c))

	// fmt.Println("end")

	// 案例四
	// for range 可以讀取 close 後 buffered channel 中的值
	fmt.Println("start")

	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	// 若沒有close -> all goroutines are asleep - deadlock!
	close(c)

	for v := range c {
		fmt.Printf("v: %v\n", v)
	}

	fmt.Println("end")
}
