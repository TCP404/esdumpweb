package testkit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type req struct {
	method string
	u      url.URL
}
type GetReq struct {
	req
	values url.Values
	cookie *http.Cookie
}

func NewGetReq(path string, query map[string]string) *GetReq {
	val := url.Values{}
	for k, v := range query {
		val.Add(k, v)
	}
	u := url.URL{Path: path, RawQuery: val.Encode()}
	return &GetReq{req: req{method: http.MethodGet, u: u}, values: val}
}

func (r *GetReq) Cookie(name, value string) *GetReq {
	r.cookie = &http.Cookie{Name: name, Value: value}
	return r
}

func (r *GetReq) Req() (*http.Request, error) {
	request, err := http.NewRequest(r.method, r.String(), nil)
	if err != nil {
		return nil, err
	}
	if r.cookie != nil {
		request.AddCookie(r.cookie)
	}
	return request, nil
}

func (r *GetReq) String() string {
	return r.u.String()
}

type PostReq struct {
	req
	body map[string]any
}

func NewPostReq(path string, body map[string]any) *PostReq {
	u := url.URL{Path: path}
	return &PostReq{req: req{method: http.MethodPost, u: u}, body: body}
}

func (r *PostReq) Req() (*http.Request, error) {
	body, err := json.Marshal(r.body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(r.method, r.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (r *PostReq) String() string {
	return r.u.String()
}
