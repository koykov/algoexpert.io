package first_non_repeating_character

import (
	"strconv"
	"testing"
)

var stages = []struct {
	str    string
	expect int
}{
	{
		str:    "abcdcaf",
		expect: 1,
	},
	{
		str:    "faadabcbbebdf",
		expect: 6,
	},
	{
		str:    "a",
		expect: 0,
	},
	{
		str:    "ab",
		expect: 0,
	},
	{
		str:    "abc",
		expect: 0,
	},
	{
		str:    "abac",
		expect: 1,
	},
	{
		str:    "ababac",
		expect: 5,
	},
	{
		str:    "ababacc",
		expect: -1,
	},
	{
		str:    "lmnopqldsafmnopqsa",
		expect: 7,
	},
	{
		str:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy",
		expect: 25,
	},
	{
		str:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		expect: -1,
	},
	{
		str:    "",
		expect: -1,
	},
	{
		str:    "ggyllaylacrhdzedddjsc",
		expect: 10,
	},
	{
		str:    "aaaaaaaaaaaaaaaaaaaabbbbbbbbbbcccccccccccdddddddddddeeeeeeeeffghgh",
		expect: -1,
	},
	{
		str:    "aabbccddeeff",
		expect: -1,
	},
}

func TestFirstNonRepeatingCharacter(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := FirstNonRepeatingCharacter(stg.str)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkFirstNonRepeatingCharacter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		FirstNonRepeatingCharacter(stg.str)
	}
}
