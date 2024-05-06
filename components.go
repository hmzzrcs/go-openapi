package openapi3

import (
	"encoding/json"
)

type (
	Callbacks       Refs[*Callback]
	Examples        Refs[*Example]
	Headers         Refs[*Header]
	Links           Refs[*Link]
	ParametersMap   Refs[*Parameter]
	RequestBodies   Refs[*RequestBody]
	ResponseBodies  Refs[*Response]
	Schemas         map[string]*SchemaRef
	SecuritySchemes Refs[*SecurityScheme]
)

var (
	_ baseMarshaller = (*Components)(nil)
)

// Components are specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#components-object
type Components struct {
	extensions

	Schemas         Schemas         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Parameters      ParametersMap   `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Headers         Headers         `json:"headers,omitempty" yaml:"headers,omitempty"`
	RequestBodies   RequestBodies   `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Responses       ResponseBodies  `json:"responses,omitempty" yaml:"responses,omitempty"`
	SecuritySchemes SecuritySchemes `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Examples        Examples        `json:"examples,omitempty" yaml:"examples,omitempty"`
	Links           Links           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       Callbacks       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

func (components *Components) MarshalYAML() (interface{}, error) {
	return components.marshal(), nil
}

func (components *Components) marshal() any {
	m := components.extensions.Export(9)
	if x := components.Schemas; len(x) != 0 {
		m["schemas"] = x
	}
	if x := components.Parameters; len(x) != 0 {
		m["parameters"] = x
	}
	if x := components.Headers; len(x) != 0 {
		m["headers"] = x
	}
	if x := components.RequestBodies; len(x) != 0 {
		m["requestBodies"] = x
	}
	if x := components.Responses; len(x) != 0 {
		m["responses"] = x
	}
	if x := components.SecuritySchemes; len(x) != 0 {
		m["securitySchemes"] = x
	}
	if x := components.Examples; len(x) != 0 {
		m["examples"] = x
	}
	if x := components.Links; len(x) != 0 {
		m["links"] = x
	}
	if x := components.Callbacks; len(x) != 0 {
		m["callbacks"] = x
	}
	return m
}

func NewComponents() *Components {
	return &Components{}
}

// MarshalJSON returns the JSON encoding of Components.
func (components *Components) MarshalJSON() ([]byte, error) {

	return json.Marshal(components.marshal())
}

func (components *Components) AddHeaders() {

}
