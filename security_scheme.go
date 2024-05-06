package openapi3

import (
	"encoding/json"
)

// SecurityScheme is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#security-scheme-object
type SecurityScheme struct {
	extensions

	Type             string      `json:"type,omitempty" yaml:"type,omitempty"`
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIdConnectUrl string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

func NewSecurityScheme() *SecurityScheme {
	return &SecurityScheme{}
}

func NewCSRFSecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		Type: "apiKey",
		In:   "header",
		Name: "X-XSRF-TOKEN",
	}
}

func NewOIDCSecurityScheme(oidcUrl string) *SecurityScheme {
	return &SecurityScheme{
		Type:             "openIdConnect",
		OpenIdConnectUrl: oidcUrl,
	}
}

func NewJWTSecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
	}
}

// MarshalJSON returns the JSON encoding of SecurityScheme.
func (ss *SecurityScheme) marshal() any {

	m := ss.extensions.Export(8)
	if x := ss.Type; x != "" {
		m["type"] = x
	}
	if x := ss.Description; x != "" {
		m["description"] = x
	}
	if x := ss.Name; x != "" {
		m["name"] = x
	}
	if x := ss.In; x != "" {
		m["in"] = x
	}
	if x := ss.Scheme; x != "" {
		m["scheme"] = x
	}
	if x := ss.BearerFormat; x != "" {
		m["bearerFormat"] = x
	}
	if x := ss.Flows; x != nil {
		m["flows"] = x
	}
	if x := ss.OpenIdConnectUrl; x != "" {
		m["openIdConnectUrl"] = x
	}
	return m
}

func (ss *SecurityScheme) WithType(value string) *SecurityScheme {
	ss.Type = value
	return ss
}

func (ss *SecurityScheme) WithDescription(value string) *SecurityScheme {
	ss.Description = value
	return ss
}

func (ss *SecurityScheme) WithName(value string) *SecurityScheme {
	ss.Name = value
	return ss
}

func (ss *SecurityScheme) WithIn(value string) *SecurityScheme {
	ss.In = value
	return ss
}

func (ss *SecurityScheme) WithScheme(value string) *SecurityScheme {
	ss.Scheme = value
	return ss
}

func (ss *SecurityScheme) WithBearerFormat(value string) *SecurityScheme {
	ss.BearerFormat = value
	return ss
}

var (
	_ baseMarshaller = (*OAuthFlows)(nil)
)

// OAuthFlows is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#oauth-flows-object
type OAuthFlows struct {
	extensions

	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

func (flows *OAuthFlows) MarshalYAML() (interface{}, error) {
	return flows.marshal(), nil
}

func (flows *OAuthFlows) marshal() any {
	m := flows.extensions.Export(4)
	if x := flows.Implicit; x != nil {
		m["implicit"] = x
	}
	if x := flows.Password; x != nil {
		m["password"] = x
	}
	if x := flows.ClientCredentials; x != nil {
		m["clientCredentials"] = x
	}
	if x := flows.AuthorizationCode; x != nil {
		m["authorizationCode"] = x
	}
	return m
}

// MarshalJSON returns the JSON encoding of OAuthFlows.
func (flows *OAuthFlows) MarshalJSON() ([]byte, error) {
	return json.Marshal(flows.marshal())
}

var (
	_ baseMarshaller = (*OAuthFlow)(nil)
)

// OAuthFlow is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#oauth-flow-object
type OAuthFlow struct {
	extensions

	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"` // required
}

func (flow *OAuthFlow) MarshalYAML() (interface{}, error) {
	return flow.marshal(), nil
}

func (flow *OAuthFlow) marshal() any {
	m := flow.extensions.Export(4)
	if x := flow.AuthorizationURL; x != "" {
		m["authorizationUrl"] = x
	}
	if x := flow.TokenURL; x != "" {
		m["tokenUrl"] = x
	}
	if x := flow.RefreshURL; x != "" {
		m["refreshUrl"] = x
	}
	m["scopes"] = flow.Scopes

	return m
}

// MarshalJSON returns the JSON encoding of OAuthFlow.
func (flow *OAuthFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(flow.marshal())
}

func (ss *SecurityScheme) MarshalJSON() ([]byte, error) { return json.Marshal(ss.marshal()) }
