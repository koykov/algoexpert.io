package best_digits

import "strconv"

func BestDigits(n string, d int) string {
	d1 := d
	buf := make([]int, 0, len(n))
	for i := 0; i < len(n); i++ {
		n1, _ := strconv.Atoi(string(n[i]))
		for d > 0 && len(buf) > 0 && n1 > buf[len(buf)-1] {
			buf = buf[:len(buf)-1]
			d--
		}
		buf = append(buf, n1)
	}
	r := make([]byte, 0, len(buf))
	for i := 0; i < len(n)-d1; i++ {
		r = strconv.AppendInt(r, int64(buf[i]), 10)
	}
	return string(r)
}
