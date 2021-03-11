package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CustomHeader is custom header
type CustomHeader map[string]string

// ParamQuery for querystring
type ParamQuery map[string]string

// ToParamQuery convert from other type to paramQuery type
func ToParamQuery(params interface{}) ParamQuery {
	var paramQuery ParamQuery
	for k, v := range params.(map[string]string) {
		paramQuery[k] = v
	}
	return paramQuery
}

// ReqApp is request application
type ReqApp struct {
	URL           string
	ContentType   string
	Body          []byte
	Authorization string
	QueryString   map[string]string
	Headers       CustomHeader

	httpClient http.Client
}

func (ra *ReqApp) ChangeURL(url string) error {
	if url == "" {
		return fmt.Errorf("url is empty")
	}

	ra.URL = url
	return nil
}

func (ra *ReqApp) ChangeBody(body []byte) error {
	if body == nil {
		return fmt.Errorf("body data is empty")
	}

	ra.Body = body
	return nil
}

func (ra *ReqApp) ChangeAuthorization(authorization string) error {
	if authorization == "" {
		return fmt.Errorf("authorization is empty")
	}

	ra.Authorization = authorization
	return nil
}

func (ra *ReqApp) ChangeHeaders(headers CustomHeader) error {
	if headers == nil {
		return fmt.Errorf("authorization is empty")
	}

	ra.Headers = headers
	return nil
}

// GET is request
func (app *ReqApp) GET() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodGet, app.URL, nil)
	if err != nil {
		return nil, err
	}

	if app.Headers != nil {
		request = buildHeader(request, app.Headers)
	}

	request = buildQuery(request, app.QueryString)

	return app.send(request)
}

// POST is request
func (app *ReqApp) POST() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodPost, app.URL, bytes.NewBuffer(app.Body))
	if err != nil {
		return nil, err
	}

	if app.Headers != nil {
		request = buildHeader(request, app.Headers)
	}

	return app.send(request)
}

// DELETE is request
func (app *ReqApp) DELETE() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodDelete, app.URL, bytes.NewBuffer(app.Body))
	if err != nil {
		return nil, err
	}

	if app.Headers != nil {
		request = buildHeader(request, app.Headers)
	}

	return app.send(request)
}

// PATCH is request
func (app *ReqApp) PATCH() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodPatch, app.URL, bytes.NewBuffer(app.Body))
	if err != nil {
		return nil, err
	}

	if app.Headers != nil {
		request = buildHeader(request, app.Headers)
	}

	return app.send(request)
}

func (app *ReqApp) send(r *http.Request) (*ReqResponse, error) {
	r.Header.Set("content-type", app.ContentType)

	// check authorization
	if app.Authorization != "" {
		r.Header.Add("Authorization", app.Authorization)
	}

	resp, err := app.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &ReqResponse{
		HTTP: resp,
		Body: data,
	}, nil
}

// ReqResponse is response request
type ReqResponse struct {
	HTTP *http.Response
	Body []byte
}

// HTTPClient is interface
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

// New is initialize
func New(url, contentType, authorization string, body interface{}, query map[string]string, headers CustomHeader) (*ReqApp, error) {
	b, err := validationBody(body)
	if err != nil {
		return nil, err
	}

	return &ReqApp{
		URL:           url,
		ContentType:   contentType,
		Body:          b,
		Authorization: authorization,
		QueryString:   query,
		Headers:       headers,
	}, nil
}
