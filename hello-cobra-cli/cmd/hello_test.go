package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ClientMock struct {
	Response map[string][]byte
}

func (m *ClientMock) Get(url string) (resp *http.Response, err error) {
	body := m.Response[url]

	response := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header:     make(http.Header),
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}

	return response, nil
}

func TestGetResponseBody(t *testing.T) {

	mock := &ClientMock{
		Response: map[string][]byte{
			"https://api.github.com/emojis": []byte("Hello, I'm pretending to be desired response!"),
		},
	}

	service := &HttpService{
		HTTPClient: mock,
	}

	response := getResponseBody(service, "https://api.github.com/emojis")

	assert.EqualValues(t, response, "Hello, I'm pretending to be desired response!")
}
