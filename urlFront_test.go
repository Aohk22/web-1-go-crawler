package main

import "testing"

func TestUrlParsePathType(t *testing.T) {
	var tests = []struct {
		name  string
		inputParent string
		inputUrl string
		want  string
	}{
		{"Full", "http://parentpage.com/", "http://localhost:123/path/?query=aa23#fragment", "http://localhost:123/path/?query=aa23#fragment"},

		{"No query, frag", "http://parentpage.com/", "http://local.host:123/path/com", "http://local.host:123/path/com"},
		{"No scheme, query, frag", "http://parentpage.com/", "localhost.net/path/her-tehe", "http://localhost.net/path/her-tehe"},

		{"Rel 0 dot", "http://parentpage.com/", "/relative/path/here", "http://parentpage.com//relative/path/here"},
		{"Rel 1 dot", "http://parentpage.com/", "./relativepath2", "http://parentpage.com/./relativepath2"},
		{"Rel 2 dots", "http://parentpage.com/", "../here/there", "http://parentpage.com/../here/there"},
		{"Rel 2 dots", "http://parentpage.com/", "../relative/path", "http://parentpage.com/../relative/path"},

		// {"Single frag", "http://parentpage.com/", "#", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uri := Uri{}
			uri.Parse(test.inputParent, test.inputUrl)
			got := uri.Full
			if got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
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
