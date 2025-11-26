package main

import ( 
	"fmt"
	"log"
	"time"
	"strings"
)

/*
Components

- URL Frontier (data structure)
	- 1. Parse URLs.
	- 2. Assign URL to queues using mapping table.
		- Mapping table self assigns.

- HTML Downloader
	- 1. A process dequeues from a URL queue to download.
		- This process should spawn its own worker thread.
*/

/* Customize these. */
const ITERATIONS = 2
const MAX_QUEUES = 3
const (
	SCHEME string = "http"
	HOST string = "localhost:8003"
	PATH string = "en/Main_page.html"
	SEED string = SCHEME + "://" + HOST + "/" + PATH
)

/* Don't customize these. */
var urlQueue []Queue[string] = make([]Queue[string], MAX_QUEUES, MAX_QUEUES)

func main() {
	if MAX_QUEUES > 255 {
		log.Fatal("MAX_QUEUES should not be larger than uint8 which is the type of the queue selector.")
	}

	var uf UrlFrontier = NewUrlFrontier()
	var hd HtmlDownloader = NewHtmlDownloader()

	urlQueue[0].Enqueue(SEED)
	printQueues(urlQueue)

	for range ITERATIONS {
		fromUrl, links, err := hd.DownloadAPage(&urlQueue)
		if err != nil {
			log.Fatal(err)
		}
		for _, link := range links {
			uf.ProcessUrl(fromUrl, link, &urlQueue)
		}
		printQueues(urlQueue)
		time.Sleep(4 * time.Second)
	}
}

// Utility functions.
func printQueues(ql []Queue[string]) {
	// Just print formatting stuff.
	// Init with a loop is better.
	var maxElementLengthPerQueue []int = []int{-1, -1, -1}
	var maxQLength int = len(ql[0].GetElements())
	const minLen = 7
	const margin = 2

	// Find the max values.
	for queueIdx := range MAX_QUEUES {
		// log.Printf("On Queue %d", queueIdx)
		queueSlice := ql[queueIdx].GetElements()
		queueLen := len(queueSlice)
		// Get max queue length.
		if maxQLength < queueLen {
			maxQLength = queueLen
		}
		// Get longest element in each queue.
		if queueLen == 0 {
			maxElementLengthPerQueue[queueIdx] = minLen // Or else.
		} else {
			for _, e := range queueSlice {
				if maxElementLengthPerQueue[queueIdx] < len(e) {
					maxElementLengthPerQueue[queueIdx] = len(e)
				}
			}
		}
	}

	// Print the headers.
	fmt.Printf("%s\n", strings.Repeat("=", 1 + margin + getSumIntArr(maxElementLengthPerQueue) + 2 * margin * len(maxElementLengthPerQueue) + 1))
	fmt.Printf("|")
	for queueIdx := range MAX_QUEUES {
		padding := getPad(maxElementLengthPerQueue[queueIdx], len("Queue n"), margin)
		fmt.Printf("%sQueue %d%s|", 
			strings.Repeat(" ", margin), 
			queueIdx, 
			strings.Repeat(" ", padding))
	}
	fmt.Println("")

	for elementIdx := range maxQLength {
		fmt.Printf("|")
		for queueIdx := range MAX_QUEUES {
			if elementIdx > (len(ql[queueIdx].GetElements())-1) {
				padding := getPad(maxElementLengthPerQueue[queueIdx], len("(empty)"), margin)
				fmt.Printf("%s(empty)%s|", 
					strings.Repeat(" ", margin),
					strings.Repeat(" ", padding))
			} else {
				element := ql[queueIdx].GetElements()[elementIdx]
				padding := getPad(maxElementLengthPerQueue[queueIdx], len(element), margin)
				fmt.Printf("%s%s%s|",
					strings.Repeat(" ", margin),
					element, 
					strings.Repeat(" ", padding))
			}
		}
		fmt.Println("")
	}
	fmt.Printf("%s\n", strings.Repeat("=", 1 + margin + getSumIntArr(maxElementLengthPerQueue) + 2 * margin * len(maxElementLengthPerQueue) + 1))
}

func getPad(maxLen, strLen, margin int) (padding int) {
	padding = maxLen - strLen + margin
	return
}

func getSumIntArr(arr []int) int {
	var sum int = 0
	for _, m := range arr {
		sum += m
	}
	return sum
}
