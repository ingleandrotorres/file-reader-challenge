package gateways

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Response struct {
	response *http.Response
	body     []byte
}

func NewResponse(response *http.Response, body []byte) *Response {
	return &Response{response: response, body: body}
}

func (r *Response) GetBody() string {
	return string(r.body)
}

func (r *Response) GetResponse() *http.Response {
	return r.response
}

func (r *Response) GetStatusCode() int {
	return r.response.StatusCode
}

func (r *Response) GetStatus() string {
	return http.StatusText(r.response.StatusCode)
}

func (r *Response) GetHeaders() *http.Header {
	return &r.response.Header
}

func (r *Response) GetHeader(name string) []string {
	return r.response.Header[name]
}

func (r *Response) StatusCodeIs2xx() bool {
	return r.response.StatusCode >= 200 && r.response.StatusCode <= 299
}

func (r *Response) StatusCodeIs4xx() bool {
	return r.response.StatusCode >= 400 && r.response.StatusCode <= 499
}

func (r *Response) StatusCodeIs5xx() bool {
	return r.response.StatusCode >= 500 && r.response.StatusCode <= 599
}
func (r *Response) Hydrate(fill interface{}) error {
	contentTypeHeaders := r.GetHeader("Content-Type")
	if len(contentTypeHeaders) == 0 {
		return r.StructCouldNotBeHydrated(fill, r.GetBody())
	}

	contentTypeHeader := contentTypeHeaders[0]
	var err error

	if strings.Contains(contentTypeHeader, "application/json") {
		if err = json.Unmarshal(r.body, fill); err == nil {
			return nil
		}
	}

	if strings.Contains(contentTypeHeader, "application/xml") {
		if err = xml.Unmarshal(r.body, fill); err == nil {
			return nil
		}
	}

	return r.StructCouldNotBeHydrated(fill, r.GetBody(), err)
}
func (r *Response) StructCouldNotBeHydrated(typ interface{}, detail interface{}, stacktrace ...error) error {
	structType := reflect.TypeOf(typ).Elem()
	message := fmt.Sprintf("'%s' could not be hydrated with: '%s'", structType.Name(), detail)
	return errors.New(message)
}
