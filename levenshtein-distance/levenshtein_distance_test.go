package levenshtein_distance

import (
	"strconv"
	"testing"
)

var stages = []struct {
	a, b   string
	expect int
}{
	{
		a:      "abc",
		b:      "yabd",
		expect: 2,
	},
	{
		a:      "biting",
		b:      "mitten",
		expect: 4,
	},
}

func TestLevenshteinDistance(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := LevenshteinDistance(stg.a, stg.b)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		LevenshteinDistance(stg.a, stg.b)
	}
}
