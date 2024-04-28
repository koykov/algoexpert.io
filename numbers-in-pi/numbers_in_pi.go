package numbers_in_pi

import (
	"bytes"
)

func NumbersInPi(pi string, numbers []string) int {
	n, m := len(pi), len(numbers)
	x := make([][]byte, len(numbers))
	for i := 0; i < m; i++ {
		x[i] = append(x[i], pi...)
	}
	for i := 0; i < m; i++ {
		nb := []byte(numbers[i])
		for {
			p := bytes.Index(x[i], nb)
			if p == -1 {
				break
			}
			for j := p; j < p+len(nb); j++ {
				x[i][j] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		println(string(x[i]))
	}

	var c int
	for i := 0; i < n; i++ {
		var s byte
		for j := 0; j < m; j++ {
			s += x[j][i]
		}
		if s > 0 {
			c++
		}
	}
	if c == 0 {
		return -1
	}
	return c
}
