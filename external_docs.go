package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*ExternalDocs)(nil)
)

// ExternalDocs is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#external-documentation-object
type ExternalDocs struct {
	extensions

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}

func (e *ExternalDocs) MarshalYAML() (interface{}, error) {
	return e.marshal(), nil
}

func (e *ExternalDocs) marshal() any {
	m := e.extensions.Export(2)
	if x := e.Description; x != "" {
		m["description"] = x
	}
	if x := e.URL; x != "" {
		m["url"] = x
	}
	return m
}

// MarshalJSON returns the JSON encoding of ExternalDocs.
func (e *ExternalDocs) MarshalJSON() ([]byte, error) {

	return json.Marshal(e.marshal())
}
