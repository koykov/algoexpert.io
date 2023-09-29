package nth_fibonacci

import (
	"strconv"
	"testing"
)

var stages = []struct {
	n, expect int
}{
	{6, 5},
	{1, 0},
	{2, 1},
	{3, 1},
	{4, 2},
	{5, 3},
	{7, 8},
	{8, 13},
	{9, 21},
	{10, 34},
	{11, 55},
	{12, 89},
	{13, 144},
	{14, 233},
	{15, 377},
	{16, 610},
	{17, 987},
	{18, 1597},
}

func TestGetNthFib(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := GetNthFib(stg.n)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkGetNthFib(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetNthFib(stages[i%len(stages)].n)
	}
}
