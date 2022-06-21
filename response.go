package request

import (
	"encoding/json"
	"net/http"
)

// Response is response request
type Response struct {
	Detail *http.Response `json:"http,omitempty"`
	Body   []byte         `json:"body,omitempty"`
	Error  *ResponseError `json:"error,omitempty"`
}

func (r *Response) setError(err error) *Response {
	if r.Error == nil {
		r.Error = &ResponseError{}
	}

	r.Error.setError(err)
	return r
}

func (r *Response) setErrorDesc(desc string) *Response {
	if r.Error == nil {
		r.Error = &ResponseError{}
	}

	r.Error.Description = desc
	return r
}

// ResponseError wrap error response
type ResponseError struct {
	Err         error  `json:"err,omitempty"`
	Description string `json:"description,omitempty"`
}

func (re *ResponseError) setError(err error) *ResponseError {
	if err == nil {
		return re
	}

	re.Err = err
	re.Description = err.Error()

	return re
}

// Parse from response data to pointer
func (r *Response) Parse(data interface{}) *Response {
	if r.Body == nil {
		return r
	}

	err := json.Unmarshal(r.Body, data)
	if err != nil {
		r.setError(err).setErrorDesc(ErrParse.Error())
	}

	return r
}
