package water_area

func WaterArea(h []int) (a int) {
	if len(h) == 0 {
		return 0
	}
	mx, mxi := 0, 0
	for i := 0; i < len(h); i++ {
		if h[i] > mx {
			mx = h[i]
			mxi = i
		}
	}

	cx := 0
	for i := 0; h[i] < mx; i++ {
		if h[i] > cx {
			cx = h[i]
			continue
		}
		a += cx
		if h[i] > 0 {
			a -= h[i]
		}
	}

	cx = 0
	for i := len(h) - 1; i > mxi; i-- {
		if h[i] > cx {
			cx = h[i]
			continue
		}
		a += cx
		if h[i] > 0 {
			a -= h[i]
		}
	}
	return
}
