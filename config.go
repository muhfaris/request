package request

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	defaultRetryDelay = 1 * time.Second
)

// Config is request application
//
// Retry Mechanism
// You can implementation retry mechanism,
// set the retry and set delay as interval time for distance to the next request
type Config struct {
	URL           string
	ContentType   string
	Body          []byte
	Authorization string
	QueryString   map[string]string
	Headers       Header

	// Retry
	Retry int
	Delay time.Duration

	// UserAgent
	UserAgent string

	httpClient http.Client
	method     string
	Context    *context.Context
}

// New is initialize
func New() *Config {
	return &Config{ContentType: MimeTypeJSON, Delay: defaultRetryDelay}
}

func (c *Config) ChangeURL(url string) error {
	if url == "" {
		return fmt.Errorf("url is empty")
	}

	c.URL = url
	return nil
}

// ChnageContentType is change content type
func (c *Config) ChangeContentType(contentType string) error {
	if contentType == "" {
		return errors.New("error missing argument of content-type")
	}
	c.ContentType = contentType
	return nil
}

// ChangeBody is change body
func (c *Config) ChangeBody(body interface{}) error {
	if body == nil {
		return fmt.Errorf("body data is empty")
	}

	b, err := validationBody(body)
	if err != nil {
		return err
	}

	c.Body = b
	return nil
}

// ChangeAuthorization is change authorization request
func (c *Config) ChangeAuthorization(authorization string) error {
	if authorization == "" {
		return fmt.Errorf("authorization is empty")
	}

	c.Authorization = authorization
	return nil
}

// ChangeQueryString is change params of query string
func (c *Config) ChangeQueryString(qs map[string]string) error {
	if len(qs) == 0 {
		return fmt.Errorf("error missing argument of query string")
	}

	c.QueryString = qs
	return nil
}

// ChangeHeaders is change header request
func (c *Config) ChangeHeaders(headers Header) error {
	if headers == nil {
		return fmt.Errorf("authorization is empty")
	}

	c.Headers = headers
	return nil
}

// ChangeRetry is change total try to retry
func (c *Config) ChangeRetry(retry int) error {
	if retry == 0 {
		return fmt.Errorf("error missing argument of retry")
	}

	c.Retry = retry
	return nil
}

// ChangeDelay is change delay Retry
func (c *Config) ChangeDelay(delay time.Duration) error {
	if delay == 0 {
		return fmt.Errorf("error missing argument of delay retry")
	}

	c.Delay = delay
	return nil
}

// ChangeUserAgent is change user agent request
func (c *Config) ChangeUserAgent(userAgent string) error {
	if userAgent == "" {
		return fmt.Errorf("error user agent is empty")
	}

	c.UserAgent = userAgent
	return nil
}

// Params is wrap query string
func (c *Config) Params(params map[string]string) *Config {
	if len(params) == 0 {
		return c
	}
	c.QueryString = params
	return c
}

// onRetry is check the request use retry mechanism
func (c *Config) onRetry() bool {
	return c.Retry > 0
}

// validate is validate the config request
func (c *Config) validate() *Response {
	var res *Response

	if c == nil {
		err := &ResponseError{
			Err:         ErrConfig,
			Description: ErrValidate.Error(),
		}

		res.Error = err
	}

	return nil
}

// newRequest is create initialize request
// default the context is background
func (c *Config) newRequest() (*http.Request, error) {
	payload := bytes.NewBuffer(c.Body)

	if c.Context != nil {
		return http.NewRequestWithContext(*c.Context, c.method, c.URL, payload)
	}

	return http.NewRequest(c.method, c.URL, payload)
}

func (c *Config) setMethod(method string) error {
	if method == "" {
		return ErrMethod
	}

	c.method = method
	return nil
}
