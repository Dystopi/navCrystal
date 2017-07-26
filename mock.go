package sslMock

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

/*
	Kudos to krait on SO (http://stackoverflow.com/users/721861/krait) for his answer that inspired this repo. Most of
	this is copy-pasta that has been modified a bit to not effect global configurations and be a bit more dynamic
*/

type Mock struct {
	Client *http.Client
	URL       *url.URL
}

func NewMock(callback func(http.ResponseWriter, *http.Request)) (*Mock, error) {
	server := httptest.NewServer(http.HandlerFunc(callback))
	serverURL, _ := url.Parse(server.URL)
	httpClient := &http.Client{}
	return &Mock{httpClient, serverURL}, nil
}

func (m Mock) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = m.URL.Scheme
	req.URL.Host = m.URL.Host
	rt := m.Client.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}
