package main

import (
	"sync"
	"time"
)

var l sync.RWMutex

func readAndRead() { // 可读锁内使用可读锁
	l.RLock()
	defer l.RUnlock()

	l.RLock()
	defer l.RUnlock()
}
func lockAndRead() { // 全局锁内使用可读锁
	l.Lock()
	defer l.Unlock() // 由于 defer 是栈式执行，所以这两个锁是嵌套结构

	l.RLock()
	defer l.RUnlock()
}
func main() {
	readAndRead()
	time.Sleep(5 * time.Second)
}
