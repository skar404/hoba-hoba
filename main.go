package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"time"

	"github.com/skar404/hoba-hoba/libs"
	"github.com/skar404/hoba-hoba/requests"
	"github.com/skar404/hoba-hoba/telegram"
)

type Rss struct {
	XMLName    xml.Name `xml:"rss"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	Atom       string   `xml:"atom,attr"`
	Content    string   `xml:"content,attr"`
	Googleplay string   `xml:"googleplay,attr"`
	Itunes     string   `xml:"itunes,attr"`
	Channel    struct {
		Text string `xml:",chardata"`
		Link []struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Rel   string `xml:"rel,attr"`
			Title string `xml:"title,attr"`
			Type  string `xml:"type,attr"`
			Xmlns string `xml:"xmlns,attr"`
		} `xml:"link"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Description   string `xml:"description"`
		Copyright     string `xml:"copyright"`
		Language      string `xml:"language"`
		PubDate       string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		Image         struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Link  string `xml:"link"`
			Title string `xml:"title"`
			URL   string `xml:"url"`
		} `xml:"image"`
		Type       string `xml:"type"`
		Summary    string `xml:"summary"`
		Author     string `xml:"author"`
		Explicit   string `xml:"explicit"`
		NewFeedURL string `xml:"new-feed-url"`
		Keywords   string `xml:"keywords"`
		Owner      struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name"`
			Email string `xml:"email"`
		} `xml:"owner"`
		Category []struct {
			Text     string `xml:",chardata"`
			AttrText string `xml:"text,attr"`
			Category struct {
				Text     string `xml:",chardata"`
				AttrText string `xml:"text,attr"`
			} `xml:"category"`
		} `xml:"category"`
		Item []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Text string `xml:",chardata"`
	Guid struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Author      string `xml:"author"`
	Link        string `xml:"link"`
	Encoded     string `xml:"encoded"`
	Enclosure   struct {
		Text   string `xml:",chardata"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
		URL    string `xml:"url,attr"`
	} `xml:"enclosure"`
	Duration    string `xml:"duration"`
	Summary     string `xml:"summary"`
	Subtitle    string `xml:"subtitle"`
	Explicit    string `xml:"explicit"`
	EpisodeType string `xml:"episodeType"`
	Episode     string `xml:"episode"`
}

var FileClient = requests.RequestClient{
	Url:     "",
	Timeout: 10 * time.Second,
}

func main() {
	res := requests.Response{}
	{
		feedClient := requests.RequestClient{
			Url:     "https://feeds.simplecast.com/jWytY2EF",
			Timeout: 1 * time.Second,
			Header: map[string][]string{
				"Content-Type": {"application/json"},
				"charset":      {"utf-8"},
			},
		}

		req := requests.Request{
			Flags: requests.RequestFlags{
				IsBodyString: true,
			},
		}

		err := feedClient.NewRequest(&req, &res)
		if err != nil {
			panic(err)
		}
	}

	rssFeed := Rss{}
	{
		err := xml.Unmarshal(res.BodyRaw, &rssFeed)
		if err != nil {
			panic(err)
		}
	}

	// Send to telegram

	for _, v := range rssFeed.Channel.Item {

		err := createPost(v)
		if err != nil {
			log.Panicf("error send post v=%+v err=%s", v, err)
		}

		//panic("")
	}

}

func downloadAudioFile(url string) ([]byte, error) {
	//req := requests.Request{
	//	Method: http.MethodGet,
	//	Uri:    url,
	//}
	//res := requests.Response{}
	//err := FileClient.NewRequest(&req, &res)
	//return res.BodyRaw, err

	return ioutil.ReadFile("file.mp3")
}

func createPost(v RssItem) error {
	chatId := -1001497299213

	//file, err := downloadAudioFile(v.Enclosure.URL)
	//if err != nil {
	//	return err
	//}
	//log.Printf("download audio file url=%s", v.Enclosure.URL)
	var messageId = 0
	//messageId, err := telegram.SendAudio(chatId, "Хоба", file)
	//if err != nil {
	//	return err
	//}
	//log.Printf("send audio")

	//validHtml, err := libs.HtmlToMarkdown(v.Description)
	//if err != nil {
	//	return err
	//}

	//err = telegram.SendMessage(chatId, validHtml, messageId, "")
	//if err != nil {
	//	return err
	//}

	err := telegram.SendMessage(chatId, libs.ValidateHTML(v.Description), messageId, "Markdown")

	log.Printf("send message")
	return err
}
