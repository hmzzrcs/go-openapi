package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*XML)(nil)
)

// XML is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#xml-object
type XML struct {
	extensions

	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

func (xml *XML) MarshalYAML() (interface{}, error) {
	return xml.marshal(), nil
}

// MarshalJSON returns the JSON encoding of XML.
func (xml *XML) marshal() any {

	m := xml.extensions.Export(5)
	if x := xml.Name; x != "" {
		m["name"] = x
	}
	if x := xml.Namespace; x != "" {
		m["namespace"] = x
	}
	if x := xml.Prefix; x != "" {
		m["prefix"] = x
	}
	if x := xml.Attribute; x {
		m["attribute"] = x
	}
	if x := xml.Wrapped; x {
		m["wrapped"] = x
	}

	return m
}

func (xml *XML) MarshalJSON() ([]byte, error) { return json.Marshal(xml.marshal()) }
