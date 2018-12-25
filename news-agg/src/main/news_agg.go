package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Locations []string `xml:"url>loc"`
	Keywords  []string `xml:"url>news>keywords"`
	Titles    []string `xml:"url>news>title"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {

	var s Sitemapindex
	var n News
	url := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map = make(map[string]NewsMap)

}
