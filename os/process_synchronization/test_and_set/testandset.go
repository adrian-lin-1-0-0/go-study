package testandset

func TestAndSet(lock *bool) bool {
	old := *lock
	*lock = true
	return old
}
