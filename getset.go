package getset

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/tidwall/gjson"
)

const (
	// TypeJSON will get a value from a given input using JSON dot notation.
	TypeJSON = "json"
)

type accumulator struct {
	m *smap
}

// New returns a pointer to an instance of an accumulator, the construct
// that holds the values you extract.
func New() *accumulator {
	return &accumulator{
		m: newSMAP(),
	}
}

// Get extracts information from input assuming it's of type t (JSON etc.).
// The path property will be used to access data within that input, so in
// the case of JSON, path will be dot notation to the key you're wanting a
// value for.  Use the name property to provide your own name to the value
// for setting later.  A name of "hello" will be usable as a placeholder
// later as "{{.hello}}".
func (a *accumulator) Get(t string, input []byte, path, name string) error {
	switch t {
	case TypeJSON:
		return a.getJSON(input, path, name)
	default:
		return nil
	}
}

// HeaderHeader extracts information from a given HTTP header.
func (a *accumulator) GetHeader(input map[string][]string, key, name string) error {
	value, ok := input[key]
	if !ok {
		return fmt.Errorf("no value in header at %q", key)
	}

	a.m.set(name, value)
	return nil
}

func (a *accumulator) getJSON(input []byte, path, name string) error {
	result := gjson.GetBytes(input, path)
	if !result.Exists() {
		return fmt.Errorf("no value in body at %q", path)
	}

	a.m.set(name, result.String())
	return nil
}

// Set takes an input, parses it as a Go template and applies values
// extracted from calls to Get.
func (a *accumulator) Set(input string) (string, error) {
	if a.m.empty() {
		return input, nil
	}

	buf := &bytes.Buffer{}
	t, err := template.New("").Parse(input)
	if err != nil {
		return "", err
	}
	if err = a.m.apply(buf, t); err != nil {
		return "", err
	}

	return buf.String(), nil
}
