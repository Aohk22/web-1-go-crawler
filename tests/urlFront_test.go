package test

import "testing"
import "github.com/Aohk22/web-1-go-crawler"

func TestUrlParse(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"Full params", "http://localhost:123/path/?query=aa23#fragment", "localhost:123"},
		{"Single frag", "#", ""},
		{"Relative path", "../relative/path", ".."},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := Url{}
			url.Parse(test.input)
			actual := url.Host
			if actual != test.want {
				t.Errorf("got %s, want %s", actual, test.want)
			}
		})
	}
}

func TestUrlFrontierCreate(t *testing.T) {
	uf := NewUrlFrontier()
	queueList := uf.UrlQueues
	if len(queueList) != MAX_QUEUES {
		t.Errorf("MAX_QUEUES(%d) != len(queues)(%d)", MAX_QUEUES, len(queueList))
	}
}
