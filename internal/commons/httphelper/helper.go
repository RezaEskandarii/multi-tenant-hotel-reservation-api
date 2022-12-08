// Package httphelper is helper to manage send http requests
// this package uses from golang builtin http package and supports POST,PUT,DELETE and GET methods.
///
package httphelper

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"reflect"
	"time"
)

type httpHelper struct {
}

type Request struct {
	Ctx     context.Context
	Url     string
	Body    map[string]interface{}
	Headers map[string]string
	Out     interface{}
}

// New creates new instance of httpHelper with url parameter
func New() *httpHelper {

	return &httpHelper{}
}

// Post sends an HTTP Request with post method and returns an error,response body dump and HTTP response,
// this method parameters usage is:
// ctx is base context.
// url is the address to which the Request is to be sent.
// header uses to set http headers and is optional, and it can br nil.
// out is a struct that we want to decode our json response to given struct.
// Example: for this response { total_amount:200000,user_id:369 }
// out struct should be
// type response struct {
//	TotalAmount float64 `json:"total_amount"`
//	UserId      uint64  `json:"user_id"`
// }
func (h *httpHelper) Post(request Request) (error, string, *http.Response) {
	return h.sendHttpRequest(request.Ctx, http.MethodPost, request.Url, request.Body, request.Headers, request.Out)
}

// Get sends an HTTP Request with get method and returns an error,response body dump and HTTP response,
// this method parameters usage is:
// ctx is base context.
// url is the address to which the Request is to be sent.
// header uses to set http headers and is optional, and it can br nil.
// out is a struct that we want to decode our json response to given struct.
// Example: for this response { total_amount:200000,user_id:369 }
// out struct should be
// type response struct {
//	TotalAmount float64 `json:"total_amount"`
//	UserId      uint64  `json:"user_id"`
// }
func (h *httpHelper) Get(request Request) (error, string, *http.Response) {
	return h.sendHttpRequest(request.Ctx, http.MethodGet, request.Url, request.Body, request.Headers, request.Out)
}

// Put sends an HTTP Request with put method and returns an error,response body dump and HTTP response,
// this method parameters usage is:
// ctx is base context.
// url is the address to which the Request is to be sent.
// header uses to set http headers and is optional, and it can br nil.
// out is a struct that we want to decode our json response to given struct.
// Example: for this response { total_amount:200000,user_id:369 }
// out struct should be
// type response struct {
//	TotalAmount float64 `json:"total_amount"`
//	UserId      uint64  `json:"user_id"`
// }
func (h *httpHelper) Put(request Request) (error, string, *http.Response) {
	return h.sendHttpRequest(request.Ctx, http.MethodPut, request.Url, request.Body, request.Headers, request.Out)
}

// Del sends an HTTP Request with delete method and returns an error,response body dump and HTTP response,
// this method parameters usage is:
// ctx is base context.
// url is the address to which the Request is to be sent.
// header uses to set http headers and is optional, and it can br nil.
// out is a struct that we want to decode our json response to given struct.
// Example: for this response { total_amount:200000,user_id:369 }
// out struct should be
// type response struct {
//	TotalAmount float64 `json:"total_amount"`
//	UserId      uint64  `json:"user_id"`
// }
func (h *httpHelper) Del(request Request) (error, string, *http.Response) {
	return h.sendHttpRequest(request.Ctx, http.MethodDelete, request.Url, request.Body, request.Headers, request.Out)
}

// Patch sends an HTTP Request with patch method and returns an error,response body dump and HTTP response,
// this method parameters usage is:
// ctx is base context.
// url is the address to which the Request is to be sent.
// header uses to set http headers and is optional, and it can br nil.
// out is a struct that we want to decode our json response to given struct.
// Example: for this response { total_amount:200000,user_id:369 }
// out struct should be
// type response struct {
//	TotalAmount float64 `json:"total_amount"`
//	UserId      uint64  `json:"user_id"`
// }
func (h *httpHelper) Patch(request Request) (error, string, *http.Response) {
	return h.sendHttpRequest(request.Ctx, http.MethodPatch, request.Url, request.Body, request.Headers, request.Out)
}

// sendHttpRequest sends http Request by given method and parameters.
func (h *httpHelper) sendHttpRequest(ctx context.Context, method string, url string, body map[string]interface{}, header map[string]string, out interface{}) (error, string, *http.Response) {

	// Request body
	var bodyBytes []byte

	// file Request body if given body is not nil.
	if body != nil {
		bytesResult, err := json.Marshal(body)
		if err != nil {
			return err, "", nil
		}
		bodyBytes = bytesResult
	}

	// Request context
	requestCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	r, err := http.NewRequestWithContext(requestCtx, method, url, bytes.NewBuffer(bodyBytes))

	if err != nil {
		return err, "", nil
	}

	// set http headers if given header if not nil
	if header != nil {
		for k, v := range header {
			r.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		return err, "", nil
	}

	// send Request
	defer resp.Body.Close()

	// decode error
	var decodeErr error

	if out != nil {

		if reflect.ValueOf(out).Kind() == reflect.Pointer {
			decodeErr = json.NewDecoder(resp.Body).Decode(out)
		} else {
			decodeErr = json.NewDecoder(resp.Body).Decode(&out)
		}

		if decodeErr != nil {
			return err, "", nil
		}
	}

	bodyDump, err := httputil.DumpResponse(resp, true)

	return err, string(bodyDump), resp

}
