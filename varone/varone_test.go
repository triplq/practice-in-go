package varone

import (
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name  string
		numbs []int
		exp   int
	}{
		{"pos", []int{3, 4, 5}, 12},
		{"neg", []int{5, -9, 1}, -3},
		{"empty", []int{}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := sum(tc.numbs...)
			if res != tc.exp {
				t.Errorf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestClean(t *testing.T) {
	cases := []struct {
		name string
		str  []string
		exp  []string
	}{
		{"remove spaces", []string{"hello world", "  leading and trailing  ", "no spaces"}, []string{"hello world", "leading and trailing", "no spaces"}},
		{"empty strings", []string{"", "  ", "test"}, []string{"test"}},
		{"mixed strings", []string{"a b c", "d", " e f "}, []string{"a b c", "d", "e f"}},
		{"no strings", []string{}, []string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := clean(tc.str)
			for i := range res {
				if res[i] != tc.exp[i] {
					t.Errorf("expected %s, got %s", tc.exp[i], res[i])
				}
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	cases := []struct {
		name string
		str  string
		exp  map[rune]int
	}{
		{
			name: "simple sentence",
			str:  "hello world",
			exp:  map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, 'w': 1, 'r': 1, 'd': 1},
		},
		{
			name: "sentence with punctuation",
			str:  "hello, world!",
			exp:  map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, ',': 1, 'w': 1, 'r': 1, 'd': 1, '!': 1},
		},
		{
			name: "empty string",
			str:  "",
			exp:  map[rune]int{},
		},
		{
			name: "string with only spaces",
			str:  "   ",
			exp:  map[rune]int{},
		},
		{
			name: "string with numbers and symbols",
			str:  "123 abc !@#",
			exp:  map[rune]int{'1': 1, '2': 1, '3': 1, 'a': 1, 'b': 1, 'c': 1, '!': 1, '@': 1, '#': 1},
		},
		{
			name: "string with varied casing",
			str:  "Hello World",
			exp:  map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, 'w': 1, 'r': 1, 'd': 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := wordCount(tc.str)
			if len(res) != len(tc.exp) {
				t.Errorf("test %s is incorrect by len", tc.name)
			} else {
				for k, v1 := range tc.exp {
					v2, ok := res[k]
					if !ok || v1 != v2 {
						t.Errorf("test %s is incorrect by match", tc.name)
					}
				}
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	acc := Account{
		balance: 100,
	}
	cases := []struct {
		name   string
		amount float64
		exp    float64
		err    bool
	}{
		{"correct", 100, 200, false},
		{"incorrect", -50, 200, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := acc.Deposit(tc.amount)
			if (err != nil) != tc.err {
				t.Errorf("error = %v, wantErr %v", err, tc.err)
			}
			if acc.balance != tc.exp {
				t.Errorf("expected %f, got %f", tc.exp, acc.balance)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	acc := Account{
		balance: 100,
	}
	cases := []struct {
		name   string
		amount float64
		exp    float64
		err    bool
	}{
		{"correct", 50, 50, false},
		{"incorrect", 200, 50, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := acc.Withdraw(tc.amount)
			if (err != nil) != tc.err {
				t.Errorf("error = %v, wantErr %v", err, tc.err)
			}
			if acc.balance != tc.exp {
				t.Errorf("expected %f, got %f", tc.exp, acc.balance)
			}
		})
	}
}
