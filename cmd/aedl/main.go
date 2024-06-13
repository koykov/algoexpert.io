package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/koykov/jsonvector"
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
	var vec jsonvector.Vector
	var buf bytes.Buffer

	idxRaw, err := dlIndex()
	if err != nil {
		log.Fatalln(err)
	}
	if err = vec.Parse(idxRaw); err != nil {
		log.Fatalln(err)
	}
	_ = vec.Beautify(&buf)
	os.WriteFile("list.json", buf.Bytes(), 0644)

	qRaw, err := dlQuestion("disk-stacking")
	if err != nil {
		log.Fatalln(err)
	}
	vec.Reset()
	buf.Reset()
	if err = vec.Parse(qRaw); err != nil {
		log.Fatalln(err)
	}
	_ = vec.Beautify(&buf)
	os.WriteFile("dist-stacking.json", buf.Bytes(), 0644)
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
