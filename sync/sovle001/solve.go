package sovle001

import (
	"time"
)

type ChanelImplementLock interface {
	Lock()
	UnLock()
	TryLock() bool
	TryLockWithTimeout(d time.Duration) bool
	isLocked() bool
}

const (
	LOCKED = 1 << iota
	UNLOCKED
)

type MyLock struct {
	l      chan struct{}
	status int32
}

func NewMyLock() *MyLock {
	l := make(chan struct{}, 1)
	return &MyLock{
		l: l,
	}
}
func (m *MyLock) Lock() {
	m.l <- struct{}{}
}

//// before
//func (m *MyLock) UnLock() {
//    <-m.l
//}

// fix
func (m *MyLock) UnLock() {
	select {
	case <-m.l:
	default:
		panic("unlock of unlocked mutex")
	}
}

// before
//func (m *MyLock) IsLocked() bool {
//	return len(m.l) == 1
//}

// after

////before
//func (m *MyLock) TryLock() bool {
//    return !m.IsLocked()
//}

// fix
func (m *MyLock) TryLock() bool {
	select {
	case m.l <- struct{}{}:
		return true
	default:
	}
	return false
}

// before
//func (m *MyLock) TryLockWithTimeOut(d time.Duration) bool {
//    timer := time.NewTimer(d)
//    defer timer.Stop()
//    for m.TryLock() == false {
//        select {
//        case <-timer.C:
//            return false
//        default:
//        }
//    }
//    return false
//}

// after
func (m *MyLock) TryLockWithTimeOut(d time.Duration) bool {
	select {
	case m.l <- struct{}{}:
		return true
	case <-time.After(d):
		return false
	default:
	}
	return false
}
