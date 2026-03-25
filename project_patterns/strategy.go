package projectpatterns

import (
	"fmt"
	"os"
	"reflect"
)

type PaymentStrategy interface {
	Pay(amount float64)
}

type Card struct{}

type Cash struct{}

func (c *Card) Pay(amount float64) {
	fmt.Printf("Payed %f with card\n", amount)
}

func (c *Cash) Pay(amount float64) {
	fmt.Printf("Payed %f with cash\n", amount)
}

type Shop struct {
	strategy PaymentStrategy
}

func (s *Shop) SetStrategy(strategy PaymentStrategy) {
	s.strategy = strategy
	switch value := strategy.(type) {
	case *Card:
		typeName := reflect.TypeOf(value).Elem().Name()
		fmt.Printf("Payment strategy is changed on %s\n", typeName)
	case *Cash:
		typeName := reflect.TypeOf(value).Elem().Name()
		fmt.Printf("Payment strategy is changed on %s\n", typeName)
	default:
		fmt.Fprint(os.Stderr, "Error with type switch")
	}
}

func (s *Shop) MakePay(amount float64) {
	s.strategy.Pay(amount)
}
