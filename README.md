# getset
Get information from various serialisation types and set them into templates.

[![Go Report Card](https://goreportcard.com/badge/github.com/codingconcepts/getset)](https://goreportcard.com/report/github.com/codingconcepts/getset)
[![Build Status](https://travis-ci.org/codingconcepts/getset.svg?branch=master)](https://travis-ci.org/codingconcepts/getset)

## Installation

``` bash
$ go get -u -d github.com/codingconcepts/getset
```

## Usage

``` bash
a := getset.New()

if err := a.Get(getset.TypeJSON, []byte(`{"a":{"b":"c"}}`), "a.b", "b"); err != nil {
	log.Fatalf("error getting value: %v", err)
}

resp, err := http.DefaultClient.Get("https://google.com")
if err != nil {
	log.Fatalf("error making request: %v", err)
}

if err := a.GetHeader(resp.Header, "Alt-Svc", "altsvc"); err != nil {
	log.Fatalf("error getting value: %v", err)
}

output, err := a.Set("Get value = {{.b}}\nGetHeader value = {{range .altsvc}}{{.}}{{end}}")
if err != nil {
	log.Fatalf("error setting placeholders: %v", err)
}

// Outputs
// Get value = c
// GetHeader value = quic=":443"; ma=2592000; v="46,44,43,39"
```