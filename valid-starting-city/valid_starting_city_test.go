package valid_starting_city

import (
	"strconv"
	"testing"
)

var stages = []struct {
	d, f   []int
	mpg    int
	expect int
}{
	{
		d:      []int{5, 25, 15, 10, 15},
		f:      []int{1, 2, 1, 0, 3},
		mpg:    10,
		expect: 4,
	},
	{
		d:      []int{30, 40, 10, 10, 17, 13, 50, 30, 10, 40},
		f:      []int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		mpg:    25,
		expect: 0,
	},
}

func TestValidStartingCity(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := ValidStartingCity(stg.d, stg.f, stg.mpg)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkValidStartingCity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := stages[i%len(stages)]
		ValidStartingCity(stg.d, stg.f, stg.mpg)
	}
}
