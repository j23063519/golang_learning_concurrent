# golang_learning_concurrent

在這個範例中，可以使用”匿名函式(anonymous)“的goroutine，也了解到 channel 得基本使用。在這個範例中，使用的是 unbuffered channel，原本在上一個案例中，需要 time.Sleep 才能印出 log，這時使用 unbufferd channel，這樣 main 就會等到 channel 值被讀出來才會結束。

# reference
https://pjchender.dev/golang/goroutine-channels-concurrency/