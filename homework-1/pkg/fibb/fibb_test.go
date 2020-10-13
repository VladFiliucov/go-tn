package fibb

import "testing"

func TestFib(t *testing.T) {
	if Fib(3) != 3 {
		t.Errorf("Nooo %v", Fib(3))
	}
}
