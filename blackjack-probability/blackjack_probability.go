package blackjack_probability

func BlackjackProbability(t int, h int) float64 {
	pb := []float64{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}
	for v := t - 5; v >= h; v-- {
		s := float64(0)
		for _, prob := range pb {
			s += prob / 10
		}
		copy(pb[1:], pb[:len(pb)-1])
		pb[0] = s
	}
	return pb[0]
}
