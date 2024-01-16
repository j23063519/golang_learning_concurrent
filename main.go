package main

import (
	"fmt"
	"time"
)

// // 範例二
// var start time.Time

// func init() {
// 	start = time.Now()
// }

// func s1(c chan string) {
// 	time.Sleep(1 * time.Second)
// 	c <- "Hello from s1"
// }

// func s2(c chan string) {
// 	time.Sleep(3 * time.Second)
// 	c <- "Hello from s2"
// }

// // 範例三
// var start time.Time

// func init() {
// 	start = time.Now()
// }

// // 範例七
// var start time.Time

// func init() {
// 	start = time.Now()
// }

// func s1(c chan string) {
// 	fmt.Println("s1 start", time.Since(start))
// 	time.Sleep(3 * time.Second)
// 	c <- "Hello from s1"
// }

// func s2(c chan string) {
// 	fmt.Println("s2 start", time.Since(start))
// 	time.Sleep(5 * time.Second)
// 	c <- "Hello from s2"
// }

// // 範例八
// func s() {
// 	for {
// 		fmt.Println("Hello from s")
// 		time.Sleep(400 * time.Millisecond)
// 	}
// }

// // 範例十
// var start time.Time

// func init() {
// 	start = time.Now()
// }

// func s1() {
// 	for {
// 		fmt.Println("Hello from s1 ", time.Since(start))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func s2() {
// 	for {
// 		fmt.Println("Hello from s2 ", time.Since(start))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func main() {
	fmt.Println("start")

	// select case 與 switch case 不同，select case 中 case 接收的是 channel (而不是 boolean)。當程式執行到 select 的位置，會阻塞 (blocking) 在那，直到有任何一個 case 收到 channel 傳來的資料後 (除非有用 default) 才會 unblock，因此通常會有另一個 channel 用來實作 timeout 機制
	// // 範例一
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// select {
	// case res := <-ch1: // 如果需要取用 channel 中的值
	// 	fmt.Println(res)
	// case <-ch2: // 如果不需要用 channel 中得值
	// 	fmt.Println("receive value")
	// }

	// select case 運作流程：
	// 如果所有的 case 都沒有接收到 channel 傳來的資料，那麼 select 會一直阻塞 (block)，直到任何得 case 收到資料後 (unblock) 才會繼續執行
	// 如果同一時間有多個 case 收到 channel 傳來的資料 (有多個 channel 同時 non-blocking)，那個會從所有這些 non-blocking 的 case 中隨機挑選一個，接著才繼續執行
	// // 範例二
	// // 此例子是使用 unbuffered channel，因此 channel 只要有 send 或 receive 的動作都會 block:
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go s1(ch1)
	// go s2(ch2)

	// fmt.Println("[main] select(blocking)")
	// select {
	// case res := <-ch1:
	// 	fmt.Println("[main] get response from s1", res, time.Since(start))
	// case res := <-ch2:
	// 	fmt.Println("[main] get response from s2", res, time.Since(start))
	// }
	// fmt.Println("end", time.Since(start))

	// // 上面的程式使用的是 unbuffered channel，所以對該 channel 任何的 send 或 receive 都會出現阻塞。我們可以使用 buffered channel 來模擬實際上 web service 處理回應的情況：
	// // 由於 buffered channel 的 capacity 是 2，但傳入 channel 的 size 並沒有超過 2 (沒有 overflow)，因此程式會繼續執行不會阻塞(non-blocking)
	// // 當 buffered channel 中的有資料時，直到整個 buffer 都被清空前，從 buffered channel 讀取資料的動作都是 non-blocking 的，而且在下面的程式又讀取了一個值，因此整個 case 的操作都會是 non-blocking
	// // 由於 select 中所有的 case 都是 non-blocking 的，因此 select 會從所有的 case 中隨機挑一個加以執行
	// // 範例三
	// ch1 := make(chan string, 2)
	// ch2 := make(chan string, 2)

	// // buffered channel：因為 channel 中的資料沒有 overflow (>2)，所以不會阻塞
	// ch1 <- "val 1"
	// ch1 <- "val 2"
	// ch2 <- "val 1"
	// ch2 <- "val 2"

	// // buffered channel 中有資料時，讀取資料會是 non-blocking 的
	// // 由於 select 中的 case 都是 non-blocking 的，因此會隨機挑選一個執行
	// select {
	// case res := <-ch1:
	// 	fmt.Println("[main] get response from s1", res, time.Since(start))
	// case res := <-ch2:
	// 	fmt.Println("[main] get response from s2", res, time.Since(start))
	// }
	// fmt.Println("end", time.Since(start))

	// // 當 select 中的 case 同時收到 channel 的資料時，會隨機選取一個 channel：
	// // 範例四
	// ch := make(chan int, 1)

	// ch <- 1
	// select {
	// case <-ch:
	// 	fmt.Println("1")
	// case <-ch:
	// 	fmt.Println("2")
	// }

	// // default:
	// // default case 本身是非阻塞的 (non-blocking)，同時也會使得 select statment 總是變成 non-blocking，也就是說，不論是 buffered 或是 unbuffered channel 都會變成非阻塞。
	// // 當有任何資料可以從 channel 中取出時，select 就會執行該 case，但若沒有，就會直接進到 default case。
	// // 簡單來說，當 channel 本身就有值，就不會走到 default，但如果 channel 執行的當下沒有值，還需要等其他 goroutine 設值到 channel 的話，就會走到 default。
	// // 範例五
	// ch := make(chan int, 1)

	// // 如果沒有資料送進 channel，也就是註解 ch <- 1的話，就會往 default 走，但會使得 select case 不會被阻塞住，導致還沒收到 channel 的訊息前，main goroutine 就執行完畢
	// ch <- 1

	// select {
	// case <-ch:
	// 	fmt.Println("1")
	// case <-ch:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }

	// // Timeout 超時機制:使用 time.After
	// // 單純使用 default case 並不是非常有用，由時我們希望的是有 timeout 的機制，也就是超過一定時間後，並沒有收到任何回應時，才做預設的行為，這時可以使用 time.After 來完成
	// // time.After 和 time.Tick 都是會回傳 time.Time 型別的 reveive channel (<- channel)
	// // 範例六
	// ch := make(chan int)
	// select {
	// case <-ch:
	// 	fmt.Println("receive value from channel")

	// // 超過一秒沒有收到主要 channel 的 value，就會收到 time.After 送來的訊息
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("timeout after 1 second")
	// }

	// // 範例七
	// // s1 會需要花 3 秒
	// // s2 會需要花 5 秒
	// // time.After 設定 2 秒，因此在 s1 和 s2 還沒完成前就會觸發 timeout
	// ch1 := make(chan string, 1)
	// ch2 := make(chan string, 1)

	// go s1(ch1)
	// go s2(ch2)

	// select {
	// case res := <-ch1:
	// 	fmt.Println("get response from s1", res, time.Since(start))
	// case res := <-ch2:
	// 	fmt.Println("get response from s2", res, time.Since(start))
	// case <-time.After(2 * time.Second):
	// 	fmt.Println("no response received", time.Since(start))
	// }

	// // empty select
	// // 如同 for{} 迴圈可以不帶任何條件一樣，select {} 也可以不搭配 case 使用 (稱作，empty select)。
	// // 從前面的例子中可以看到，因為 select statement 會一直組塞 (blocking)，直到其中一個 case unblocks，才會繼續往後執行，但因為 empty select 中並沒有任何的 case statement，因此 main goroutine 將會永遠阻塞在那，如果沒有其他 goroutine 可以持續運行的話，最終導致 deadlock。
	// // 範例八
	// go func() {
	// 	fmt.Println("Hello, world!")
	// }()

	// // 這個 select 永遠 block 在這
	// select {}

	// // 範例九
	// // 如果在 main goroutine 使用 empty select 後，main goroutine 將會完全阻塞，需要其他的 goroutine 持續運作才不至於進入 deadlock。
	// go s()

	// // 這個 select 永遠 block 在這
	// select {}

	// // 範例十
	// // 另外透過 empty select 導致 main goroutine 阻塞的這種方式，可以在 server 啟動兩個不同的 service
	// go s1()
	// go s2()

	// // 這個 select 會永遠 block 在這，service1 和 service2 輪流輸出訊息
	// select {}

	// // 判斷是否超過 channel 的 buffer size
	// // 範例十ㄧ
	// // 建立一個只能裝 buffer size 為 1 資料
	// ch := make(chan int, 1)
	// ch <- 1

	// select {
	// case ch <- 2:
	// 	fmt.Println("ch value is", <-ch)
	// 	fmt.Println("ch value is", <-ch)
	// default:
	// 	// ch 中的內容超過 1 時，但若把 channel buffer size 的容量改成 2，就不會走到 default
	// 	fmt.Println("blocking")
	// }

	// // 使用 for + select 讀取多個 channel 的 value
	// // 範例十二
	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)

	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Println("tick")
	// 	case <-boom:
	// 		fmt.Println("boom")
	// 		return // 如果沒有 return 的話程式將不會結束
	// 	default:
	// 		fmt.Println("default")
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// }

	// 範例十三
	ch1 := make(chan string)
	ch2 := make(chan int, 1)

	defer func() {
		fmt.Println("---- In defer ----")
		close(ch1)
		close(ch2)
	}()

	i := 0

	// 建立一個 goroutine
	go func() {
		fmt.Println("In Go Routine")

		// 透過 for loop 來不停監控不同 channels 傳回來的資料
	LOOP:
		for {
			// 透過 sleep 讓他每 500 毫秒檢查有無 channel 傳訊息進來
			time.Sleep(500 * time.Millisecond)
			i++
			fmt.Printf("In Go Routine, i: %v, time: %v \n", i, time.Now().Unix())

			// 透過 select 判斷不同 channel 傳入的資料
			select {
			case m := <-ch1: // 收到 ch1 傳入資料時就 break
				fmt.Printf("In Go Routine, get message from channel 1: %v \n", m)
				break LOOP
			case m := <-ch2: // 收到 ch2 傳入資料時，輸出訊息
				fmt.Printf("In Go Routine, get message from channel 2: %v \n", m)
			default: // 如果沒有收到任何 channel 則走 default
				fmt.Println("In Go Routine to default")
			}
		}
	}()

	ch2 <- 666 // 在 sleep 前將訊息丟入 ch2

	fmt.Println("Start Sleep")

	// 雖然這裡 sleep，但 goroutine 中的 for 迴圈仍然不斷在檢查有無收到訊息
	time.Sleep(4 * time.Second)

	fmt.Println("After Sleep: send value to channel")

	// 4秒後把 "stop"傳進 ch1，for 迴圈收到訊息後 break
	ch1 <- "stop"

	fmt.Println("end")
}
