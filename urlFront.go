package main

import (
	"log"
	"regexp"
	"errors"
	"strings"
)

const (
	UrlTypeAbs = iota
	UrlTypeRel
)

type Uri struct {
	Full     string
	Schema   string
	Host     string
	Path     string
	Query    string
	Frag     string
	PathType int
}

type UrlFrontier struct {
	Urls []string
	HostMap map[string]uint8
	queueIdx uint8
}

func NewUrlFrontier() UrlFrontier {
	return UrlFrontier{make([]string, 0, 1), make(map[string]uint8, 1), 0}
}

func (uf *UrlFrontier) CreateMapping(hostname string) (uint8) {
	log.Printf("Creating mapping for %s - %d", hostname, uf.queueIdx)
	uf.HostMap[hostname] = uf.queueIdx
	uf.queueIdx = (uf.queueIdx + 1) % uint8(MAX_QUEUES)
	return uf.HostMap[hostname]
}

func (uf *UrlFrontier) ProcessUrl(parentUrl string, url string, pQueues *([]Queue[string])) error {
	uri := Uri{}

	url = strings.ReplaceAll(url, "\\", "/")
	uri, err := parseUrl(parentUrl, url)
	if err != nil {
		log.Printf("Skipping URL: %s", url)
		return err
	}

	queueIdx, ok := uf.HostMap[uri.Host]

	if !ok {
		queueIdx = uf.CreateMapping(uri.Host)
	} 

	switch uri.PathType {
	case UrlTypeRel:
		queues := *pQueues
		queue := &(queues[queueIdx])
		if queue.Exists(uri.Full) == false {
			log.Printf("Enqueueing: %s", uri.Full)
			(*queue).Enqueue(uri.Full)
		}
	case UrlTypeAbs:
		// queues := *pQueues
		// queue := &(queues[queueIdx])
		// if queue.Exists(uri.Full) == false {
		// 	log.Printf("Enqueueing: %s", uri.Full)
		// 	(*queue).Enqueue(uri.Full)
		// }
	default:
		return errors.New("Undefined URL type.")
	}

	return nil
}

/*
Parsing is heavily dependent on regular expression. 
[0] = full path
[1] = schema
[2] = host
[3] = path
[4] = query
[5] = fragment
*/
func parseUrl(parentUrl string, url string) (uri Uri, err error) {
	regexUri := regexp.MustCompile(`^(https?://)?((?:[a-zA-Z0-9\-:]\.?)+)?(\.{0,2}[a-zA-Z0-9_./:\-]+)?(\?.*)?(#.*)?`)
	fields := regexUri.FindStringSubmatch(url)
	parentFields := regexUri.FindStringSubmatch(parentUrl)
	// for i, a := range fields {
	// 	log.Printf("[%d] %s", i, a)
	// }
	if len(fields) != 6 {
		return uri, errors.New("No matches.")
	}

	uri.Full = fields[0]
	uri.Schema = fields[1]
	uri.Host = fields[2]
	uri.Path = fields[3]
	uri.Query = fields[4]
	uri.Frag = fields[5]

	if uri.Host == "" && uri.Path == "" && uri.Schema == "" {
		return uri, errors.New("Parse error: no host nor path nor schema.")
	} else if uri.Host == "" {
		uri.PathType = UrlTypeRel
		uri.Schema = parentFields[1]
		uri.Host = parentFields[2]
		uri.Full = uri.Schema + uri.Host + "/" + uri.Path
		// log.Printf("Reconstructed path %s", uri.Full)
	} else {
		if uri.Schema == "" {
			uri.Schema = "http://"
			uri.Full = uri.Schema + uri.Full
		}
		uri.PathType = UrlTypeAbs
	}

	return uri, nil
}
