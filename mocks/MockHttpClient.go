package mocks

import "net/http"

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var GetDoFunc func(req *http.Request) (*http.Response, error)

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
