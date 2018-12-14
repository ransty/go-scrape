package main

import (
	"net/http"
	"time"
	"log"
	"fmt"
	
	"golang.org/x/net/html"
)

func main() {
	links, err := ExtractUrls("https://keano.me")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(links)
}

func ExtractUrls(url string) ([]string, error) {
	// Create HTTP client with Timeout
	client := &http.Client {
		Timeout: 30 * time.Second,
	}
	
	// HTTP request to the link
	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to be scraped for links
	body, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Now get the links within the body and search them with the forEachNode method, store in array of strings
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := response.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}

		}
	}
	forEachNode(body, visitNode, nil)
	return links, nil
}	

// Basic traversal of nodes
// See gopl.io/ch5/outline2
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

