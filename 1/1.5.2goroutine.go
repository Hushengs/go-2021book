package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// var total struct {
// 	sync.Mutex
// 	value int
// }

// //sync.Mutex来实现
// func worker(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i <= 100; i++ {
// 		total.Lock()
// 		total.value += 1
// 		total.Unlock()
// 	}
// }

var total uint64

//sync/atomic包来实现
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

func main() {
	fmt.Println(total)
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	// fmt.Println(total.value)
	fmt.Println(total)
}
