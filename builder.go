package requestbuidler

import (
	"bytes"
	"net/http"
)

type basicAuth struct {
	user     string
	password string
}

type request struct {
	url       string
	path      string
	method    string
	headers   map[string]string
	body      []byte
	basicAuth *basicAuth
}

func (r *request) Execute(client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest(r.method, r.path, bytes.NewBuffer(r.body))
	if err != nil {
		return nil, err
	}

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	if r.basicAuth != nil {
		req.SetBasicAuth(r.basicAuth.user, r.basicAuth.password)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func newRequest() *request {
	return &request{
		headers: map[string]string{},
	}
}

type RequestBuilder struct {
	url     string
	request *request
}

func New(url string) *RequestBuilder {
	return &RequestBuilder{
		url:     url,
		request: newRequest(),
	}
}

func (b *RequestBuilder) SetMethod(method string) *RequestBuilder {
	b.request.method = method

	return b
}

func (b *RequestBuilder) SetPath(path string) *RequestBuilder {
	b.request.path = b.url + path

	return b
}

func (b *RequestBuilder) AddHeader(key, value string) *RequestBuilder {
	b.request.headers[key] = value

	return b
}

func (b *RequestBuilder) SetHeaders(headers map[string]string) *RequestBuilder {
	b.request.headers = headers

	return b
}

func (b *RequestBuilder) SetContentTypeJson() *RequestBuilder {
	return b.AddHeader("Content-Type", "application/json")
}

func (b *RequestBuilder) SetContentURLEncoded() *RequestBuilder {
	return b.AddHeader("Content-Type", "application/x-www-form-urlencoded")
}

func (b *RequestBuilder) SetBasicAuth(user, password string) *RequestBuilder {
	b.request.basicAuth = &basicAuth{
		user:     user,
		password: password,
	}

	return b
}

func (b *RequestBuilder) SetBody(body []byte) *RequestBuilder {
	b.request.body = body

	return b
}

func (b *RequestBuilder) Build() *request {
	req := b.request
	b.request = newRequest()

	return req
}
