# golang_learning_concurrent

在這個範例中會知道如何使用 waitgroup 及 waitgroup 與 channel、workergroup或是mutex 等等的應用

WaitGroup:
WaitGroup 的用法適合在，需要將單一任務拆成多次任務，待所有任務完成後才繼續執行的情境
!! 這種作法適合用在單純等待任務完成，而不需要從 goroutine 中取得所需資料的情況，如果會需要從 goroutine 中返回資料，那麼比較好的做法是使用 channel

# reference
https://pjchender.dev/golang/goroutine-channels-concurrency/