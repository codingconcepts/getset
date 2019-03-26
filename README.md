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

if err := a.Get(TypeJSON, []byte(`{"a":{"b":"c"}}`), "a.b", "b"); err != nil {
	log.Fatalf("error getting value: %v", err)
}

output, err := a.Set("the value of b is {{.b}}")
if err != nil {
	log.Fatalf("error setting placeholder: %v", err)
}

fmt.Println(output)
// Output: the value of b is c
```