package sovle001

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

func TestLockAndUnLock(t *testing.T) {
    m := NewMyLock()
    var count int
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            m.Lock()
            count++
            m.UnLock()
        }()
    }
    wg.Wait()
    fmt.Println(count)
}

func TestTryLock(t *testing.T) {
    m := NewMyLock()
    wg := sync.WaitGroup{}
    var count int
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            if m.TryLock() {
                t.Log("抢到锁")
                count++
                defer m.UnLock()
            }
        }()
    }
    wg.Wait()
    fmt.Println(count)
}
func TestTryLockWithTimeOut(t *testing.T) {
    m := NewMyLock()
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            if m.TryLockWithTimeOut(time.Nanosecond * 50) {
                defer m.UnLock()
                t.Logf("%v'th goroutine acquired lock", i)
                //模拟一下耗时
                time.Sleep(time.Second)
            } else {
                t.Logf("%v'th goroutine trylock failed", i)
            }
        }(i)
    }
    wg.Wait()
}
