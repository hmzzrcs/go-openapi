package openapi3

import "encoding/json"

var (
	_ baseMarshaller = (*Info)(nil)
)

// Info is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#info-object
type Info struct {
	extensions

	Title          string   `json:"title" yaml:"title"` // Required
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"` // Required
}

func (info *Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(info.marshal())
}

func (info *Info) MarshalYAML() (interface{}, error) {
	return info.marshal(), nil

}

// MarshalJSON returns the JSON encoding of Info.
func (info *Info) marshal() any {
	m := info.extensions.Export(6)
	m["title"] = info.Title
	if x := info.Description; x != "" {
		m["description"] = x
	}
	if x := info.TermsOfService; x != "" {
		m["termsOfService"] = x
	}
	if x := info.Contact; x != nil {
		m["contact"] = x
	}
	if x := info.License; x != nil {
		m["license"] = x
	}
	m["version"] = info.Version
	return m
}
