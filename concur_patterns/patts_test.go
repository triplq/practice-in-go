package concurpatterns

import "testing"

func TestSender(t *testing.T) {
	t.Run("Sender", func(t *testing.T) {
		Sender()
	})
}

func TestPool(t *testing.T) {
	t.Run("Pool", func(t *testing.T) {
		Demo_pool()
	})
}
