package vartwo

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Sender interface {
	Send(msg string) error
}

type Email struct{}
type SMS struct{}

func (e *Email) Send(msg string) error {

	if strings.TrimSpace(msg) == "" {
		return fmt.Errorf("Message is empty")
	} else if !strings.Contains(msg, "@") {
		return fmt.Errorf("Message has no @")
	}

	return nil
}
func (s *SMS) Send(msg string) error {
	if strings.TrimSpace(msg) == "" {
		return fmt.Errorf("Message is empty")
	}

	return nil
}

func Fan([]Sender, m)

func ExternalCall() string {
	time.Sleep(2 * time.Second)

	return "Done"
}

func FetchWithTimeout() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan string, 1)

	select {
	case ch <- ExternalCall():
		return "API work is done", nil
	case <-ctx.Done():
		return "nil", fmt.Errorf("Waiting time is ended")
	}
}
