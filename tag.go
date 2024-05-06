package openapi3

import (
	"encoding/json"
)

// Tags is specified by OpenAPI/Swagger 3.0 standard.
type Tags []*Tag

func (tags Tags) Get(name string) *Tag {
	for _, tag := range tags {
		if tag.Name == name {
			return tag
		}
	}
	return nil
}

var (
	_ baseMarshaller = (*Tag)(nil)
)

// Tag is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#tag-object
type Tag struct {
	extensions

	Name         string        `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

func (t *Tag) MarshalYAML() (interface{}, error) {
	return t.marshal(), nil
}

// MarshalJSON returns the JSON encoding of Tag.
func (t *Tag) marshal() any {

	m := t.extensions.Export(3)
	if x := t.Name; x != "" {
		m["name"] = x
	}
	if x := t.Description; x != "" {
		m["description"] = x
	}
	if x := t.ExternalDocs; x != nil {
		m["externalDocs"] = x
	}

	return m
}

func (t *Tag) MarshalJSON() ([]byte, error) { return json.Marshal(t.marshal()) }
