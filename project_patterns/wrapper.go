package projectpatterns

import "fmt"

type Service interface {
	DoSomething(input string) string
}

type BaseService struct{}

func (bs *BaseService) DoSomething(input string) string {
	return "Procces: " + input
}

type LoggingWrapper struct {
	next Service
}

func (w *LoggingWrapper) DoSomething(input string) string {
	fmt.Println("Before:", input)
	result := w.next.DoSomething(input)
	fmt.Println("After:", result)

	return result
}
