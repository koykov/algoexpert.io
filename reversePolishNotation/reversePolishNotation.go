package reversePolishNotation

import "strconv"

var op = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
}

func ReversePolishNotation(t []string) int {
	if len(t) == 1 {
		x, _ := strconv.Atoi(t[0])
		return x
	}
	for i := 0; i < len(t); i++ {
		x := t[i]
		if _, ok := op[x]; ok {
			rawl, rawr := t[i-2], t[i-1]
			l, _ := strconv.Atoi(rawl)
			r, _ := strconv.Atoi(rawr)
			var s int
			switch x {
			case "+":
				s = l + r
			case "-":
				s = l - r
			case "*":
				s = l * r
			case "/":
				s = l / r
			}
			t[i-2] = strconv.Itoa(s)
			t = append(t[:i-1], t[i+1:]...)
			i -= 2
		}
	}
	x, _ := strconv.Atoi(t[0])
	return x
}
