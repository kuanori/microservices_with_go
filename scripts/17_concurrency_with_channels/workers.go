package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// worker pool pattern
// это паттерн, где создаётся фикс. число работников (горутин), которые
// обрабатывают задачи из общей очереди
// Плюсы: 2. Предотвращение "наводнения" системы - Работники автоматически забирают задачи по мере освобождения
// Если придут 10,000 задач, система не создаст 10,000 горутин, а обработает их постепенно.
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
