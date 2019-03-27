package getset

import (
	"fmt"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name     string
		t        string
		input    []byte
		path     string
		pathName string
		exp      interface{}
		expKey   bool
		expErr   bool
	}{
		{
			name:     "invalid type",
			t:        "invalid",
			input:    []byte(`{"a":"b"}`),
			path:     "a",
			pathName: "a",
		},
		{
			name:     "json not found top level",
			t:        TypeJSON,
			input:    []byte(`{"a":"b"}`),
			path:     "not_a_key",
			pathName: "not_a_key",
			expErr:   true,
		},
		{
			name:     "json not found",
			t:        TypeJSON,
			input:    []byte(`{"a":{"b":"c"}}`),
			path:     "not_a_key",
			pathName: "not_a_key",
			expErr:   true,
		},
		{
			name:     "json found top level",
			t:        TypeJSON,
			input:    []byte(`{"a":"b"}`),
			expKey:   true,
			path:     "a",
			pathName: "a",
			exp:      "b",
		},
		{
			name:     "json found",
			t:        TypeJSON,
			input:    []byte(`{"a":{"b":"c"}}`),
			expKey:   true,
			path:     "a.b",
			pathName: "ab",
			exp:      "c",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a := New()
			err := a.Get(c.t, c.input, c.path, c.pathName)
			if err != nil {
				if !c.expErr {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil && c.expErr {
				t.Fatal("expected error but didn't get one")
			}

			act, ok := a.m.m[c.pathName]
			if c.expKey && !ok {
				t.Fatalf("missing key: %s", c.path)
			}
			if !c.expKey && ok {
				t.Fatalf("unexpected value: %s", act)
			}

			equals(t, c.exp, act)
		})
	}
}

func TestGetHeader(t *testing.T) {
	cases := []struct {
		name    string
		input   map[string][]string
		key     string
		keyName string
		expKey  bool
		exp     []string
		expErr  bool
	}{
		{
			name:    "not found",
			input:   map[string][]string{"a": []string{"b"}},
			key:     "not_a_key",
			keyName: "not_a_key",
			expErr:  true,
		},
		{
			name:    "single value",
			input:   map[string][]string{"a": []string{"b"}},
			key:     "a",
			keyName: "a",
			expKey:  true,
			exp:     []string{"b"},
		},
		{
			name:    "multiple values",
			input:   map[string][]string{"a": []string{"b", "c"}},
			key:     "a",
			keyName: "a",
			expKey:  true,
			exp:     []string{"b", "c"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a := New()
			err := a.GetHeader(c.input, c.key, c.keyName)
			if err != nil {
				if !c.expErr {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil && c.expErr {
				t.Fatal("expected error but didn't get one")
			}

			act, ok := a.m.m[c.keyName]
			if c.expKey && !ok {
				t.Fatalf("missing key: %s", c.key)
			}
			if !c.expKey && ok {
				t.Fatalf("unexpected value: %s", act)
			}

			if !c.expKey {
				return
			}
			equals(t, c.exp, act)
		})
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		m      map[string]interface{}
		exp    string
		expErr bool
	}{
		{
			name:   "invalid template",
			input:  "invalid {{.template",
			m:      map[string]interface{}{"key": "value"},
			exp:    "input without a placeholder",
			expErr: true,
		},
		{
			name:  "without placeholder",
			input: "input without a placeholder",
			m:     map[string]interface{}{"key": "value"},
			exp:   "input without a placeholder",
		},
		{
			name:  "with placeholder",
			input: "input with a {{.placeholder}}",
			m:     map[string]interface{}{"placeholder": "value"},
			exp:   "input with a value",
		},
		{
			name:  "missing key",
			input: "input with a {{.placeholder}}",
			m:     map[string]interface{}{},
			exp:   "input with a {{.placeholder}}",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a := New()
			a.m.m = c.m

			act, err := a.Set(c.input)
			if err != nil {
				if !c.expErr {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil && c.expErr {
				t.Fatal("expected error but didn't get one")
			}

			equals(t, c.exp, act)
		})
	}
}

func Example() {
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
}
