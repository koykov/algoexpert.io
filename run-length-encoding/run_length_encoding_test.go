package run_length_encoding

import (
	"strconv"
	"testing"
)

var stages = []struct {
	str, expect string
}{
	{
		str:    "AAAAAAAAAAAAABBCCCCDD",
		expect: "9A4A2B4C2D",
	},
	{
		str:    "aA",
		expect: "1a1A",
	},
	{
		str:    "122333",
		expect: "112233",
	},
	{
		str:    "************^^^^^^^$$$$$$%%%%%%%!!!!!!AAAAAAAAAAAAAAAAAAAA",
		expect: "9*3*7^6$7%6!9A9A2A",
	},
	{
		str:    "aAaAaaaaaAaaaAAAABbbbBBBB",
		expect: "1a1A1a1A5a1A3a4A1B3b4B",
	},
	{
		str:    "                          ",
		expect: "9 9 8 ",
	},
	{
		str:    "1  222 333    444  555",
		expect: "112 321 334 342 35",
	},
	{
		str:    "1A2BB3CCC4DDDD",
		expect: "111A122B133C144D",
	},
	{
		str:    "........______=========AAAA   AAABBBB   BBB",
		expect: "8.6_9=4A3 3A4B3 3B",
	},
	{
		str:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		expect: "9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a1a",
	},
	{
		str:    "        aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		expect: "8 9a9a9a9a9a4a",
	},
	{
		str:    " ",
		expect: "1 ",
	},
	{
		str:    "[(aaaaaaa,bbbbbbb,ccccc,dddddd)]",
		expect: "1[1(7a1,7b1,5c1,6d1)1]",
	},
	{
		str:    ";;;;;;;;;;;;''''''''''''''''''''1233333332222211112222111s",
		expect: "9;3;9'9'2'111273524142311s",
	},
	{
		str:    "AAAAAAAAAAAAABBCCCCDDDDDDDDDDD",
		expect: "9A4A2B4C9D2D",
	},
}

func TestRunLengthEncoding(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := RunLengthEncoding(stg.str)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkRunLengthEncoding(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		RunLengthEncoding(stg.str)
	}
}
