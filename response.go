package openapi3

import (
	"encoding/json"
	"strconv"
)

var (
	_ baseMarshaller = (*Responses)(nil)
)

// Responses is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#responses-object
type Responses struct {
	extensions

	m map[string]*ResponseRef
}

// NewResponses builds a responses object with response objects in insertion order.
// Given no arguments, NewResponses returns a valid responses object containing a default match-all reponse.
func NewResponses(opts ...NewResponsesOption) *Responses {
	if len(opts) == 0 {
		return NewResponses(WithName("default", NewResponse().WithDescription("")))
	}
	responses := NewResponsesWithCapacity(len(opts))
	for _, opt := range opts {
		opt(responses)
	}
	return responses
}

// NewResponsesOption describes options to NewResponses func
type NewResponsesOption func(*Responses)

// WithStatus adds a status code keyed ResponseRef
func WithStatus(status int, responseRef *ResponseRef) NewResponsesOption {
	return func(responses *Responses) {
		if r := responseRef; r != nil {
			code := strconv.FormatInt(int64(status), 10)
			responses.Set(code, r)
		}
	}
}

// WithName adds a name-keyed Response
func WithName(name string, response *Response) NewResponsesOption {
	return func(responses *Responses) {
		if r := response; r != nil && name != "" {
			responses.Set(name, &ResponseRef{Value: r})
		}
	}
}

// Default returns the default response
func (responses *Responses) Default() *ResponseRef {
	return responses.Value("default")
}

// Status returns a ResponseRef for the given status
// If an exact match isn't initially found a patterned field is checked using
// the first digit to determine the range (eg: 201 to 2XX)
// See https://spec.openapis.org/oas/v3.0.3#patterned-fields-0
func (responses *Responses) Status(status int) *ResponseRef {
	st := strconv.FormatInt(int64(status), 10)
	if rref := responses.Value(st); rref != nil {
		return rref
	}
	if 99 < status && status < 600 {
		st = string(st[0]) + "XX"
		switch st {
		case "1XX", "2XX", "3XX", "4XX", "5XX":
			return responses.Value(st)
		}
	}
	return nil
}

// Response is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#response-object
type Response struct {
	extensions

	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     Headers `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
	Links       Links   `json:"links,omitempty" yaml:"links,omitempty"`
}

func NewResponse() *Response {
	return &Response{}
}

func (response *Response) WithDescription(value string) *Response {
	response.Description = &value
	return response
}

func (response *Response) WithContent(content Content) *Response {
	response.Content = content
	return response
}

func (response *Response) WithJSONSchema(schema *Schema) *Response {
	response.Content = NewContentWithJSONSchema(schema)
	return response
}

func (response *Response) WithJSONSchemaRef(schema *SchemaRef) *Response {
	response.Content = NewContentWithJSONSchemaRef(schema)
	return response
}

// MarshalJSON returns the JSON encoding of Response.
func (response *Response) marshal() any {

	m := response.extensions.Export(4)
	if x := response.Description; x != nil {
		m["description"] = x
	}
	if x := response.Headers; len(x) != 0 {
		m["headers"] = x
	}
	if x := response.Content; len(x) != 0 {
		m["content"] = x
	}
	if x := response.Links; len(x) != 0 {
		m["links"] = x
	}

	return m
}

func (response *Response) MarshalJSON() ([]byte, error) { return json.Marshal(response.marshal()) }
