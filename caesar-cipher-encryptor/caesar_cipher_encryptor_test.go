package caesar_cipher_encryptor

import (
	"strconv"
	"testing"
)

var stages = []struct {
	str    string
	key    int
	expect string
}{
	{
		str:    "xyz",
		key:    2,
		expect: "zab",
	},
	{
		str:    "abc",
		key:    0,
		expect: "abc",
	},
	{
		str:    "abc",
		key:    3,
		expect: "def",
	},
	{
		str:    "xyz",
		key:    5,
		expect: "cde",
	},
	{
		str:    "abc",
		key:    26,
		expect: "abc",
	},
	{
		str:    "abc",
		key:    52,
		expect: "abc",
	},
	{
		str:    "abc",
		key:    57,
		expect: "fgh",
	},
	{
		str:    "xyz",
		key:    25,
		expect: "wxy",
	},
	{
		str:    "iwufqnkqkqoolxzzlzihqfm",
		key:    25,
		expect: "hvtepmjpjpnnkwyykyhgpel",
	},
	{
		str:    "ovmqkwtujqmfkao",
		key:    52,
		expect: "ovmqkwtujqmfkao",
	},
	{
		str:    "mvklahvjcnbwqvtutmfafkwiuagjkzmzwgf",
		key:    7,
		expect: "tcrshocqjuidxcabatmhmrdpbhnqrgtgdnm",
	},
	{
		str:    "kjwmntauvjjnmsagwgawkagfuaugjhawgnawgjhawjgawbfawghesh",
		key:    15,
		expect: "zylbcipjkyycbhpvlvplzpvujpjvywplvcplvywplyvplquplvwthw",
	},
}

func TestCaesarCipherEncryptor(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := CaesarCipherEncryptor(stg.str, stg.key)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkCaesarCipherEncryptor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		CaesarCipherEncryptor(stg.str, stg.key)
	}
}
