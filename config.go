package request

import (
	"errors"
	"fmt"
	"net/http"
)

// Config is request application
type Config struct {
	URL           string
	ContentType   string
	Body          []byte
	Authorization string
	QueryString   map[string]string
	Headers       Header

	// Retry
	Retry int
	Delay int

	httpClient http.Client
}

// New is initialize
func New() *Config {
	return &Config{ContentType: MimeTypeJSON}
}

func (c *Config) reinit(config *Config) *Config {
	c.ChangeURL(config.URL)
	c.ChangeContentType(config.ContentType)
	c.ChangeBody(config.Body)
	c.ChangeAuthorization(config.Authorization)
	c.ChangeQueryString(config.QueryString)
	c.ChangeHeaders(config.Headers)
	c.ChangeRetry(config.Retry)
	c.ChangeDelay(config.Delay)
	return c
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
func (c *Config) ChangeBody(body []byte) error {
	if body == nil {
		return fmt.Errorf("body data is empty")
	}

	c.Body = body
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
func (c *Config) ChangeDelay(delay int) error {
	if delay == 0 {
		return fmt.Errorf("error missing argument of delay retry")
	}

	c.Delay = delay
	return nil
}
