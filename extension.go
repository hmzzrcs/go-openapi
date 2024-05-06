package openapi3

type IExtensions interface {
	AddExtensions(key string, value interface{})
	RemoveExtensions(key string)
	Export(gr int) map[string]interface{}
}

type extensions struct {
	data map[string]interface{}
}

func (e *extensions) AddExtensions(key string, value interface{}) {
	if e.data == nil {
		e.data = make(map[string]interface{})
	}
	e.data[key] = value
}

func (e *extensions) RemoveExtensions(key string) {
	if e.data != nil {
		delete(e.data, key)
	}
}

func (e *extensions) Export(gr int) map[string]interface{} {
	m := make(map[string]interface{}, len(e.data)+gr)
	for k, v := range e.data {
		m[k] = v
	}
	return m
}
