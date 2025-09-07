package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
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

// 2 Задача
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

// 3 Задача
// func main() {
// 	const parts = 10
// 	nums := make([]int, 1000)
// 	for i := range nums {
// 		nums[i] = rand.Intn(100)
// 	}

// 	resultCh := make(chan int, parts)

// 	var wg sync.WaitGroup

// 	for i := 0; i < parts; i++ {
// 		wg.Add(1)
// 		go sumPart(nums[i*100:(i+1)*100], resultCh, &wg)
// 	}

// 	wg.Wait()

// 	result := 0
// 	for i := 0; i < parts; i++ {
// 		result += <-resultCh
// 	}

// 	fmt.Println(result)
// }

func generator(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 6; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}

// 4 Задача
// func main() {
// 	ch := make(chan int)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go generator(ch, &wg)
// 	go consumer(ch, &wg)

// 	wg.Wait()
// }

func producer(ch1 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 6; i++ {
		ch1 <- i
	}
	close(ch1)
}

func square(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch1 {
		ch2 <- val * val
	}
	close(ch2)
}

func printer(ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch2 {
		fmt.Println(val)
	}
}

// 5 Задача
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go producer(ch1, &wg)
// 	go square(ch1, ch2, &wg)
// 	go printer(ch2, &wg)
// 	wg.Wait()
// }

func worker(ctx context.Context) {
	for true {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("tick")
		}
	}
}

// 6 Задача
// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

//		go worker(ctx)
//		time.Sleep(4 * time.Second)
//	}
func void1(ctx context.Context, ch1 chan int) {
	timer := time.NewTicker(200 * time.Millisecond)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			ch1 <- 1
			return
		}
	}
}

func void2(ctx context.Context, ch2 chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(199 * time.Millisecond):
			ch2 <- 1
			return
		}
	}
}

// 7 Задача
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	go void1(ctx, ch1)
// 	go void2(ctx, ch2)

// 	select {
// 	case res := <-ch1:
// 		fmt.Printf("Горутина 1 быстрее: %d", res)
// 		cancel()
// 	case res := <-ch2:
// 		fmt.Printf("Горутина 2 быстрее: %d", res)
// 		cancel()
// 	}
// }

type SafeNumber struct {
	value int
	mu    sync.RWMutex
}

func (s *SafeNumber) Get() int {
	s.mu.RLock()
	v := s.value
	s.mu.RUnlock()
	return v
}

func (s *SafeNumber) Set(v int) {
	s.mu.Lock()
	s.value = v
	s.mu.Unlock()
}

func reader(id int, sn *SafeNumber, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			v := sn.Get()
			fmt.Printf("R%d - %d\n", id, v)
		}
	}
}

func writer(sn *SafeNumber, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	tick := time.NewTicker(80 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			sn.Set(sn.value + 1)
		}
	}
}

// 8 Задача
// func main() {
// 	sn := SafeNumber{}
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()
// 	var wg sync.WaitGroup

// 	wg.Add(3)
// 	go reader(1, &sn, ctx, &wg)
// 	go reader(2, &sn, ctx, &wg)
// 	go writer(&sn, ctx, &wg)
// 	wg.Wait()
// }

func increment(n *atomic.Int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		n.Add(1)
	}
}

// 9 Задача v1 более простая
// func main() {
// 	var total atomic.Int32
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go increment(&total, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(total.Load())
// }

type Counter struct {
	count atomic.Int32
}

func (c *Counter) increment() {
	c.count.Add(1)
}

func incrementer(count *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		count.increment()
	}
}

// 9 Задача v2 более прпавильная потому что есть инкапсуляция, а на проде так и надо
// func main() {
// 	c := Counter{}
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go incrementer(&c, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(c.count.Load())
// }

func worker1(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range jobs {
		fmt.Printf("Worker%d is processing number %d\n", id, val)
	}
}

// 10 Задача
// func main() {
// 	const (
// 		numJobs    = 20
// 		numWorkers = 3
// 	)

// 	jobs := make(chan int, numJobs)
// 	var wg sync.WaitGroup

// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go worker1(i, jobs, &wg)
// 	}

// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// 	wg.Wait()
// }

func worker2(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range jobs {
		results <- val * val
	}
}

func main() {
	const (
		numJobs    = 20
		numWorkers = 3
	)

	results := make(chan int, numJobs-10)
	jobs := make(chan int, numJobs-10)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker2(jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()
	var result []int
	for val := range results {
		result = append(result, val)
	}

	fmt.Println(result)
}
