package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*Example)(nil)
)

// Example is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#example-object
type Example struct {
	extensions

	Summary       string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{} `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

func (example *Example) MarshalYAML() (interface{}, error) {
	return example.marshal(), nil
}

func (example *Example) marshal() any {

	m := example.extensions.Export(4)
	if x := example.Summary; x != "" {
		m["summary"] = x
	}
	if x := example.Description; x != "" {
		m["description"] = x
	}
	if x := example.Value; x != nil {
		m["value"] = x
	}
	if x := example.ExternalValue; x != "" {
		m["externalValue"] = x
	}
	return m
}

func NewExample(value interface{}) *Example {
	return &Example{Value: value}
}

// MarshalJSON returns the JSON encoding of Example.
func (example *Example) MarshalJSON() ([]byte, error) {

	return json.Marshal(example.marshal())
}
