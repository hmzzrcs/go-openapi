package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*RequestBody)(nil)
)

// RequestBody is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#request-body-object
type RequestBody struct {
	extensions

	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Content     Content `json:"content" yaml:"content"`
}

func (requestBody *RequestBody) MarshalYAML() (interface{}, error) {
	return requestBody.marshal(), nil
}

func NewRequestBody() *RequestBody {
	return &RequestBody{}
}

func (requestBody *RequestBody) WithDescription(value string) *RequestBody {
	requestBody.Description = value
	return requestBody
}

func (requestBody *RequestBody) WithRequired(value bool) *RequestBody {
	requestBody.Required = value
	return requestBody
}

func (requestBody *RequestBody) WithContent(content Content) *RequestBody {
	requestBody.Content = content
	return requestBody
}

func (requestBody *RequestBody) WithSchemaRef(value *SchemaRef, consumes []string) *RequestBody {
	requestBody.Content = NewContentWithSchemaRef(value, consumes)
	return requestBody
}

func (requestBody *RequestBody) WithSchema(value *Schema, consumes []string) *RequestBody {
	requestBody.Content = NewContentWithSchema(value, consumes)
	return requestBody
}

func (requestBody *RequestBody) WithJSONSchemaRef(value *SchemaRef) *RequestBody {
	requestBody.Content = NewContentWithJSONSchemaRef(value)
	return requestBody
}

func (requestBody *RequestBody) WithJSONSchema(value *Schema) *RequestBody {
	requestBody.Content = NewContentWithJSONSchema(value)
	return requestBody
}

func (requestBody *RequestBody) WithFormDataSchemaRef(value *SchemaRef) *RequestBody {
	requestBody.Content = NewContentWithFormDataSchemaRef(value)
	return requestBody
}

func (requestBody *RequestBody) WithFormDataSchema(value *Schema) *RequestBody {
	requestBody.Content = NewContentWithFormDataSchema(value)
	return requestBody
}

func (requestBody *RequestBody) GetMediaType(mediaType string) *MediaType {
	m := requestBody.Content
	if m == nil {
		return nil
	}
	return m[mediaType]
}

// MarshalJSON returns the JSON encoding of RequestBody.
func (requestBody *RequestBody) marshal() any {

	m := requestBody.extensions.Export(3)
	if x := requestBody.Description; x != "" {
		m["description"] = requestBody.Description
	}
	if x := requestBody.Required; x {
		m["required"] = x
	}
	if x := requestBody.Content; true {
		m["content"] = x
	}

	return m
}

func (requestBody *RequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(requestBody.marshal())
}
