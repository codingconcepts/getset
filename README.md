# getset
Get information from various serialisation types and set them into templates.

## Installation

``` bash
$ go get -u -d github.com/codingconcepts/getset
```

## Usage

``` bash
a := New()

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