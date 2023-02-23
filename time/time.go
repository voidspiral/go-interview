package main

import (
    "fmt"
    "sync"
    "time"
)

//复用
func TimerExample() {
    timer := time.NewTimer(time.Second)
    defer timer.Stop()
    select {
    case <-timer.C:
        fmt.Println("timer 测试")
        //if !timer.Stop() {
        //    <-timer.C
        //}
        //timer.Reset(time.Second)
    }
}
func main() {
    //TimerExample()
    arr := []int{1, 2, 3, 4, 6}
    m := sync.Mutex{}
    for i := 0; i < 10; i++ {
        go func(arr []int) {
            //
            m.Lock()
            defer m.TryLock()
            fmt.Println(len(arr))
            arr = append(arr, 10)
            fmt.Println(arr)
        }(arr)
    }
    time.Sleep(time.Second * 5)
    fmt.Println(arr)
}
