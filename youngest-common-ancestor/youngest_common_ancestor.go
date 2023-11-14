package youngest_common_ancestor

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, d1, d2 *AncestralTree) (r *AncestralTree) {
	l1, l2 := lvl(d1, top), lvl(d2, top)
	if l1 > l2 {
		d1 = sh(d1, l1-l2)
	} else {
		d2 = sh(d2, l2-l1)
	}
	for {
		if d1.Name == d2.Name {
			r = d1
			break
		}
		d1, d2 = d1.Ancestor, d2.Ancestor
	}
	return
}

func lvl(a, b *AncestralTree) (r int) {
	for {
		if a.Name == b.Name {
			return r
		}
		r++
		a = a.Ancestor
	}
}

func sh(a *AncestralTree, n int) *AncestralTree {
	for i := 0; i < n; i++ {
		a = a.Ancestor
	}
	return a
}

// func GetYoungestCommonAncestor(top, d1, d2 *AncestralTree) *AncestralTree {
// 	p1, p2 := la(d1, top), la(d2, top)
// 	if p1 == nil || p2 == nil {
// 		return nil
// 	}
// 	if p1.Name == p2.Name {
// 		return p1
// 	}
// 	return nil
// }
//
// func la(n, t *AncestralTree) *AncestralTree {
// 	p := n.Ancestor
// 	for {
// 		p_ := p
// 		if p = p.Ancestor; p == nil {
// 			return nil
// 		}
// 		if p.Name == t.Name {
// 			p = p_
// 			break
// 		}
// 	}
// 	return p
// }
