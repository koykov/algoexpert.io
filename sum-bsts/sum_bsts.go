package sum_bsts

import "math"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type t struct {
	valid   bool
	mn, mx  int
	c, t, n int
}

func newT() t {
	return t{
		valid: true,
		mn:    math.MaxInt,
		mx:    math.MinInt,
	}
}

func walk(node *BinaryTree) (t t) {
	t = newT()
	if t.valid = node == nil; t.valid {
		return
	}
	v := node.Value
	l, r := walk(node.Left), walk(node.Right)
	if t.valid = l.valid && r.valid && l.mx < v && v <= r.mn; t.valid {
		if t.mn = v; t.mn > l.mn {
			t.mn = l.mn
		}
		if t.mx = v; t.mx < r.mx {
			t.mx = r.mx
		}
		t.c = l.c + r.c + 1
		t.t = l.t + r.t + v
		if t.c > 2 {
			t.n, t.t = t.t, t.n
		}
	}
	t.n += l.n + r.n
	return
}

func SumBsts(tree *BinaryTree) int {
	return walk(tree).n
}
