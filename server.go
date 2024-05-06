package openapi3

import (
	"encoding/json"
)

// Servers is specified by OpenAPI/Swagger standard version 3.
type Servers []*Server

var (
	_ baseMarshaller = (*Server)(nil)
)

// Server is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#server-object
type Server struct {
	extensions

	URL         string                     `json:"url" yaml:"url"` // Required
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

func (server *Server) MarshalYAML() (interface{}, error) {

	return server.marshal(), nil
}

// MarshalJSON returns the JSON encoding of Server.
func (server *Server) marshal() any {

	m := server.extensions.Export(3)
	m["url"] = server.URL
	if x := server.Description; x != "" {
		m["description"] = x
	}
	if x := server.Variables; len(x) != 0 {
		m["variables"] = x
	}
	return m
}

var (
	_ baseMarshaller = (*ServerVariable)(nil)
)

// ServerVariable is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#server-variable-object
type ServerVariable struct {
	extensions

	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

func (serverVariable *ServerVariable) MarshalYAML() (interface{}, error) {
	return serverVariable.marshal(), nil
}

// MarshalJSON returns the JSON encoding of ServerVariable.
func (serverVariable *ServerVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(serverVariable.marshal())
}
func (serverVariable *ServerVariable) marshal() any {
	m := serverVariable.extensions.Export(4)
	if x := serverVariable.Enum; len(x) != 0 {
		m["enum"] = x
	}
	if x := serverVariable.Default; x != "" {
		m["default"] = x
	}
	if x := serverVariable.Description; x != "" {
		m["description"] = x
	}

	return m
}

func (server *Server) MarshalJSON() ([]byte, error) { return json.Marshal(server.marshal()) }
