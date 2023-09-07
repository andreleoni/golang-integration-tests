package request

import (
	"net/http"

	"gopkg.in/resty.v1"
)

type Request struct{}

func NewRequest() RequestInterface {
	return &Request{}
}

func (r *Request) KeepAlive(url string) bool {
	client := resty.New()

	resp, _ := client.R().Get(url)

	return resp.StatusCode() == http.StatusOK
}
