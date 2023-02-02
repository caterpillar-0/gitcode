package main

import (
	"fmt"
	"sync"
	"time"
)

//---------goroutine------------lock-----------------

var (
	x    int64
	lock sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("withoutlock :", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("withlock :", x)

}

// -------------------waitgroup------------------------
func hello(i int) {
	println("hello goroutine: " + fmt.Sprint(i))
}

func HelloGoRouinte() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

// ----------------main主函数入口-------------------
func main() {
	//Add()
	HelloGoRouinte()
}
