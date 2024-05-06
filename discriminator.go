package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*Discriminator)(nil)
)

// Discriminator is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#discriminator-object
type Discriminator struct {
	extensions

	PropertyName string            `json:"propertyName" yaml:"propertyName"` // required
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

func (discriminator *Discriminator) MarshalYAML() (interface{}, error) {
	return discriminator.marshal(), nil
}

func (discriminator *Discriminator) marshal() any {
	m := discriminator.extensions.Export(2)
	m["propertyName"] = discriminator.PropertyName
	if x := discriminator.Mapping; len(x) != 0 {
		m["mapping"] = x
	}
	return m
}

// MarshalJSON returns the JSON encoding of Discriminator.
func (discriminator *Discriminator) MarshalJSON() ([]byte, error) {

	return json.Marshal(discriminator.marshal())
}
