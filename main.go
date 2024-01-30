package main

import (
	"fmt"
	"net/http"
)

// // 範例一
// var start time.Time

// func init() {
// 	start = time.Now()
// }

// func service(wg *sync.WaitGroup, instance int) {
// 	time.Sleep(time.Duration(instance) * 500 * time.Millisecond)
// 	fmt.Println("Service called on instance: ", instance, time.Since(start))
// 	wg.Done()
// }

// // 範例二
// func notifying(wg *sync.WaitGroup, s string) {
// 	fmt.Printf("Starting to notifying %s... \n", s)
// 	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
// 	fmt.Printf("Finish notifying %s \n", s)
// 	wg.Done()
// }

// func notify(services ...string) {
// 	var wg sync.WaitGroup

// 	for _, service := range services {
// 		wg.Add(1) // 添加 counter 次數
// 		go notifying(&wg, service)
// 	}

// 	wg.Wait()

// 	fmt.Println("All service notified")
// }

// // 範例三
// func notifying(res chan string, s string) {
// 	fmt.Printf("Starting to notifying %s... \n", s)
// 	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
// 	res <- fmt.Sprintf("Finish notifying %s", s)
// }

// func notify(services ...string) {
// 	res := make(chan string)
// 	var count int = 0

// 	for _, service := range services {
// 		count++
// 		go notifying(res, service)
// 	}

// 	for i := 0; i < count; i++ {
// 		fmt.Println(<-res)
// 	}

// 	fmt.Println("All service notified")
// }

// // 範例四
// // STEP 3 : 在 worker goroutines 中會做相同的工作
// // tasks is receive only channel
// // results is send only channel
// func sqrWorker(tasks <-chan int, results chan<- int, instance int) {
// 	// 一旦收到 tasks channel 傳來資料，就可以動工並傳回結果
// 	for v := range tasks {
// 		time.Sleep(500 * time.Millisecond) // 模擬組塞的任務
// 		fmt.Printf("[worker %v] Sending result of task %v \n", instance, v)
// 		results <- v * v
// 	}
// }

// // 範例五
// // STEP 3 : 在 worker goroutines 中會做相同的工作
// // tasks is receive only channel
// // results is send only channel
// // defer 先進後出，所以會在最後執行 [減少一次] WaitGroup counter
// func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
// 	defer wg.Done()

// 	// 一旦收到 tasks channel 傳來資料，就可以動工並回傳結果
// 	for v := range tasks {
// 		time.Sleep(500 * time.Millisecond) // 模擬會組塞的任務
// 		fmt.Printf("[worker %v] Sending result of task %v \n", instance, v)
// 		results <- v * v
// 	}
// }

// // 範例六
// var x int

// func worker(wg *sync.WaitGroup) {
// 	x++
// 	wg.Done()
// }

// // 範例七
// var x int

// func worker(wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	x++
// 	m.Unlock()
// 	wg.Done()
// }

// // 範例八
// // fib 會回傳 read-only channel
// func fib(length int) <-chan int {
// 	c := make(chan int, length)

// 	// run generation concurrently
// 	go func() {
// 		for i, j := 0, 1; i < length; i, j = i+j, i {
// 			c <- i
// 		}
// 		close(c)
// 	}()

// 	return c
// }

// // 範例九
// func worker(wg *sync.WaitGroup, c chan<- int, i int) {
// 	fmt.Println("[worker] start i:", i)
// 	time.Sleep(time.Second * 1)
// 	defer wg.Done()
// 	c <- i
// 	fmt.Println("[worker] stop i:", i)
// }

// // 範例十
// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}

// 	// STEP 3 : 把加總後的值丟回 channel
// 	c <- sum
// }

// 範例十一
func checkLink(link string) {
	if _, err := http.Get(link); err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

func main() {
	fmt.Println("start")

	// sync.WaitGroup
	// var wg sync.WaitGroup 可以建立 waitgroup，預設 counter 是 0
	// wg.Add(delta int) 增加要等待的次數 (increment counter)，也是可以負值，通常就是要等待完成的 goroutine 數目
	// wg.Done() 會把要等待的次數減 1 (decrement counter)，可以使用 defer wg.Done()
	// wg.Wait() 會阻塞在這，直到 counter 歸零，也就是所有 WaitGroup 都呼叫過 done 後才往後執行
	// // 範例一
	// fmt.Println("[main]: ", time.Since(start))
	// var wg sync.WaitGroup // 建立 waitgroup (empty struct)
	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go service(&wg, i)
	// }

	// wg.Wait() // 阻塞到 counter 為 0

	// // 範例二
	// // 這裡的 wg 需要把 pointer 傳進去 goroutine 中，如果不是傳 pointer 進去而是傳 value 的話，將沒辦法有效把 main goroutine 中的 waitGroup 的 counter 減 1。
	// notify("Service-1", "Service-2", "Service-3")

	// // 範例三
	// // 如果我們需要使用到 goroutine 中回傳的資料，那個應該要使用 channel 而不是 waitGroup
	// notify("Service-1", "Service-2", "Service-3")

	// // 範例四
	// // worker pool
	// // worker pool 指的是有許多的 goroutine 同步的進行一個工作。要建立 worker pool，會先建立許多的 worker goroutine，這些 goroutine 中會:
	// // 。進行相同的 job
	// // 。有兩個 channel，一個用來接受任務 (task channel)，一個用來回傳結果 (result channel)
	// // 。都等待 task channel 傳來要進行的 tasks
	// // 。一旦收到 tasks 就可以做事並透過 result channel 回傳結果
	// // STEP 1 : 建立兩個 channel，一個用來傳送 tasks，一個用來接收 results
	// tasks := make(chan int, 10)
	// results := make(chan int, 10)

	// // STEP 2 : 啟動三個不同的 worker goroutines
	// for i := 0; i < 3; i++ {
	// 	go sqrWorker(tasks, results, i)
	// }

	// // STEP 4 : 發送 5 個不同的任務
	// for i := 0; i < 5; i++ {
	// 	tasks <- i // non-blocking
	// }

	// fmt.Println("[main] Wrote 5 tasks")

	// // STEP 5 : 發送完任務把 channel 關閉 (非必要，但可減少 bug)
	// close(tasks)

	// // STEP 6 : 等待各個 worker 從 result channel 回傳結果
	// for i := 0; i < 5; i++ {
	// 	result := <-results
	// 	fmt.Println("[main] Result ", i, ":", result)
	// }

	// // 範例五
	// // WorkerGroup 搭配 WaitGroup
	// // 有些時候，我們希望所有 tasks 都執行完後才讓 main goroutine 繼續往後做，這時可以搭配 WaitGroup :
	// var wg sync.WaitGroup

	// // STEP 1 : 建立兩個 channel，一個用來傳送 tasks，一個用來接收 results
	// tasks := make(chan int, 10)
	// results := make(chan int, 10)

	// // STEP 2 : 啟動三個不同的 worker goroutines 並添加三次 WaitGroup counter
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go sqrWorker(&wg, tasks, results, i)
	// }

	// // STEP 4 : 發送 5 個不同的任務
	// for i := 0; i < 5; i++ {
	// 	tasks <- i // non-blocking (因為buffered channel 的 capacity 是 10)
	// }

	// fmt.Println("[main] Wrote 5 tasks")

	// // STEP 5 : 發送完任務把 channel 關閉，有用 waitGroup 的話，必須使用 close
	// close(tasks)

	// // STEP 6 :必須等到所有 worker goroutine 把所有 tasks 都做完後才繼續往後
	// wg.Wait()

	// // STEP 7 : 等待各個 worker 從 result channel 回傳結果
	// for i := 0; i < 5; i++ {
	// 	result := <-results // blocking (因為 buffer 是空的)
	// 	fmt.Println("[main] Result", i, ":", result)
	// }

	// // 結論：
	// // 這時會等到所有的 worker 都完工下班後，才開始輸出計算好的結果。搭配 WaitGroup 的好處是可以等到所有的 worker 都完工後讓程式繼續，但相對的會需要花更長的時間在等待所有人完工

	// // mutex
	// // 在 goroutines 中，由於有獨立的 stack，因此並不會在彼此之間共享資料 (此為 scope 中的變數); 然而在 heap 中的資料是會在不同的 goroutine 之間共享的 (此為 global 的變數)，在這情況下，許多的 goroutine 會試著操作相同記憶體位址的資料，導致未預期的結果。
	// // 範例六
	// // 在這個例子中可以看到，我們預期 i 的結果會是 1000，但是因為 race condition 的情況，最終的結果並不會是 1000：
	// var wg sync.WaitGroup

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go worker(&wg)
	// }

	// wg.Wait()
	// fmt.Println(x)

	// // 結論：
	// // 之所以會有這個情況發生，是因為多個 goroutine 在執行時，在為 x 賦值前，拿到同個 x，所以雖然跑了很多次 goroutine，但對於 x 來說，並沒增加
	// // 為了要避免多個 goroutine 同時取到一個 heap 中的變數，第一原則是要盡可能避免在多個 goroutine 中使用共享的資源(變數)，如果無法避免則需要用 Mutex (mutual exclusion)
	// // 簡單來說在一個時間區段內只有一個 goroutine 可以對該變數進行操作，在該變數進行操作前，會把它先上鎖，操作完後再解鎖，當一個變數被上鎖後，其他 goroutine 都不能對該變數進行讀取和寫入
	// // mutex 是 map 型別得方法，被放在 sync package 中，使用 mutex.Lock() 可以上鎖，使用 mutex.Unlock() 可以解鎖
	// // 範例七
	// var wg sync.WaitGroup
	// var m sync.Mutex

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go worker(&wg, &m)
	// }

	// wg.Wait()

	// fmt.Println(x)
	// // 補充：
	// // mutex 和 waitGroup 一樣都是把 『記憶體位址』傳入 goroutine 中使用
	// // 如同前面所說，第一原則應該是要避免 race condition，也就是不要在 goroutine 中對共用的變數進行操作，在 go CLI 中可以透過：
	// // go run -race program.go
	// // 檢查程式是否有 race condition 的情況發生

	// // concurrency pattern
	// // generator
	// // 範例八
	// for v := range fib(10) {
	// 	fmt.Println("Current fibonacci number is ", v)
	// }

	// // waitGroup 搭配 channel
	// // 範例九
	// numOfFacilities := 6
	// var wg sync.WaitGroup

	// c := make(chan int, numOfFacilities)

	// for i := 0; i < numOfFacilities; i++ {
	// 	fmt.Println("[main] add i: ", i)
	// 	wg.Add(1)
	// 	go worker(&wg, c, i)
	// }

	// wg.Wait()

	// var numbers []int
	// for i := 0; i < numOfFacilities; i++ {
	// 	numbers = append(numbers, <-c)
	// }
	// fmt.Println("[main] ---all finish---", numbers)

	// defer close(c)

	// // 範例十
	// s := []int{7, 2, 8, -9, 4, 0}

	// // STEP 1 : 建立一個 channel，該 channel 會傳出 int
	// c := make(chan int)

	// // STEP 2 : 使用 goroutine，並把 channel 當成參數傳入
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)

	// // STEP 4 : 從 channel 取得計算好的結果
	// x, y := <-c, <-c

	// // 寫在這裡的內容會在 channel 傳回結果後才會執行
	// fmt.Println(x, y, x+y)

	// 範例十一
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}

	fmt.Println("end")
}
