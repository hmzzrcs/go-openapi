package openapi3

type baseMarshaller interface {
	MarshalJSON() ([]byte, error)
	MarshalYAML() (interface{}, error)
	marshaller
}

type marshaller interface {
	marshal() any
}
