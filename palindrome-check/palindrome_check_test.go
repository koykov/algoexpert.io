package palindrome_check

import (
	"strconv"
	"testing"
)

var stages = []struct {
	str    string
	expect bool
}{
	{str: "abcdcba", expect: true},
	{str: "a", expect: true},
	{str: "ab", expect: false},
	{str: "aba", expect: true},
	{str: "abb", expect: false},
	{str: "abba", expect: true},
	{str: "abcdefghhgfedcba", expect: true},
	{str: "abcdefghihgfedcba", expect: true},
	{str: "abcdefghihgfeddcba", expect: false},
}

func TestIsPalindrome(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := IsPalindrome(stg.str)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		IsPalindrome(stg.str)
	}
}
