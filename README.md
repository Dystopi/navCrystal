# sslMock

sslMock is a GO HTTPS mocking package. The primary use case for this is mocking 3rd party API calls that rely upon https

## Installation
```golang
go get github.com/Dystopi/sslMock
```

## Overview

This is a very basic package that effectivly morphs the URL scheme from HTTPS to HTTP in order to return mock requests.

## Usage

```golang
	mock, err := navCrystal.NewMock({CUSTOM DEFINED REQUEST HANDLER})
	if err != nil {
		fmt.Println("Failed to create Mock Client")
	}
	resp, _ := mock.Client.Get("https://www.example.com")
	// Treat the response as you would anything else
```

In order to handle a method that makes multiple requests to multiple endpoints we can do this :

```golang
func multiReqHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.URL.Path {
	case "/v2/search":
		fmt.Fprintln(w, `{"result":"foo"}`)
	case "/v2/status":
		fmt.Fprintln(w, `{"result":"bar"}`)
	}
}

mock, err := navCrystal.NewMock(multiReqHandler)
if err != nil {
	fmt.Println("Failed to create Mock Client")
}
searchResp, _ := mock.Client.Get("https://www.example.com/v2/search")
statResp, _ := mock.Client.Get("https://www.example.com/v2/status")
```
