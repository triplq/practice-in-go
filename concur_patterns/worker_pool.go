package concurpatterns

import (
	"fmt"
	"sync"
	"time"
)

func worker(in, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range in {
		time.Sleep(1 * time.Second)
		out <- (job - 3) * 2
	}
}

func Demo_pool() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
