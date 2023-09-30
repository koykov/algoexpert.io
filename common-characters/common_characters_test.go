package common_characters

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

var stages = []struct {
	ss     []string
	expect string
}{
	{
		ss:     []string{"abc", "bcd", "cbad"},
		expect: "[b c]",
	},
	{
		ss:     []string{"a"},
		expect: "[a]",
	},
	{
		ss:     []string{"a", "b", "c"},
		expect: "[]",
	},
	{
		ss:     []string{"aa", "aa"},
		expect: "[a]",
	},
	{
		ss:     []string{"aaaa", "a"},
		expect: "[a]",
	},
	{
		ss:     []string{"abcde", "aa", "foobar", "foobaz", "and this is a string", "aaaaaaaa", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeea"},
		expect: "[a]",
	},
	{
		ss:     []string{"abcdef", "fedcba", "abcefd", "aefbcd", "efadbc", "effffffffffffbcda", "eeeeeffffffbbbbbaaaaaccccdddd", "**************abdcef************"},
		expect: "[a b c d e f]",
	},
	{
		ss:     []string{"ab&cdef!", "f!ed&cba", "a&bce!d", "ae&fb!cd", "efa&!dbc", "eff!&fff&fffffffbcda", "eeee!efff&fffbbbbbaaaaaccccdddd", "*******!***&****abdcef************", "*******!***&****a***********f*", "*******!***&****b***********c*"},
		expect: "[! &]",
	},
	{
		ss:     []string{"*abcd", "def*", "******d*****"},
		expect: "[* d]",
	},
	{
		ss:     []string{"*abc!d", "de!f*", "**!!!****d*****"},
		expect: "[! * d]",
	},
}

func TestCommonCharacters(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := CommonCharacters(stg.ss)
			sort.Strings(r)
			if fmt.Sprintf("%v", r) != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkCommonCharacters(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		CommonCharacters(stg.ss)
	}
}
