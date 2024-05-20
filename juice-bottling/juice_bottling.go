package juice_bottling

func JuiceBottling(p []int) []int {
	n := len(p)
	var (
		mxp float32
		xpi int
	)
	mxpidx := make([]int, 1, n)

	for i := 1; i < n; i++ {
		up := float32(p[i]) / float32(i)
		if up >= mxp {
			xpi = i
			mxp = up
		}
		mxpidx = append(mxpidx, xpi)
	}

	buf := make([]int, 0, n)
	r := n - 1
	for r > 0 {
		buf = append(buf, mxpidx[r])
		r -= mxpidx[r]
	}

	res := make([]int, len(buf))
	for i, v := range buf {
		res[len(buf)-1-i] = v
	}

	return res
}
