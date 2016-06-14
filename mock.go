package navCrystal

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
}

type MockHTTPSTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func NewMock(callback func(http.ResponseWriter, *http.Request)) (*Mock, error) {
	server := httptest.NewServer(http.HandlerFunc(callback))
	serverURL, _ := url.Parse(server.URL)
	tPort := MockHTTPSTransport{URL: serverURL}
	httpClient := &http.Client{Transport: tPort}
	return &Mock{httpClient}, nil
}

func (m MockHTTPSTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = m.URL.Scheme
	req.URL.Host = m.URL.Host
	rt := m.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}
