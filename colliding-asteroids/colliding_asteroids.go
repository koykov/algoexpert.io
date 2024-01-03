package colliding_asteroids

func CollidingAsteroids(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] < 0 && i > 0 && a[i-1] > 0 {
			switch {
			case a[i-1] > -a[i]:
				a = append(a[:i], a[i+1:]...)
				i--
			case a[i-1] < -a[i]:
				for j := i - 1; j >= 0 && a[j] > 0 && a[j] < -a[i]; j-- {
					a = append(a[:i-1], a[i:]...)
					i--
				}
			case a[i-1] == -a[i]:
				a = append(a[:i-1], a[i+1:]...)
				i--
			}
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] > 0 && i < len(a)-1 && a[i+1] < 0 {
			switch {
			case a[i] < -a[i+1]:
				a = append(a[:i], a[i+1:]...)
				i--
			case a[i] > -a[i+1]:
				for j := i + 1; j < len(a) && a[j] < 0 && -a[j] < a[i]; j++ {
					a = append(a[:i+1], a[i+2:]...)
					i--
				}
			case a[i] == -a[i+1]:
				a = append(a[:i-1], a[i+1:]...)
				i--
			}
		}
	}
	return a
}
