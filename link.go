package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*Link)(nil)
)

// The Link is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#link-object
type Link struct {
	extensions

	OperationRef string                 `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Server       *Server                `json:"server,omitempty" yaml:"server,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
}

func (link *Link) MarshalYAML() (interface{}, error) {
	return link.marshal(), nil
}

// MarshalJSON returns the JSON encoding of Link.
func (link *Link) MarshalJSON() ([]byte, error) {
	return json.Marshal(link.marshal())
}
func (link *Link) marshal() any {

	m := link.extensions.Export(6)

	if x := link.OperationRef; x != "" {
		m["operationRef"] = x
	}
	if x := link.OperationID; x != "" {
		m["operationId"] = x
	}
	if x := link.Description; x != "" {
		m["description"] = x
	}
	if x := link.Parameters; len(x) != 0 {
		m["parameters"] = x
	}
	if x := link.Server; x != nil {
		m["server"] = x
	}
	if x := link.RequestBody; x != nil {
		m["requestBody"] = x
	}

	return m
}
