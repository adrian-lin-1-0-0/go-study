package compareandswap

func CompareAndSwap(addr *int32, oldVal, newVal int32) bool {
	if *addr == oldVal {
		*addr = newVal
		return true
	}
	return false
}
