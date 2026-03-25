package main

import (
	projectpatterns "github.com/triplq/practice-in-go/project_patterns"
)

func main() {
	shop := &projectpatterns.Shop{}
	shop.SetStrategy(&projectpatterns.Cash{})
	shop.MakePay(20)
	shop.SetStrategy(&projectpatterns.Card{})
	shop.MakePay(50)
}
