package product_sum

import (
	"strconv"
	"testing"
)

var stages = []struct {
	array  []any
	expect int
}{
	{
		array:  []any{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4}},
		expect: 12,
	},
	{
		array:  []any{1, 2, 3, 4, 5},
		expect: 15,
	},
	{
		array:  []any{1, 2, SpecialArray{3}, 4, 5},
		expect: 18,
	},
	{
		array:  []any{SpecialArray{1, 2}, 3, SpecialArray{4, 5}},
		expect: 27,
	},
	{
		array:  []any{SpecialArray{SpecialArray{SpecialArray{SpecialArray{5}}}}},
		expect: 600,
	},
	{
		array:  []any{9, SpecialArray{2, -3, 4}, 1, SpecialArray{1, 1, SpecialArray{1, 1, 1}}, SpecialArray{SpecialArray{SpecialArray{SpecialArray{3, 4, 1}}}, 8}, SpecialArray{1, 2, 3, 4, 5, SpecialArray{6, 7}, -7}, SpecialArray{1, SpecialArray{2, 3, SpecialArray{4, 5}}, SpecialArray{6, 0, SpecialArray{7, 0, -8}}, -7}, SpecialArray{1, -3, 2, SpecialArray{1, -3, 2, SpecialArray{1, -3, 2}, SpecialArray{1, -3, 2, SpecialArray{1, -3, 2}}, SpecialArray{1, -3, 2}}}, -3},
		expect: 1351,
	},
}

func TestProductSum(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := ProductSum(stg.array)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkProductSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ProductSum(stages[i%len(stages)].array)
	}
}
