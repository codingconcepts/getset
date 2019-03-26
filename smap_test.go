package getset

import (
	"bytes"
	"html/template"
	"testing"
)

func TestSMAP_Set(t *testing.T) {
	m := newSMAP()
	m.set("key", "value")

	equals(t, "value", m.m["key"])
}

func TestSMAP_Apply(t *testing.T) {
	m := newSMAP()
	m.set("key", "value")

	input := "this has a {{.key}}"
	exp := "this has a value"

	buf := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(input)
	if err != nil {
		t.Fatalf("error parting template: %v", err)
	}
	if err = m.apply(buf, tmpl); err != nil {
		t.Fatalf("error applying template: %v", err)
	}

	equals(t, exp, buf.String())
}

func TestSMAP_Empty(t *testing.T) {
	m := newSMAP()
	equals(t, true, m.empty())

	m.set("key", "value")
	equals(t, false, m.empty())
}
