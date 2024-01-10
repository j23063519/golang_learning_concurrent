package main

import (
	"fmt"
)

// // 範例一
// func forloop(c chan int) {
// 	// 把 0 ~ 9 寫入 channel 後便把 channel 關閉
// 	for i := 0; i < 9; i++ {
// 		c <- i
// 	}
// 	close(c)
// }

// // 範例二
// func forrange(c chan int) {
// 	// 把 0 ~ 9 寫入 channel 後便把 channel 關閉
// 	for i := 0; i <= 9; i++ {
// 		c <- i
// 	}

// 	// 若沒有 close 會 deadlock (for range)
// 	close(c)
// }

// 範例三
func forrange2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	fmt.Println("start")

	// // 範例一
	// // for loop 搭配 close:需要手動 break 迴圈
	// // 單純透過 for 迴圈需要自行判斷 channel 是否已經 close，如果是的話需要自行使用 break 把 for loop 終止
	// c := make(chan int)

	// go forloop(c)

	// // 監聽 channel 的值：週期性的 block/unblock main goroutine 直到 squares goroutine close
	// for {
	// 	val, ok := <-c
	// 	if !ok {
	// 		fmt.Println(val, ok, "<- loop broke case channel closed")
	// 		break
	// 	} else {
	// 		fmt.Println(val, ok)
	// 	}
	// }

	// // 範例二
	// // for range 搭配 close:會自動關閉迴圈
	// // 在 go 中有 for range loop ，只要該 channel 被關閉後，loop 則會自動終止
	// // 需特別留意，如果是在 main goroutine 中使用 for val := range channel {}，最後 channel 沒有被 close 的話程式會 deadlock。但如果是在其他的 goroutine 中使用，即使沒有 close 也不會 deadlock，但為了不必要的 bug 產生，一般都還是將其關閉

	// c := make(chan int)

	// go forrange(c)

	// // 使用 for range 的寫法，一旦 channel close，loop 會自動 break
	// for v := range c {
	// 	fmt.Println(v)
	// }

	// 範例三
	// for range 其他範例
	c := make(chan int, 10)

	go forrange2(cap(c), c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("end")
}
