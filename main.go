package main

import "fmt"
import "log"
import "net/http"

import "golang.org/x/net/html"
import "golang.org/x/net/html/atom"

/*
Components

- URL Frontier (data structure)
	- 1. Extract URLs.
	- 2. Assign URL to queues using mapping table.
		- Mapping table self init.
	- 3. A process dequeues from a URL queue to download.
		- This process should spawn a worker thread.

- HTML Downloader
*/

const MAX_QUEUES int = 3

// https://datatracker.ietf.org/doc/html/rfc3986#section-3
const scheme string = "http"
const host string = "localhost:8003"
const path string = "en/Main_page.html"

var seed string = fmt.Sprintf("%s://%s/%s", scheme, host, path)

var uf UrlFrontier = NewUrlFrontier()

func main() {
	Start()
	printList(uf.UrlQueues[0].GetElements())
}

func Start() {
	resp, err := http.Get(seed)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" {
					uf.ProcessUrl(a.Val)
					break
				}
			}
		}
	}
}

func printList[T string | int](a []T) {
	for i, m := range a {
		fmt.Printf("Element %d: ", i)
		fmt.Println(m)
	}
}
