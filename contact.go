package openapi3

import (
	"encoding/json"
)

var (
	_ baseMarshaller = (*Contact)(nil)
)

// Contact is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#contact-object
type Contact struct {
	extensions

	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

func (contact *Contact) MarshalYAML() (interface{}, error) {
	return contact.marshal(), nil
}

func (contact *Contact) marshal() any {
	m := contact.extensions.Export(3)
	if x := contact.Name; x != "" {
		m["name"] = x
	}
	if x := contact.URL; x != "" {
		m["url"] = x
	}
	if x := contact.Email; x != "" {
		m["email"] = x
	}
	return m
}

// MarshalJSON returns the JSON encoding of Contact.
func (contact *Contact) MarshalJSON() ([]byte, error) {

	return json.Marshal(contact.marshal())
}
