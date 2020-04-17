package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// ParamQuery for querystring
type ParamQuery map[string]string

// ReqApp is request application
type ReqApp struct {
	URL           string
	ContentType   string
	Body          []byte
	Authorization string
	QueryString   ParamQuery

	httpClient http.Client
}

// ReqResponse is response request
type ReqResponse struct {
	Response *http.Response
	Body     []byte
}

// HTTPClient is interface
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
}

// New is initialize
func New(url, ct, au string, body interface{}, qs ParamQuery) (*ReqApp, error) {
	b, err := validationBody(body)
	if err != nil {
		return nil, err
	}

	return &ReqApp{
		URL:           url,
		ContentType:   ct,
		Body:          b,
		Authorization: au,
		QueryString:   qs,
	}, nil
}

func validationBody(body interface{}) ([]byte, error) {
	r := reflect.TypeOf(body)
	log.Println(r.Kind())
	if r.Kind() == reflect.Uint8 || r.Kind() == reflect.Slice {
		return body.([]byte), nil
	}
	return BodyByte(body)
}

// BodyByte is build body data to byte
func BodyByte(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Error BodyByte:%v", err)
	}
	return b, nil
}

// GET is request
func (app *ReqApp) GET() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodGet, app.URL, nil)
	if err != nil {
		return nil, err
	}

	request = buildQuery(request, app.QueryString)
	return app.send(request)
}

// POST is request
func (app *ReqApp) POST() (*ReqResponse, error) {
	log.Println(app.URL)
	log.Println(string(app.Body))
	request, err := http.NewRequest(http.MethodPost, app.URL, bytes.NewBuffer(app.Body))
	if err != nil {
		return nil, err
	}

	return app.send(request)

}

// DELETE is request
func (app *ReqApp) DELETE() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodDelete, app.URL, bytes.NewBuffer(app.Body))
	if err != nil {
		return nil, err
	}

	return app.send(request)
}

// PATCH is request
func (app *ReqApp) PATCH() (*ReqResponse, error) {
	request, err := http.NewRequest(http.MethodPatch, app.URL, bytes.NewBuffer(app.Body))
	if err != nil {
		return nil, err
	}

	return app.send(request)
}

func buildQuery(request *http.Request, querystring ParamQuery) *http.Request {
	if querystring == nil {
		return request
	}

	q := request.URL.Query()
	for k, v := range querystring {
		q.Add(k, v)
	}

	request.URL.RawQuery = q.Encode()
	return request
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
		Body:     data,
		Response: resp,
	}, nil
}
