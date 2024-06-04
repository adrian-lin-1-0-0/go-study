package testandset

import "testing"

func Test_testandset(t *testing.T) {

	lock := false

	for TestAndSet(&lock) {
	}
	// critical section
	lock = false
}
