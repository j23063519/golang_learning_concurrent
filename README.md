# golang_learning_concurrent

在這個範例中將會發現為何沒有印出 Hello，原因是因爲 main 已經執行完畢，且不會等待其他 goroutine，但是我們知道 goroutine 被阻塞時，就會把控制全轉交給其他 goroutine，所以我們可以用 time.Sleep() 來阻塞 main 這個 goroutine

# reference
https://pjchender.dev/golang/goroutine-channels-concurrency/