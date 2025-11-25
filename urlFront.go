package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	UrlTypeAbs = iota
	UrlTypeRel
)

var regexUri *regexp.Regexp = regexp.MustCompile(`^(https?://)?([a-zA-Z0-9.\-:]+)/([a-zA-Z0-9_./:-]+)?(\?.+)?(#.+)?`)

type Url struct {
	Full     string
	Schema   string
	Host     string
	Path     string
	Query    string
	Frag     string
	PathType int
}

func (u *Url) Parse(s string) error {
	fields := regexUri.FindStringSubmatch(s)
	if len(fields) != 6 {
		return errors.New("No matches.")
	}

	u.Full = fields[0]
	u.Schema = fields[1]
	u.Host = fields[2]
	u.Path = fields[3]
	u.Query = fields[4]
	u.Frag = fields[5]

	if u.Host == ".." {
		u.PathType = UrlTypeRel
	} else {
		u.PathType = UrlTypeAbs
	}

	return nil
}

type UrlFrontier struct {
	UrlQueues  []Queue[string]
	UrlMap     map[string]uint8
	MapRrIndex uint8 // Round robin.
}

func NewUrlFrontier() UrlFrontier {
	// queue list pre init.
	return UrlFrontier{make([]Queue[string], 3, MAX_QUEUES), make(map[string]uint8, 1), 0}
}

func (uf *UrlFrontier) GetQueueMap(hostname string) uint8 {
	_, ok := uf.UrlMap[hostname]
	if !ok {
		uf.UrlMap[hostname] = uf.MapRrIndex
		uf.MapRrIndex = (uf.MapRrIndex + 1) % uint8(MAX_QUEUES)
	}
	return uf.UrlMap[hostname]
}

func (uf *UrlFrontier) ProcessUrl(uri string) error {
	url := Url{}

	uri = strings.ReplaceAll(uri, "\\", "/")
	err := url.Parse(uri)

	if err == nil {
		switch url.PathType {
		case UrlTypeRel:
			queueIndex := uf.GetQueueMap(url.Host)
			if int(queueIndex) >= len(uf.UrlQueues) {
				return errors.New("Index out of range.")
			}
			uf.UrlQueues[queueIndex].Enqueue(url.Full)
		case UrlTypeAbs:
		default:
			return errors.New("Undefined URL type.")
		}
	} else {
		log.Printf("Skipping string %s", uri)
	}
	return nil
}

func printFields(list []string) {
	for i, m := range list {
		fmt.Printf("Element %d", i)
		fmt.Println(m)
	}
}
