package varone

import (
	"fmt"
	"strings"
	"unicode"
)

type Account struct {
	balance float64
}

func (acc *Account) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("insufficient funds")
	}
	acc.balance += amount
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount > acc.balance {
		return fmt.Errorf("insufficient funds")
	}
	acc.balance -= amount
	return nil
}

func sum(nums ...int) (sum int) {
	for _, n := range nums {
		sum += n
	}

	return
}

func clean(data []string) []string {
	var i int
	for _, s := range data {
		sTrim := strings.TrimSpace(s)
		if sTrim != "" {
			data[i] = sTrim
			i++
		}
	}

	data = data[:i]

	return data
}

func wordCount(text string) map[rune]int {
	count := make(map[rune]int)

	for _, s := range text {
		lwrS := unicode.ToLower(s)
		if lwrS != ' ' {
			count[lwrS] += 1
		}
	}

	return count
}
