package main

import (
	projectpatterns "github.com/triplq/practice-in-go/project_patterns"
)

func main() {
	smart := projectpatterns.NewSmartHouseFacade()
	smart.GoodNight()
}
