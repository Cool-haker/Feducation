// Вы разрабатываете сервис, который обрабатывает изображения.
// Каждое изображение проходит дорогостоящую обработку, например, наложение подяного знака.
// Поскольку обработка кажлдого изображения занимает значительное время,
// необходимо обрабатывать их параллельно, чтобы ускорить процесс.
// Однако, чтобы избежать излишней нагрузки на систему, вы хотите ограничить количество
// одновременно работающих горутин.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	const (
		numWorkers = 5
		numTasks   = 10
	)

	taskCh := make(chan Task, numTasks)
	resCh := make(chan string, numTasks)

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			RunWorker(ctx, int64(i), taskCh, resCh)
		}()
	}

	go func() {
		for i := 0; i < numTasks; i++ {
			taskCh <- Task{ID: int64(i), Filename: fmt.Sprintf("File_%d.jpg", i)}
		}

		close(taskCh)
	}()

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for res := range resCh {
		fmt.Println(res)
	}

	fmt.Println("All tasks are done")
}

type Task struct {
	ID       int64
	Filename string
}

func processImage(task Task) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("Файл %v обработан (File ID — %d)", task.Filename, task.ID)
}

func RunWorker(ctx context.Context, id int64, taskCh <-chan Task, resCh chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task canceled for context")
			return
		case task := <-taskCh:
			fmt.Printf("Worker %d started task %s\n", id, task.Filename)
			resCh <- processImage(task)
			fmt.Printf("Worker %d finished task %s\n", id, task.Filename)
		}
	}
}
