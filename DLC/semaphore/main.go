// Вы разрабатываете систему для загрузки файлов с удаленного сервера.
// Каждый файл загружается через удаленное соединение.
// Однако, сервер ограничивает количество одновременно активных соединений,
// и ваше приложение не должго превышать этот лимит.
// В программе нельзя просто запускать фиксированное количество горутин.
// Нужно обеспечить динамически изменяемое ограничение
// количество одновременно работающих соединений.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const (
		goroutinesLimit = 3
	)

	files := []string{"file1", "file2", "file3", "file4", "file5", "file6", "file7", "file8", "file9", "file10"}

	semaphore := make(chan struct{}, goroutinesLimit)

	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, file := range files {
		fmt.Println(runtime.NumGoroutine())
		semaphore <- struct{}{} // забивать канал надо именно тут, а не в самой горутине.

		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			downloadFile(file)
		}()
	}
	// fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func downloadFile(fileName string) {
	fmt.Printf("downloading %s\n", fileName)
	time.Sleep(time.Second)
	fmt.Printf("downloaded %s\n", fileName)
}
