# navCrystal
![Zhaan with the Nav Crystal from DNA Mad Scientist] (http://1.bp.blogspot.com/-NS2_TUS5mVg/TjDFkbnF6PI/AAAAAAAAOTU/2fM_KioX5J8/s320/mesmerized.JPG)

navCrystal is a GO HTTPS mocking package

## Installation
```golang
go get github.com/Dystopi/navCrystal
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
