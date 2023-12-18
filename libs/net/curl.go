package net

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Response struct {
	PrimaryIndex int
	StatusCode   int
	Header       http.Header
	Body         []byte
}

type Option struct {
	Method  string // http.MethodX
	Params  string
	Body    io.Reader
	Header  http.Header
	Timeout time.Duration
	Must    int
}

/*
CurlOption function
*/
func CurlOption() *Option {
	return &Option{
		Method:  http.MethodGet,
		Header:  http.Header{},
		Timeout: 30 * time.Second,
		Must:    -1,
	}
}

/*
Curl function.
By default, if the option is NIL then timeout is 30 seconds.
*/
func Curl(url string, option *Option) (*Response, error) {

	if option == nil {
		option = CurlOption()
	}

	if strings.HasPrefix(url, "ws") {
		return nil, errors.WithStack(errors.New("the `websocket` protocol is not supported"))
	}
	if strings.HasPrefix(url, "http") {
		//pass
	} else {
		url = "http://" + url
	}

	if len(option.Params) > 0 {
		url += "/" + option.Params
	}

	request, err := http.NewRequest(option.Method, url, option.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for key, arr := range option.Header {
		for i, value := range arr {
			if i == 0 {
				request.Header.Set(key, value)
			} else {
				request.Header.Add(key, value)
			}
		}
	}

	client := &http.Client{
		Timeout: option.Timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buffer, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return &Response{
		StatusCode: response.StatusCode,
		Header:     response.Header.Clone(),
		Body:       buffer,
	}, errors.WithStack(err)
}

/*
SetMethod function
*/
func (ins *Option) SetMethod(method string) error {
	var temp = map[string]bool{
		http.MethodConnect: true,
		http.MethodDelete:  true,
		http.MethodGet:     true,
		http.MethodHead:    true,
		http.MethodOptions: true,
		http.MethodPatch:   true,
		http.MethodPost:    true,
		http.MethodPut:     true,
		http.MethodTrace:   true,
	}

	value, ok := temp[method]
	if ok && value {
		ins.Method = method
		return nil
	}

	return errors.WithStack(errors.New("method invalid"))
}

/*
SetData function
*/
func (ins *Option) SetData(body []byte) {
	n := make([]byte, len(body))
	copy(n, body)
	ins.Body = bytes.NewReader(n)
}
