package getset

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/tidwall/gjson"
)

const (
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

func (a *accumulator) Get(t string, input []byte, path, name string) error {
	switch t {
	case TypeJSON:
		return a.getJSON(input, path, name)
	default:
		return nil
	}
}

func (a *accumulator) getJSON(input []byte, path, name string) error {
	result := gjson.GetBytes(input, path)
	if !result.Exists() {
		return fmt.Errorf("no valid in body at %q", path)
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
