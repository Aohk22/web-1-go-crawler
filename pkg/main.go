package main

// import "os"
import "fmt"
import "log"
import "regexp"
import "net/http"
import "golang.org/x/net/html"
import "golang.org/x/net/html/atom"

/*
Components

- HTML Downloader
- URL Frontier
*/

// https://datatracker.ietf.org/doc/html/rfc3986#section-3
const scheme string = "http"
const host string = "localhost:8003"
const path string = "en/Main_page.html"
var seed string = fmt.Sprintf("%s://%s/%s", scheme, host, path)

func main() {
	re := regexp.MustCompile(`^(https?://)?([a-zA-Z0-9.-]+){1}/([a-zA-Z0-9_./:-]+)?(\?.+)?(#.+)?`)

	resp, err := http.Get(seed)
	if err != nil { log.Fatal(err) }
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil { log.Fatal(err) }

	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" && re.MatchString(a.Val) {
					submatches := re.FindStringSubmatch(a.Val)
					if submatches[2] == ".." {
						fmt.Printf("%s ", a.Val)
						fmt.Println("Is a relative path")
					} else {
						fmt.Printf("%s ", a.Val)
						fmt.Println("Is a link")
					}
					break
				}
			}
		}
	}
}
