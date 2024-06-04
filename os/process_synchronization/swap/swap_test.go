package swap

import "testing"

func Swap(a *bool, b *bool) {
	*a, *b = *b, *a
}

func Test_swap(t *testing.T) {

	lock, old := false, true

	for old {
		Swap(&lock, &old)
	}
	// critical section
	lock = false
}
