package rss

import "encoding/xml"

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
		Item []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
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
	Season      string `xml:"season"`
}
