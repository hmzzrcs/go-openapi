package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*License)(nil)
)

// The License is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#license-object
type License struct {
	extensions

	Name string `json:"name" yaml:"name"` // Required
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

func (license *License) MarshalYAML() (interface{}, error) {
	return license.marshal(), nil
}

func (license *License) marshal() any {
	m := license.extensions.Export(2)
	m["name"] = license.Name
	if x := license.URL; x != "" {
		m["url"] = x
	}
	return m
}

// MarshalJSON returns the JSON encoding of License.
func (license *License) MarshalJSON() ([]byte, error) {

	return json.Marshal(license.marshal())
}
