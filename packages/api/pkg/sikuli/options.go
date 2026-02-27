package sikuli

import (
	"strconv"
	"sync"
)

type Options struct {
	mu      sync.RWMutex
	entries map[string]string
}

func NewOptions() *Options {
	return &Options{
		entries: map[string]string{},
	}
}

func NewOptionsFromMap(entries map[string]string) *Options {
	o := NewOptions()
	for k, v := range entries {
		o.entries[k] = v
	}
	return o
}

func (o *Options) Has(key string) bool {
	o.mu.RLock()
	defer o.mu.RUnlock()
	_, ok := o.entries[key]
	return ok
}

func (o *Options) GetString(key, def string) string {
	o.mu.RLock()
	defer o.mu.RUnlock()
	v, ok := o.entries[key]
	if !ok {
		return def
	}
	return v
}

func (o *Options) SetString(key, value string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.entries[key] = value
}

func (o *Options) GetInt(key string, def int) int {
	raw := o.GetString(key, "")
	if raw == "" {
		return def
	}
	v, err := strconv.Atoi(raw)
	if err != nil {
		return def
	}
	return v
}

func (o *Options) SetInt(key string, value int) {
	o.SetString(key, strconv.Itoa(value))
}

func (o *Options) GetFloat64(key string, def float64) float64 {
	raw := o.GetString(key, "")
	if raw == "" {
		return def
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return def
	}
	return v
}

func (o *Options) SetFloat64(key string, value float64) {
	o.SetString(key, strconv.FormatFloat(value, 'f', -1, 64))
}

func (o *Options) GetBool(key string, def bool) bool {
	raw := o.GetString(key, "")
	if raw == "" {
		return def
	}
	v, err := strconv.ParseBool(raw)
	if err != nil {
		return def
	}
	return v
}

func (o *Options) SetBool(key string, value bool) {
	o.SetString(key, strconv.FormatBool(value))
}

func (o *Options) Delete(key string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	delete(o.entries, key)
}

func (o *Options) Entries() map[string]string {
	o.mu.RLock()
	defer o.mu.RUnlock()
	out := make(map[string]string, len(o.entries))
	for k, v := range o.entries {
		out[k] = v
	}
	return out
}

func (o *Options) Merge(other *Options) {
	if other == nil {
		return
	}
	for k, v := range other.Entries() {
		o.SetString(k, v)
	}
}

func (o *Options) Clone() *Options {
	return NewOptionsFromMap(o.Entries())
}
