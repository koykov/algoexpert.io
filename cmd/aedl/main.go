package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/koykov/bytealg"
	"github.com/koykov/bytebuf"
	"github.com/koykov/byteconv"
	"github.com/koykov/jsonvector"
	"github.com/koykov/vector"
)

const (
	readmeMainHeader = `# Algoexpert.io

[algoexpert.io](https://www.algoexpert.io) is a good resource to prepare for coding interviews. This repo contains my solutions of AE problems.

See index section below, there are solution grouped by both [difficulty](#group-by-difficulty) and [category](#group-by-category).`
)

var (
	headers = map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
		"cache-control":      "no-cache",
		"content-length":     "0",
		"dnt":                "1",
		"origin":             "https://www.algoexpert.io",
		"pragma":             "no-cache",
		"priority":           "u=1, i",
		"referer":            "https://www.algoexpert.io/",
		"sec-ch-ua":          `"Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Linux"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-site",
		"user-agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
	}
	difficulty = []string{
		0: "Unknown",
		1: "Easy",
		2: "Medium",
		3: "Hard",
		4: "Very Hard",
	}

	bP   = []byte("<p>")
	bCP  = []byte("</p>")
	bNSS = []byte("\n  ")

	reComment   = regexp.MustCompile(`<span class="[^"]+">(&.*|[\s\S]*?)</span>`)
	reP         = regexp.MustCompile(`<p>\n([^<]+)</p>`)
	reP1        = regexp.MustCompile(`<p>([^<]+)</p>`)
	reH3        = regexp.MustCompile(`<h3>[\n\s]*(.*)[\n\s]*</h3>`)
	reUL        = regexp.MustCompile(`(?s)<ul>[\n\s]*(.*)[\n\s]*</ul>`)
	reOL        = regexp.MustCompile(`(?s)<ol>[\n\s]*(.*)[\n\s]*</ol>`)
	reLI        = regexp.MustCompile(`\s*<li>[\n\s]*(.*|[\s\S]*?)[\n\s]*</li>\s*`)
	reNormPre   = regexp.MustCompile(`(<pre>)([^\n])`)
	reNormCPre  = regexp.MustCompile(`([^\n])(</pre[\n\s]*>)`)
	reNormPre1  = regexp.MustCompile(`</p>\s<pre>`)
	reNormCPre1 = regexp.MustCompile(`([^\n]*)(</pre[\n\s]+>)`)

	auth = flag.String("auth", "", "authorization cookie string")
)

type tuple struct {
	name, uid, slug, cat string
	difficulty           int64
	ok                   bool
}

func init() {
	flag.Parse()
	if auth == nil || len(*auth) == 0 {
		log.Fatalln("no auth cookie provided")
	}
	headers["authorization"] = *auth
}

func main() {
	idxRaw, err := dlIndex()
	if err != nil {
		log.Fatalln(err)
	}

	var (
		diffReg = make(map[int64][]*tuple)
		catReg  = make(map[string][]*tuple)
		cats    []string
	)

	root := jsonvector.Acquire()
	defer jsonvector.Release(root)
	if err = root.Parse(idxRaw); err != nil {
		log.Fatalln(err)
	}
	var c, f, t int
	root.Get("questions").Each(func(idx int, node *vector.Node) {
		t++

		uid := node.GetString("uid")
		slug := uid
		if strings.IndexByte(uid, '\'') != -1 {
			slug = strings.ReplaceAll(slug, "'", "")
		}

		var tpl tuple
		tpl.difficulty, _ = node.GetInt("difficulty")
		tpl.name = node.GetString("name")
		tpl.cat = node.GetString("category")
		tpl.uid, tpl.slug = uid, slug

		if _, err := os.Stat(slug); !os.IsNotExist(err) {
			tpl.ok = true
		}

		diffReg[tpl.difficulty] = append(diffReg[tpl.difficulty], &tpl)
		if _, ok := catReg[tpl.cat]; !ok {
			cats = append(cats, tpl.cat)
		}
		catReg[tpl.cat] = append(catReg[tpl.cat], &tpl)

		qRaw, err := dlQuestion(uid)
		if err != nil {
			f++
			log.Printf("* %s: download error: %s\n", uid, err.Error())
			return
		}

		vec := jsonvector.Acquire()
		defer jsonvector.Release(vec)
		if err = vec.Parse(qRaw); err != nil {
			f++
			log.Printf("* %s: parse error: %s\n", uid, err.Error())
			return
		}

		b, err := composeReadme(vec)
		if err != nil {
			f++
			log.Printf("* %s: compose error: %s\n", uid, err.Error())
			return
		}

		pth := fmt.Sprintf("%s/readme.md", slug)
		if err = os.WriteFile(pth, b, 0644); err != nil {
			f++
			log.Printf("* %s: write error: %s\n", uid, err.Error())
			return
		}

		log.Printf("* %s: success\n", uid)
		c++
	})

	_ = f
	log.Printf("%d questions of %d loaded\n", c, t)

	b, err := composeReadmeMain(diffReg, catReg, cats)
	if err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile("readme.md", b, 0644); err != nil {
		log.Fatalln(err)
	}
}

func dlIndex() ([]byte, error) {
	req, _ := http.NewRequest(http.MethodPost, "https://prod.api.algoexpert.io/api/problems/v1/algoexpert/coding-questions/list", nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := io.ReadAll(resp.Body)
	return b, err
}

func dlQuestion(slug string) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(`{"name":"` + slug + `"}`)
	req, _ := http.NewRequest(http.MethodPost, "https://prod.api.algoexpert.io/api/problems/v1/algoexpert/coding-questions/get", &buf)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := io.ReadAll(resp.Body)
	return b, err
}

func composeReadmeMain(diffReg map[int64][]*tuple, catReg map[string][]*tuple, cats []string) (b []byte, err error) {
	buf := bytebuf.Chain{}

	buf.WriteString(readmeMainHeader).WriteString("\n\n")

	var mt int
	buf.WriteString("## Group by difficulty\n\n").
		WriteByte('|')
	for i := 1; i < len(difficulty); i++ {
		dn := difficulty[i]
		tuples := diffReg[int64(i)]
		var s, t int64
		for j := 0; j < len(tuples); j++ {
			tpl := tuples[j]
			t++
			if tpl.ok {
				s++
			}
		}
		buf.WriteString(dn).
			WriteString(" (").
			WriteInt(s).
			WriteByte('/').
			WriteInt(t).
			WriteString(") |")

		sort.Slice(tuples, func(i, j int) bool {
			return tuples[i].name < tuples[j].name
		})

		mt = max(mt, len(tuples))
	}
	buf.WriteString("\n| --- | --- | --- | --- |\n")

	for i := 0; i < mt; i++ {
		buf.WriteByte('|')
		for j := 1; j < len(difficulty); j++ {
			buf.WriteByte(' ')
			tuples := diffReg[int64(j)]
			if len(tuples) > i {
				tpl := tuples[i]
				buf.WriteByteIf(tpl.ok, '[').
					WriteString(tpl.name).
					WriteStringIf(tpl.ok, "](").
					WriteStringIf(tpl.ok, tpl.slug).
					WriteByteIf(tpl.ok, ')')
			} else {
				buf.WriteString(" ")
			}
			buf.WriteString(" |")
		}
		buf.WriteByte('\n')
	}
	buf.WriteString("\n\n")

	buf.WriteString("## Group by category\n\n")
	sort.Strings(cats)
	for i := 0; i < len(cats); i++ {
		cat := cats[i]
		var s, t int64
		tuples := catReg[cat]
		sort.Slice(tuples, func(i, j int) bool {
			return tuples[i].difficulty < tuples[j].difficulty
		})
		for j := 0; j < len(tuples); j++ {
			t++
			if tuples[j].ok {
				s++
			}
		}
		buf.WriteString("### ").
			WriteString(cat).
			WriteString(" (").
			WriteInt(s).
			WriteByte('/').
			WriteInt(t).
			WriteString(")\n\n").
			WriteString("| Name | Difficulty |\n").
			WriteString("| --- | --- |\n")
		for j := 0; j < len(tuples); j++ {
			buf.WriteString("| ")
			tpl := tuples[j]
			buf.WriteByteIf(tpl.ok, '[').
				WriteString(tpl.name).
				WriteStringIf(tpl.ok, "](").
				WriteStringIf(tpl.ok, tpl.slug).
				WriteByteIf(tpl.ok, ')')
			buf.WriteString(" | ")
			buf.WriteString(difficulty[tpl.difficulty])
			buf.WriteString(" |\n")
		}
		buf.WriteString("\n\n")
	}

	b = buf.Bytes()
	return
}

func composeReadme(vec vector.Interface) (b []byte, err error) {
	buf := bytebuf.Chain{}

	prompt := vec.DotString("prompt")
	prompt = strings.ReplaceAll(prompt, "<div class=\"html\">\n", "")
	prompt = strings.ReplaceAll(prompt, `</div>`, "")
	prompt = reComment.ReplaceAllString(prompt, "$1")
	prompt = strings.ReplaceAll(prompt, "<span>", "`")
	prompt = strings.ReplaceAll(prompt, `</span>`, "`")

	pb := byteconv.S2B(prompt)
	var off int
	for {
		s := bytealg.IndexAt(pb, bP, off)
		e := bytealg.IndexAt(pb, bCP, off+1)
		if s != -1 && e != -1 {
			poff := s + 3
			for {
				p := bytealg.IndexAt(pb, bNSS, poff)
				if p == -1 || p > e {
					break
				}
				copy(pb[p+1:], pb[p+3:])
				pb = pb[:len(pb)-2]
				e -= 2
			}
			off = e
			continue
		}
		break
	}
	prompt = byteconv.B2S(pb)
	prompt = reP.ReplaceAllString(prompt, "$1")
	prompt = reP1.ReplaceAllString(prompt, "\n$1\n")
	prompt = strings.ReplaceAll(prompt, `&lt;`, "<")
	prompt = strings.ReplaceAll(prompt, `&gt;`, ">")
	prompt = reH3.ReplaceAllString(prompt, "\n### $1")
	prompt = reNormPre.ReplaceAllString(prompt, "<pre>\n$2")
	prompt = reNormPre1.ReplaceAllString(prompt, "</p>\n\n<pre>")
	prompt = reNormCPre.ReplaceAllString(prompt, "$1\n</pre>")
	prompt = reNormCPre1.ReplaceAllString(prompt, "$1</pre>")
	prompt = reLI.ReplaceAllString(prompt, "\n* $1")
	prompt = reUL.ReplaceAllString(prompt, "$1\n")
	prompt = reOL.ReplaceAllString(prompt, "$1\n")
	prompt = strings.ReplaceAll(prompt, "<pre>", "```")
	prompt = strings.ReplaceAll(prompt, `</pre>`, "```")
	prompt = strings.ReplaceAll(prompt, "\n\n```", "\n```")

	d, _ := vec.DotInt("difficulty")
	buf.WriteString("# ").
		Write(vec.DotBytes("name")).
		WriteString("\n\n").
		WriteString("Category: ").WriteString(vec.DotString("category")).WriteString("\n\n").
		WriteString("Difficulty: ").WriteString(difficulty[d]).WriteString("\n\n").
		WriteString("## Description\n\n").
		WriteString(prompt).
		WriteString("\n").
		WriteString("## Optimal Space & Time Complexity\n\n").
		Write(vec.DotBytes("spaceTime")).
		WriteByte('\n')
	b = buf.Bytes()
	return
}
