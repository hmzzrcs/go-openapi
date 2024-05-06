package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*MediaType)(nil)
)

// MediaType is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#media-type-object
type MediaType struct {
	extensions

	Schema   *SchemaRef           `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  interface{}          `json:"example,omitempty" yaml:"example,omitempty"`
	Examples Examples             `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

func (mediaType *MediaType) MarshalYAML() (interface{}, error) {
	return mediaType.marshal(), nil
}

func NewMediaType() *MediaType {
	return &MediaType{}
}

func (mediaType *MediaType) WithSchema(schema *Schema) *MediaType {
	if schema == nil {
		mediaType.Schema = nil
	} else {
		mediaType.Schema = &SchemaRef{Value: schema}
	}
	return mediaType
}

func (mediaType *MediaType) WithSchemaRef(schema *SchemaRef) *MediaType {
	mediaType.Schema = schema
	return mediaType
}

func (mediaType *MediaType) WithExample(name string, value interface{}) *MediaType {
	example := mediaType.Examples
	if example == nil {
		example = make(map[string]*ExampleRef)
		mediaType.Examples = example
	}
	example[name] = &ExampleRef{
		Value: NewExample(value),
	}
	return mediaType
}

func (mediaType *MediaType) WithEncoding(name string, enc *Encoding) *MediaType {
	encoding := mediaType.Encoding
	if encoding == nil {
		encoding = make(map[string]*Encoding)
		mediaType.Encoding = encoding
	}
	encoding[name] = enc
	return mediaType
}

// MarshalJSON returns the JSON encoding of MediaType.
func (mediaType *MediaType) marshal() any {

	m := mediaType.extensions.Export(4)
	if x := mediaType.Schema; x != nil {
		m["schema"] = x
	}
	if x := mediaType.Example; x != nil {
		m["example"] = x
	}
	if x := mediaType.Examples; len(x) != 0 {
		m["examples"] = x
	}
	if x := mediaType.Encoding; len(x) != 0 {
		m["encoding"] = x
	}

	return m
}

func (mediaType *MediaType) MarshalJSON() ([]byte, error) { return json.Marshal(mediaType.marshal()) }
