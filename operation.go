package openapi3

import (
	"encoding/json"
	"strconv"
)

var (
	_ baseMarshaller = (*Operation)(nil)
)

// Operation represents "operation" specified by" OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#operation-object
type Operation struct {
	extensions

	// Optional tags for documentation.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Optional short summary.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Optional description. Should use CommonMark syntax.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Optional operation ID.
	OperationID string `json:"operationId,omitempty" yaml:"operationId,omitempty"`

	// Optional parameters.
	Parameters Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Optional body parameter.
	RequestBody *RequestBodyRef `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`

	// Responses.
	Responses *Responses `json:"responses" yaml:"responses"` // Required

	// Optional callbacks
	Callbacks Callbacks `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`

	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Optional security requirements that override top-level security.
	Security *SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`

	// Optional servers that override top-level servers.
	Servers *Servers `json:"servers,omitempty" yaml:"servers,omitempty"`

	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

func (operation *Operation) MarshalYAML() (interface{}, error) {
	return operation.marshal(), nil
}

func NewOperation() *Operation {
	return &Operation{}
}

// MarshalJSON returns the JSON encoding of Operation.
func (operation *Operation) marshal() any {

	m := operation.extensions.Export(12)
	if x := operation.Tags; len(x) != 0 {
		m["tags"] = x
	}
	if x := operation.Summary; x != "" {
		m["summary"] = x
	}
	if x := operation.Description; x != "" {
		m["description"] = x
	}
	if x := operation.OperationID; x != "" {
		m["operationId"] = x
	}
	if x := operation.Parameters; len(x) != 0 {
		m["parameters"] = x
	}
	if x := operation.RequestBody; x != nil {
		m["requestBody"] = x
	}
	m["responses"] = operation.Responses
	if x := operation.Callbacks; len(x) != 0 {
		m["callbacks"] = x
	}
	if x := operation.Deprecated; x {
		m["deprecated"] = x
	}
	if x := operation.Security; x != nil {
		m["security"] = x
	}
	if x := operation.Servers; x != nil {
		m["servers"] = x
	}
	if x := operation.ExternalDocs; x != nil {
		m["externalDocs"] = x
	}

	return m
}

func (operation *Operation) MarshalJSON() ([]byte, error) { return json.Marshal(operation.marshal()) }

func (operation *Operation) AddParameter(p *Parameter) {
	operation.Parameters = append(operation.Parameters, &ParameterRef{Value: p})
}

func (operation *Operation) AddResponse(status int, response *Response) {
	code := "default"
	if 0 < status && status < 1000 {
		code = strconv.FormatInt(int64(status), 10)
	}
	if operation.Responses == nil {
		operation.Responses = NewResponses()
	}
	operation.Responses.Set(code, &ResponseRef{Value: response})
}
