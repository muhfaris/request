package request

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Header is custom header
type Header map[string]string

// Query for querystring
type Query map[string]string

// ToQuery convert from other type to paramQuery type
func ToQuery(params interface{}) Query {
	var paramQuery = Query{}
	for k, v := range params.(map[string]string) {
		paramQuery[k] = v
	}
	return paramQuery
}

// GET is request
func (c *Config) Get() *Response {
	var res *Response

	if err := c.setMethod(http.MethodGet); err != nil {
		return res.setError(err)
	}

	request, err := c.newRequest()
	if err != nil {
		return res.setError(err)
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	request = buildQuery(request, c.QueryString)

	return c.send(request)
}

// POST is request
func (c *Config) Post() *Response {
	var res *Response

	if err := c.setMethod(http.MethodPost); err != nil {
		return res.setError(err)
	}

	request, err := c.newRequest()
	if err != nil {
		return res.setError(err)
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

// DELETE is request
func (c *Config) Delete() *Response {
	var res *Response

	if err := c.setMethod(http.MethodDelete); err != nil {
		return res.setError(err)
	}

	request, err := c.newRequest()
	if err != nil {
		return res.setError(err)
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

// PATCH is request
func (c *Config) Patch() *Response {
	var res *Response

	if err := c.setMethod(http.MethodPatch); err != nil {
		return res.setError(err)
	}

	request, err := c.newRequest()
	if err != nil {
		return res.setError(err)
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

// Put is request
func (c *Config) Put() *Response {
	var res *Response

	if err := c.setMethod(http.MethodPut); err != nil {
		return res.setError(err)
	}

	request, err := c.newRequest()
	if err != nil {
		return res.setError(err)
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

func (c *Config) send(r *http.Request) *Response {
	var res *Response

	for {
		r.Header.Set(string(ContentTypeHeader), c.ContentType)
		// set user agent
		if c.UserAgent != "" {
			r.Header.Set(string(UserAgentHeader), c.UserAgent)
		}

		// check authorization
		if c.Authorization != "" {
			r.Header.Add(string(AuthorizationHeader), c.Authorization)
		}

		resp, err := c.httpClient.Do(r)
		// case retry
		if c.onRetry() {
			// if success
			if err == nil {
				defer resp.Body.Close()
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return res.setError(err).setErrorDesc(ErrReadData.Error())
				}

				// Restore the io.ReadCloser to its original state
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))
				return &Response{Detail: resp, Body: data}
			}

			// if error
			select {
			case <-r.Context().Done():
				return res.setError(err).setErrorDesc(ErrContextDone.Error())

			case <-time.After(c.Delay):
				c.Retry--
			}

			continue
		}

		if err != nil {
			return res.setError(err).setErrorDesc(ErrServer.Error())
		}

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return res.setError(err).setErrorDesc(ErrReadData.Error())
		}

		if resp.StatusCode > 300 {
			err := errors.New(http.StatusText(resp.StatusCode))
			return res.setError(err)
		}

		// Restore the io.ReadCloser to its original state
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		return &Response{Detail: resp, Body: data}
	}
}

// HTTPClient is interface
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}
