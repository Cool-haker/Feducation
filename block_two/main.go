package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func totalPlus(tc *int, lock *sync.Mutex) {
	lock.Lock()
	*tc++
	lock.Unlock()
}

func anyGoroutines(lock *sync.Mutex, n int, f func(tc *int, lock *sync.Mutex), counter *int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			f(counter, lock)
		}()
	}
	wg.Wait()
}

// 1 Задача
// func main() {
// 	var lock sync.Mutex
// 	totalCounter := 0
// 	fmt.Println(totalCounter)
// 	anyGoroutines(&lock, 10, totalPlus, &totalCounter)
// 	fmt.Println(totalCounter)
// }

func totalPlusPlus(tc *atomic.Int32) {
	tc.Add(1)
}

func anyGoroutinchik(n int, f func(tc *atomic.Int32), counter *atomic.Int32) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			f(counter)
		}()
	}
	wg.Wait()
}

// // 2 Задача
// func main() {
// 	var totalCounter atomic.Int32
// 	fmt.Println(totalCounter.Load())
// 	anyGoroutinchik(10, totalPlusPlus, &totalCounter)
// 	fmt.Println(totalCounter.Load())
// }

func sumPart(nums []int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for _, val := range nums {
		result += val
	}
	resultCh <- result
}

func main() {
	const parts = 10
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	resultCh := make(chan int, parts)

	var wg sync.WaitGroup
	for i := 0; i < parts; i++ {
		wg.Add(1)
		go sumPart(nums[i*100:(i+1)*100], resultCh, &wg)
	}

	wg.Wait()

	result := 0
	for i := 0; i < parts; i++ {
		result += <-resultCh
	}

	fmt.Println(result)
}
