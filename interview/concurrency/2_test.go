package concurrency

import (
    "fmt"
    "testing"
)

//func TestConcurrencyMap(t *testing.T) {
//	m := NewConcurrencyMap()
//	// 往 map 写 key 的协程
//	go func() {
//		// 往 map 写入数据
//		for i := 0; i < 10000; i++ {
//			m.Set(strconv.Itoa(i), i*1000)
//		}
//	}()
//	// 从 map 读取 key 的协程
//	go func() {
//		// 从 map 读取数据
//		for i := 10000; i > 0; i-- {
//			if v, ok := m.Get(strconv.Itoa(i)); ok {
//				t.Log(v.(int))
//			}
//		}
//	}()
//	// 等待两个协程执行完毕
//	time.Sleep(time.Second)
//}

func TestMap(t *testing.T) {
    m := ConcurrentMap{items: make(map[string]interface{})}

    go func() {
        m.Set("key", "value")
    }()

    go func() {
        if val, ok := m.Get("key"); ok {
            fmt.Println(val)
        }
    }()
}
