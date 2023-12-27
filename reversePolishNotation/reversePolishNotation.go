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
	var l, r int
	rawl, rawr := t[len(t)-3], t[len(t)-2]
	if _, ok := op[rawr]; ok {
		r = ReversePolishNotation(t[:len(t)-1])
	} else {
		r, _ = strconv.Atoi(rawr)
	}
	if _, ok := op[rawl]; ok {
		l = ReversePolishNotation(t[:len(t)-2])
	} else {
		l, _ = strconv.Atoi(rawl)
	}
	switch t[len(t)-1] {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r

	}
	return -1
}
