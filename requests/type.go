package requests

import (
	"net/http"
	"net/url"
	"time"
)

type Request struct {
	Method    string
	Uri       string
	JsonBody  interface{}
	UrlValues url.Values
	Header    http.Header

	Flags RequestFlags
}

type RequestFlags struct {
	IsBodyString bool
}

type RequestClient struct {
	Url     string
	Timeout time.Duration
	Header  http.Header
}

type Response struct {
	Code int

	BodyRaw []byte
	Body    string

	Struct interface{}

	Header http.Header
}
