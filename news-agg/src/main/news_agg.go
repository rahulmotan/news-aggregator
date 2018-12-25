package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	link := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := http.Get(link)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	news_map := make(map[string]NewsMap)

	for _, location := range s.Locations {
		fmt.Println("Reading from", location)
		link := strings.TrimSpace(location)
		r, _ := http.Get(link)
		//		fmt.Println("Error:", err)
		b, _ := ioutil.ReadAll(r.Body)
		//		fmt.Println("Error:", err2)
		xml.Unmarshal(b, &n)
		//		fmt.Println("Keywords: ", n.Keywords)
		//		fmt.Println("Titles : ", n.Titles)
		//		fmt.Println("Locations : ", n.Locations)
		for index, _ := range n.Keywords {
			//			fmt.Println("Title:", n.Titles[index])
			//			fmt.Println("Link:", n.Locations[index])
			news_map[n.Titles[index]] = NewsMap{n.Keywords[index], n.Locations[index]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}

}
