package request

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type args struct {
		config *Config
	}

	type Anime struct {
		Anime     string `json:"anime"`
		Character string `json:"character"`
		Quote     string `json:"quote"`
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "test response json with parse to struct",
			args: args{
				config: &Config{URL: "https://animechan.vercel.app/api/random"},
			},
			want: &Response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var anime = Anime{}
			got := Get(tt.args.config).Parse(&anime)

			assert.NotEqual(t, nil, got.Error)
			assert.NotNil(t, got.Body)
			assert.NotNil(t, got.Detail.Body)
			assert.NotEqual(t, "", anime.Anime)
			assert.NotEqual(t, "", anime.Character)
			assert.NotEqual(t, "", anime.Quote)
		})
	}
}

func TestPost(t *testing.T) {
	type args struct {
		config *Config
	}

	body, _ := BodyByte(
		map[string]string{
			"name": "faris",
			"job":  "leader",
		},
	)

	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "create new user",
			args: args{
				config: &Config{
					URL:  "https://reqres.in/api/users",
					Body: body,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data map[string]interface{}
			got := Post(tt.args.config).Parse(&data)

			assert.Nil(t, got.Error)
			assert.NotNil(t, got.Body)
			assert.NotNil(t, got.Detail.Body)
			assert.NotEmpty(t, data["id"])
			assert.NotEmpty(t, data["name"])
			assert.NotEmpty(t, data["job"])

		})
	}
}

func TestPostForm(t *testing.T) {
	type args struct {
		config *Config
	}

	data := url.Values{
		"name":       {"John Doe"},
		"occupation": {"gardener"},
	}
	body, _ := BodyByte(data)

	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "create new user",
			args: args{
				config: &Config{
					URL:         "https://httpbin.org/post",
					Body:        body,
					ContentType: MimeTypeFormURL,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data map[string]interface{}
			got := Post(tt.args.config).Parse(&data)

			assert.Nil(t, got.Error)
			assert.NotNil(t, got.Body)
			assert.NotNil(t, got.Detail.Body)
			assert.NotEmpty(t, data["form"])
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		config *Config
	}

	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "delete user",
			args: args{
				config: &Config{
					URL: "https://reqres.in/api/users/7",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Delete(tt.args.config)

			assert.Nil(t, got.Error)
			assert.Equal(t, http.StatusNoContent, got.Detail.StatusCode)
		})
	}
}

func TestPatch(t *testing.T) {
	type args struct {
		config *Config
	}

	body, _ := BodyByte(
		map[string]string{
			"name": "faris",
			"job":  "leader",
		},
	)

	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "patch: create and then update the user",
			args: args{
				config: &Config{
					URL:  "https://reqres.in/api/users",
					Body: body,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data map[string]interface{}
			got := Post(tt.args.config).Parse(&data)

			assert.Nil(t, got.Error)
			assert.NotNil(t, got.Body)
			assert.NotNil(t, got.Detail.Body)
			assert.NotEmpty(t, data["id"])
			assert.NotEmpty(t, data["name"])
			assert.NotEmpty(t, data["job"])

			// change job to programmer
			var update map[string]interface{}
			body, _ := BodyByte(
				map[string]string{
					"name": "faris",
					"job":  "programmer",
				},
			)
			// preparation
			_ = tt.args.config.ChangeBody(body)
			_ = tt.args.config.ChangeURL(fmt.Sprintf("%s/%s", tt.args.config.URL, data["id"]))

			got = Patch(tt.args.config).Parse(&update)
			assert.Nil(t, got.Error)
			assert.NotNil(t, got.Body)
			assert.NotNil(t, got.Detail.Body)
			assert.Equal(t, data["name"], update["name"])
			assert.Equal(t, "programmer", update["job"])
		})
	}
}
