package concurpatterns

import (
	"fmt"
	"sync"
	"time"
)

func Gener() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 12; i++ {
			ch <- i * 2
		}
		close(ch)
	}()

	return ch
}

func Fanin(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				ch <- s
			case s := <-input2:
				ch <- s
			}
		}
	}()

	return ch
}

func work(wg *sync.WaitGroup, ch <-chan int) {
	wg.Done()
	for v := range ch {
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}

func Sender() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go work(&wg, ch)
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()
	fmt.Println("done")
}
