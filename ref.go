package openapi3

import (
	"encoding/json"
	"fmt"
)

// Ref is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#reference-object
type Ref struct {
	Ref string `json:"$ref" yaml:"$ref"`
}

var (
	_ baseMarshaller = (*RefValue[marshaller])(nil)
)

type RefValue[T marshaller] struct {
	Ref   string
	Value T
}

// MarshalYAML returns the YAML encoding of CallbackRef.
func (x *RefValue[T]) MarshalYAML() (interface{}, error) {
	return x.marshal(), nil
}
func (x *RefValue[T]) marshal() any {
	if ref := x.Ref; ref != "" {
		return &Ref{Ref: ref}
	}
	return x.Value.marshal()
}

// MarshalJSON returns the JSON encoding of CallbackRef.
func (x *RefValue[T]) MarshalJSON() ([]byte, error) {

	return json.Marshal(x.marshal())
}

func (x *RefValue[T]) RefTo(name string) {
	x.Ref = fmt.Sprintf("#/components/%s/%s", refNames[T](), name)

}
func (x *RefValue[T]) Set(v T) {
	x.Value = v
	x.Ref = ""
}
