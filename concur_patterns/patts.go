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
