package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*Encoding)(nil)
)

// Encoding is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#encoding-object
type Encoding struct {
	extensions

	ContentType   string  `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       Headers `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string  `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool   `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool    `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

func (encoding *Encoding) MarshalYAML() (interface{}, error) {
	return encoding.marshal(), nil
}

func (encoding *Encoding) marshal() any {
	m := encoding.extensions.Export(5)
	if x := encoding.ContentType; x != "" {
		m["contentType"] = x
	}
	if x := encoding.Headers; len(x) != 0 {
		m["headers"] = x
	}
	if x := encoding.Style; x != "" {
		m["style"] = x
	}
	if x := encoding.Explode; x != nil {
		m["explode"] = x
	}
	if x := encoding.AllowReserved; x {
		m["allowReserved"] = x
	}
	return m
}

func NewEncoding() *Encoding {
	return &Encoding{}
}

func (encoding *Encoding) WithHeader(name string, header *Header) *Encoding {
	return encoding.WithHeaderRef(name, &HeaderRef{
		Value: header,
	})
}

func (encoding *Encoding) WithHeaderRef(name string, ref *HeaderRef) *Encoding {
	headers := encoding.Headers
	if headers == nil {
		headers = make(Headers)
		encoding.Headers = headers
	}
	headers[name] = ref
	return encoding
}

// MarshalJSON returns the JSON encoding of Encoding.
func (encoding *Encoding) MarshalJSON() ([]byte, error) {

	return json.Marshal(encoding.marshal())
}

// SerializationMethod returns a serialization method of request body.
// When serialization method is not defined the method returns the default serialization method.
func (encoding *Encoding) SerializationMethod() *SerializationMethod {
	sm := &SerializationMethod{Style: SerializationForm, Explode: true}
	if encoding != nil {
		if encoding.Style != "" {
			sm.Style = encoding.Style
		}
		if encoding.Explode != nil {
			sm.Explode = *encoding.Explode
		}
	}
	return sm
}
