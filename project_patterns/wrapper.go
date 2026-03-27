package projectpatterns

import "fmt"

type Service interface {
	DoSomething(input string) string
}

type BaseService struct{}

func (bs *BaseService) DoSomething(input string) string {
	return "Procces: " + input
}

type WrappedService struct {
	next Service
}

func (w *WrappedService) DoSomething(input string) string {
	fmt.Println("Before:", input)
	result := w.next.DoSomething(input)
	fmt.Println("After:", result)

	return result
}

func demo_main() {
	simple := &BaseService{}

	wrapped := &WrappedService{
		next: simple,
	}

	fmt.Println(wrapped.DoSomething("Hello"))
}
