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
	"strings"

	"github.com/koykov/bytealg"
	"github.com/koykov/bytebuf"
	"github.com/koykov/byteconv"
	"github.com/koykov/jsonvector"
	"github.com/koykov/vector"
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
	difficulty = map[int]string{
		0: "Unknown",
		1: "Easy",
		2: "Medium",
		3: "Hard",
		4: "Very Hard",
	}

	bP   = []byte("<p>")
	bCP  = []byte("</p>")
	bNSS = []byte("\n  ")

	reComment  = regexp.MustCompile(`<span class="[^"]+">(.*)</span>`)
	reP        = regexp.MustCompile(`<p>\n([^<]+)</p>`)
	reH3       = regexp.MustCompile(`<h3>(.*)</h3>`)
	reUL       = regexp.MustCompile(`(?s)<ul>[\n\s]*(.*)[\n\s]*<\/ul>`)
	reLI       = regexp.MustCompile(`\s*<li>[\n\s]*(.*|[\s\S]*?)[\n\s]*<\/li>\s*`)
	reNormPre  = regexp.MustCompile(`(<pre>)([^\n])`)
	reNormCPre = regexp.MustCompile(`([^\n])(</pre>)`)

	auth = flag.String("auth", "", "authorization cookie string")
)

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

	root := jsonvector.Acquire()
	defer jsonvector.Release(root)
	if err = root.Parse(idxRaw); err != nil {
		log.Fatalln(err)
	}
	var c, f int
	root.Get("questions").Each(func(idx int, node *vector.Node) {
		uid := node.GetString("uid")
		// if uid != "find-kth-largest-value-in-bst" { // todo remove me
		// 	return
		// }
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

		uid1 := uid
		if strings.IndexByte(uid, '\'') != -1 {
			uid1 = strings.ReplaceAll(uid1, "'", "")
		}
		pth := fmt.Sprintf("%s/readme.md", uid1)
		if err = os.WriteFile(pth, b, 0644); err != nil {
			f++
			log.Printf("* %s: write error: %s\n", uid, err.Error())
			return
		}

		log.Printf("* %s: success\n", uid)
		c++
	})

	log.Printf("%d questions loaded\n", c)
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
	defer resp.Body.Close()

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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	return b, err
}

func composeReadme(vec vector.Interface) (b []byte, err error) {
	var buf bytebuf.Chain

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
	prompt = reH3.ReplaceAllString(prompt, "\n### $1")
	prompt = reNormPre.ReplaceAllString(prompt, "$1\n$2")
	prompt = reNormCPre.ReplaceAllString(prompt, "$1\n$2")
	prompt = reLI.ReplaceAllString(prompt, "\n* $1")
	prompt = reUL.ReplaceAllString(prompt, "$1\n")
	prompt = strings.ReplaceAll(prompt, "<pre>", "```")
	prompt = strings.ReplaceAll(prompt, `</pre>`, "```")

	d, _ := vec.DotInt("difficulty")
	buf.WriteString("# ").
		Write(vec.DotBytes("name")).
		WriteString("\n\n").
		WriteString("Category: ").WriteString(vec.DotString("category")).WriteString("\n\n").
		WriteString("Difficulty: ").WriteString(difficulty[int(d)]).WriteString("\n\n").
		WriteString("## Description\n\n").
		WriteString(prompt).
		WriteString("\n").
		WriteString("## Optimal Space & Time Complexity\n\n").
		Write(vec.DotBytes("spaceTime")).
		WriteByte('\n')
	b = buf.Bytes()
	return
}
