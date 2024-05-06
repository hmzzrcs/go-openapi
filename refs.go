package openapi3

import "encoding/json"

// CallbackRef represents either a Callback or a $ref to a Callback.
// When serializing and both fields are set, Ref is preferred over Value.
type CallbackRef = RefValue[*Callback]

// ExampleRef represents either an Example or a $ref to an Example.
// When serializing and both fields are set, Ref is preferred over Value.
type ExampleRef = RefValue[*Example]

// HeaderRef represents either a Header or a $ref to a Header.
// When serializing and both fields are set, Ref is preferred over Value.
type HeaderRef = RefValue[*Header]

// LinkRef represents either a Link or a $ref to a Link.
// When serializing and both fields are set, Ref is preferred over Value.
type LinkRef = RefValue[*Link]

// ParameterRef represents either a Parameter or a $ref to a Parameter.
// When serializing and both fields are set, Ref is preferred over Value.
type ParameterRef = RefValue[*Parameter]

// RequestBodyRef represents either a RequestBody or a $ref to a RequestBody.
// When serializing and both fields are set, Ref is preferred over Value.
type RequestBodyRef = RefValue[*RequestBody]

// ResponseRef represents either a Response or a $ref to a Response.
// When serializing and both fields are set, Ref is preferred over Value.
type ResponseRef = RefValue[*Response]

var (
	_ baseMarshaller = (*SchemaRef)(nil)
)

// SchemaRef represents either a Schema or a $ref to a Schema.
// When serializing and both fields are set, Ref is preferred over Value.
type SchemaRef RefValue[*Schema]

// 由于存在循环定义, 不能用别名,需要重新实现 marshaller
func (s *SchemaRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.marshal())
}

func (s *SchemaRef) MarshalYAML() (interface{}, error) {
	return s.marshal(), nil
}

func (s *SchemaRef) marshal() any {
	if s.Ref != "" {
		return &Ref{
			Ref: s.Ref,
		}
	}
	return s.Value.marshal()
}

// SecuritySchemeRef represents either a SecurityScheme or a $ref to a SecurityScheme.
// When serializing and both fields are set, Ref is preferred over Value.
type SecuritySchemeRef = RefValue[*SecurityScheme]

type Refs[T marshaller] map[string]*RefValue[T]

func (refs Refs[T]) AddValue(name string, value T) {
	r, has := refs[name]
	if !has {
		r = new(RefValue[T])
		refs[name] = r
	}
	r.Set(value)
}
func (refs Refs[T]) AddRef(name string, ref string) {
	r, has := refs[name]
	if !has {
		r = new(RefValue[T])
		refs[name] = r
	}
	r.RefTo(ref)
}
