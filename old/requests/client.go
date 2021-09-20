package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var BadGateway = errors.New("bat gateway")

func (c *RequestClient) getUrl(uri string) string {
	return c.Url + uri
}

func (c *RequestClient) NewRequest(req *Request, res *Response) error {
	url := c.getUrl(req.Uri)

	var body bytes.Buffer

	if req.Body != nil {
		body = *req.Body
	} else {
		//var body io.Reader
		if req.JsonBody != nil {
			byteBody, err := json.Marshal(req.JsonBody)
			if err != nil {
				return err
			}
			body = *bytes.NewBuffer(byteBody)
		}
	}

	newReq, err := http.NewRequest(req.Method, url, &body)
	if err != nil {
		return err
	}

	if c.Header != nil {
		newReq.Header = c.Header.Clone()
	}

	for i, values := range req.Header {
		for _, v := range values {
			newReq.Header.Add(i, v)
		}
	}

	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}

	client := &http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Do(newReq)
	if err != nil {
		return err
	}

	res.Code = resp.StatusCode
	defer resp.Body.Close()

	BodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res.BodyRaw = BodyRaw

	if req.Flags.IsBodyString {
		res.Body = string(res.BodyRaw)
	}

	if res.Struct != nil {
		err = json.Unmarshal(res.BodyRaw, &res.Struct)
		if err != nil {
			return err
		}
	}

	return err
}
