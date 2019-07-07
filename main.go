package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var linksList = make(chan []string, 15)
var visited = make(map[string]bool)
var allLinks []string

func setFirstLink(args []string) {
	linksList <- args[1:]
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You should provide an url to be processed")
	} else {
		counter := 1
		baseURL := os.Args[1]
		go setFirstLink(os.Args)
		for ; counter > 0; counter-- {
			list := <-linksList
			for _, link := range list {
				if !visited[link] {
					counter++
					visited[link] = true
					fmt.Println("checking the link:", link)
					go visitLink(link, baseURL)
				}
			}
		}
		sort.Strings(allLinks)
		for index, link := range allLinks {
			fmt.Println(index, link)
		}
	}
}

func visit(link string, baseURL string) []string {
	allLinks = append(allLinks, link)
	list := returnAllLinks(link, baseURL)
	return list
}

func visitLink(link string, baseURL string) {
	linksList <- visit(link, baseURL)
}

func retrieveURLBody(url string) io.ReadCloser {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error or retrieving the url:", url)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error on parsing the url:", url, err.Error())
	}
	return response.Body
}

func returnAllLinks(url string, urlBase string) []string {
	var linksFound []string
	body := retrieveURLBody(url)
	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return linksFound
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					link := formatURL(url, attr.Val)
					linksFound = addLinkToList(link, linksFound)
					linksFound = returnLocalLinks(urlBase, linksFound)
				}
			}
		}
	}
}

func formatURL(base string, link string) string {
	linkURL, err := url.Parse(link)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uriFormatted := baseURL.ResolveReference(linkURL)
	return uriFormatted.String()
}

func returnLocalLinks(baseURL string, links []string) (localLinks []string) {
	var ret []string
	for _, link := range links {
		if strings.HasPrefix(link, baseURL) {
			ret = append(ret, link)
		}
	}
	return ret
}

func addLinkToList(link string, linksFound []string) []string {
	linksFound = append(linksFound, removeLastSlash(trimHash(link)))
	return linksFound
}

func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

func removeLastSlash(l string) string {
	if strings.HasSuffix(l, "/") {
		index := len(l) - 1
		return l[:index]
	}
	return l
}
