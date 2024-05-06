package openapi3

import (
	"encoding/json"
)

// Parameters is specified by OpenAPI/Swagger 3.0 standard.
type Parameters []*ParameterRef

func NewParameters() Parameters {
	return make(Parameters, 0, 4)
}

func (parameters Parameters) GetByInAndName(in string, name string) *Parameter {
	for _, item := range parameters {
		if v := item.Value; v != nil {
			if v.Name == name && v.In == in {
				return v
			}
		}
	}
	return nil
}

var (
	_ baseMarshaller = (*Parameter)(nil)
)

// Parameter is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#parameter-object
type Parameter struct {
	extensions

	Name            string      `json:"name,omitempty" yaml:"name,omitempty"`
	In              string      `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string      `json:"description,omitempty" yaml:"description,omitempty"`
	Style           string      `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool       `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	AllowReserved   bool        `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Deprecated      bool        `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Required        bool        `json:"required,omitempty" yaml:"required,omitempty"`
	Schema          *SchemaRef  `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content         Content     `json:"content,omitempty" yaml:"content,omitempty"`
}

func (parameter *Parameter) MarshalYAML() (interface{}, error) {
	return parameter.marshal(), nil
}

const (
	ParameterInPath   = "path"
	ParameterInQuery  = "query"
	ParameterInHeader = "header"
	ParameterInCookie = "cookie"
)

func NewPathParameter(name string) *Parameter {
	return &Parameter{
		Name:     name,
		In:       ParameterInPath,
		Required: true,
	}
}

func NewQueryParameter(name string) *Parameter {
	return &Parameter{
		Name: name,
		In:   ParameterInQuery,
	}
}

func NewHeaderParameter(name string) *Parameter {
	return &Parameter{
		Name: name,
		In:   ParameterInHeader,
	}
}

func NewCookieParameter(name string) *Parameter {
	return &Parameter{
		Name: name,
		In:   ParameterInCookie,
	}
}

func (parameter *Parameter) WithDescription(value string) *Parameter {
	parameter.Description = value
	return parameter
}

func (parameter *Parameter) WithRequired(value bool) *Parameter {
	parameter.Required = value
	return parameter
}

func (parameter *Parameter) WithSchema(value *Schema) *Parameter {
	if value == nil {
		parameter.Schema = nil
	} else {
		parameter.Schema = &SchemaRef{
			Value: value,
		}
	}
	return parameter
}

// MarshalJSON returns the JSON encoding of Parameter.
func (parameter *Parameter) MarshalJSON() ([]byte, error) {

	return json.Marshal(parameter.marshal())
}
func (parameter *Parameter) marshal() any {
	m := parameter.extensions.Export(13)

	if x := parameter.Name; x != "" {
		m["name"] = x
	}
	if x := parameter.In; x != "" {
		m["in"] = x
	}
	if x := parameter.Description; x != "" {
		m["description"] = x
	}
	if x := parameter.Style; x != "" {
		m["style"] = x
	}
	if x := parameter.Explode; x != nil {
		m["explode"] = x
	}
	if x := parameter.AllowEmptyValue; x {
		m["allowEmptyValue"] = x
	}
	if x := parameter.AllowReserved; x {
		m["allowReserved"] = x
	}
	if x := parameter.Deprecated; x {
		m["deprecated"] = x
	}
	if x := parameter.Required; x {
		m["required"] = x
	}
	if x := parameter.Schema; x != nil {
		m["schema"] = x
	}
	if x := parameter.Example; x != nil {
		m["example"] = x
	}
	if x := parameter.Examples; len(x) != 0 {
		m["examples"] = x
	}
	if x := parameter.Content; len(x) != 0 {
		m["content"] = x
	}
	return m
}
