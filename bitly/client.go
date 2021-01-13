package bitly

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/skar404/hoba-hoba/requests"
)

var Client = requests.RequestClient{
	Url:     "https://api-ssl.bitly.com/v4/shorten",
	Timeout: 10 * time.Second,
	Header: map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", os.Getenv("SHORT_LINK_TOKEN"))},
	},
}

func CreateLink(link string) (string, error) {
	r := CreateLinkRes{}
	req := requests.Request{
		Method: http.MethodPost,
		JsonBody: &CreateLinkReq{
			GroupGuid: "",
			Domain:    "bit.ly",
			LongUrl:   link,
		},
		Flags: requests.RequestFlags{
			IsBodyString: true,
		},
	}
	res := requests.Response{Struct: &r}
	err := Client.NewRequest(&req, &res)
	if err != nil {
		return "", err
	}

	if res.Code != http.StatusOK && res.Code != http.StatusCreated {
		return "", fmt.Errorf("not valid code %+v, %w", res.Body, requests.BadGateway)
	}

	return r.Link, nil
}
