package main

import "log"
import "net/http"

import "golang.org/x/net/html"
import "golang.org/x/net/html/atom"

type HtmlDownloader struct {
	queueIdx uint8
}

func NewHtmlDownloader() HtmlDownloader {
	return HtmlDownloader{0}
}

/* Dequeue. */
func (hd *HtmlDownloader) GetDownloadUrl(pQueues *([]Queue[string])) (url string, err error) {
	queues := *pQueues
	queue := &(queues[hd.queueIdx])

	url, err = (*queue).Dequeue()
	return
}

func (hd *HtmlDownloader) DownloadAPage(pQueues *([]Queue[string])) (fromUrl string, urls []string, err error) {
	fromUrl, err = hd.GetDownloadUrl(pQueues)
	if err != nil {
		return fromUrl, urls, err
	}

	log.Printf("GET %s", fromUrl)
	resp, err := http.Get(fromUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	urls = parseLinks(doc)
	log.Print("Got links: ")
	log.Println(urls)

	hd.queueIdx = (hd.queueIdx + 1) % MAX_QUEUES

	return fromUrl, urls, nil
}

func parseLinks(doc *html.Node) (links []string) {
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
	}
	return links
}
