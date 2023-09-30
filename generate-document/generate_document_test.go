package generate_document

import (
	"strconv"
	"testing"
)

var stages = []struct {
	chars, doc string
	expect     bool
}{
	{
		chars:  "Bste!hetsi ogEAxpelrt x ",
		doc:    "AlgoExpert is the Best!",
		expect: true,
	},
	{
		chars:  "A",
		doc:    "a",
		expect: false,
	},
	{
		chars:  "a",
		doc:    "a",
		expect: true,
	},
	{
		chars:  "a hsgalhsa sanbjksbdkjba kjx",
		doc:    "",
		expect: true,
	},
	{
		chars:  " ",
		doc:    "hello",
		expect: false,
	},
	{
		chars:  "     ",
		doc:    "     ",
		expect: true,
	},
	{
		chars:  "aheaollabbhb",
		doc:    "hello",
		expect: true,
	},
	{
		chars:  "aheaolabbhb",
		doc:    "hello",
		expect: false,
	},
	{
		chars:  "estssa",
		doc:    "testing",
		expect: false,
	},
	{
		chars:  "Bste!hetsi ogEAxpert",
		doc:    "AlgoExpert is the Best!",
		expect: false,
	},
	{
		chars:  "helloworld ",
		doc:    "hello wOrld",
		expect: false,
	},
	{
		chars:  "helloworldO",
		doc:    "hello wOrld",
		expect: false,
	},
	{
		chars:  "helloworldO ",
		doc:    "hello wOrld",
		expect: true,
	},
	{
		chars:  "&*&you^a%^&8766 _=-09     docanCMakemthisdocument",
		doc:    "Can you make this document &",
		expect: true,
	},
	{
		chars:  "abcabcabcacbcdaabc",
		doc:    "bacaccadac",
		expect: true,
	},
}

func TestGenerateDocument(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := GenerateDocument(stg.chars, stg.doc)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkGenerateDocument(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		GenerateDocument(stg.chars, stg.doc)
	}
}
