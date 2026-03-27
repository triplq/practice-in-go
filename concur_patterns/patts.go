package concurpatterns

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
