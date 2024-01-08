# golang_learning_concurrent

在這個範例將會介紹 deadlock 及 buffered channel。
從案例一及案例二中可以了解到，何時會發生 deadlock 何時看似會發生，實際上並不會。
在 buffered channel 的 四個案例中，知道何時可以阻塞後執行，為何 goroutine 沒執行，以及 buffered channel 的相關資訊len及cap，以及 close channel 仍然可以用 for range 印出


# reference
https://pjchender.dev/golang/goroutine-channels-concurrency/