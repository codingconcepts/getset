package getset

import (
	"io"
	"sync"
	"text/template"
)

type smap struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

func newSMAP() *smap {
	return &smap{
		m: make(map[string]interface{}),
	}
}

func (m *smap) set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[key] = value
}

func (m *smap) apply(w io.Writer, t *template.Template) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return t.Execute(w, m.m)
}

func (m *smap) empty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.m) == 0
}
