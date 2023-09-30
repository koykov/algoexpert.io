package semordnilap

import (
	"strconv"
	"testing"
)

var stages = []struct {
	words  []string
	expect [][]string
}{
	{
		words:  []string{},
		expect: [][]string{},
	},
	{
		words:  []string{"aaa", "bbbb"},
		expect: [][]string{},
	},
	{
		words:  []string{"dog", "god"},
		expect: [][]string{{"dog", "god"}},
	},
	{
		words:  []string{"dog", "hello", "god"},
		expect: [][]string{{"dog", "god"}},
	},
	{
		words:  []string{"dog", "desserts", "god", "stressed"},
		expect: [][]string{{"dog", "god"}, {"desserts", "stressed"}},
	},
	{
		words:  []string{"dog", "hello", "desserts", "test", "god", "stressed"},
		expect: [][]string{{"dog", "god"}, {"desserts", "stressed"}},
	},
	{
		words:  []string{"abcde", "edcba", "edbc", "edb", "cbde", "abc"},
		expect: [][]string{{"cbde", "edbc"}, {"abcde", "edcba"}},
	},
	{
		words:  []string{"abcde", "bcd", "dcb", "edcba", "aaa"},
		expect: [][]string{{"edcba", "abcde"}, {"bcd", "dcb"}},
	},
	{
		words:  []string{"abcdefghijk", "aaa", "hello", "racecar", "kjihgfedcba"},
		expect: [][]string{{"kjihgfedcba", "abcdefghijk"}},
	},
	{
		words:  []string{"liver", "dog", "hello", "desserts", "evil", "test", "god", "stressed", "racecar", "palindromes", "semordnilap", "abcd", "live"},
		expect: [][]string{{"stressed", "desserts"}, {"palindromes", "semordnilap"}, {"god", "dog"}, {"live", "evil"}},
	},
}

func TestSemordnilap(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := Semordnilap(stg.words)
			_ = r
		})
	}
}

func BenchmarkSemordnilap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		Semordnilap(stg.words)
	}
}
