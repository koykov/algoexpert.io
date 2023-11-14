package youngest_common_ancestor

import "testing"

func TestLCA(t *testing.T) {
	tree := getTrees()
	tree['A'].addAsAncestor(tree['B'], tree['C'])
	tree['B'].addAsAncestor(tree['D'], tree['E'])
	tree['D'].addAsAncestor(tree['H'], tree['I'])
	tree['C'].addAsAncestor(tree['F'], tree['G'])
	yca := GetYoungestCommonAncestor(tree['A'], tree['E'], tree['I'])
	t.Log(yca.Name)
}

func BenchmarkLCA(b *testing.B) {
	tree := getTrees()
	tree['A'].addAsAncestor(tree['B'], tree['C'])
	tree['B'].addAsAncestor(tree['D'], tree['E'])
	tree['D'].addAsAncestor(tree['H'], tree['I'])
	tree['C'].addAsAncestor(tree['F'], tree['G'])
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetYoungestCommonAncestor(tree['A'], tree['E'], tree['I'])
	}
}

func (tree *AncestralTree) addAsAncestor(descendants ...*AncestralTree) {
	for _, descendant := range descendants {
		descendant.Ancestor = tree
	}
}

func getTrees() map[rune]*AncestralTree {
	tree := map[rune]*AncestralTree{}
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		tree[r] = &AncestralTree{Name: string(r)}
	}
	return tree
}
