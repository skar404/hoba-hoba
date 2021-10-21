package rss

import (
	"encoding/xml"
	"github.com/skar404/hoba-hoba/global"
	"time"

	"github.com/skar404/hoba-hoba/requests"
)

var Feed = requests.RequestClient{
	Url:     global.FeedUrl,
	Timeout: 1 * time.Second,
	Header: map[string][]string{
		"Content-Type": {"application/json"},
		"charset":      {"utf-8"},
	},
}

func GetFeed() (*Rss, error) {
	var rssFeed *Rss

	req := requests.Request{
		Flags: requests.RequestFlags{
			IsBodyString: true,
		},
	}
	res := requests.Response{}
	err := Feed.NewRequest(&req, &res)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(res.BodyRaw, &rssFeed)
	if err != nil {
		return nil, err
	}

	return rssFeed, err
}
