package nth_fibonacci

func GetNthFib(n int) (f int) {
	f0, f1 := 0, 1
	if n == 1 {
		return f0
	}
	if n == 2 {
		return f1
	}
	for i := 2; i < n; i++ {
		f = f0 + f1
		f1, f0 = f, f1
	}
	return
}
