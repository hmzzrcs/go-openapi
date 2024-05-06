package openapi3

var (
	_ baseMarshaller = (*Callback)(nil)
)

// Callback is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#callback-object
type Callback struct {
	extensions

	m map[string]*PathItem
}

func (callback *Callback) MarshalYAML() (interface{}, error) {
	return callback.marshal(), nil
}

func (callback *Callback) marshal() any {
	m := callback.extensions.Export(callback.Len())
	for k, v := range callback.Map() {
		m[k] = v
	}
	return m
}

// NewCallback builds a Callback object with path items in insertion order.
func NewCallback(opts ...NewCallbackOption) *Callback {
	Callback := NewCallbackWithCapacity(len(opts))
	for _, opt := range opts {
		opt(Callback)
	}
	return Callback
}

// NewCallbackOption describes options to NewCallback func
type NewCallbackOption func(*Callback)

// WithCallback adds Callback as an option to NewCallback
func WithCallback(cb string, pathItem *PathItem) NewCallbackOption {
	return func(callback *Callback) {
		if p := pathItem; p != nil && cb != "" {
			callback.Set(cb, p)
		}
	}
}
