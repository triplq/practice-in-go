package concurpatterns

import "fmt"

func producer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 12; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func producerconsumer(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range ch {
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
