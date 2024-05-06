package openapi3

var (
	_ baseMarshaller = (*Header)(nil)
)

// Header is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#header-object
type Header struct {
	Parameter
}
