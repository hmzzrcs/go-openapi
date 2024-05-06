package openapi3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sync"
)

const (
	TypeArray   = "array"
	TypeBoolean = "boolean"
	TypeInteger = "integer"
	TypeNumber  = "number"
	TypeObject  = "object"
	TypeString  = "string"
	TypeNull    = "null"

	// constants for integer formats
	formatMinInt32 = float64(math.MinInt32)
	formatMaxInt32 = float64(math.MaxInt32)
	formatMinInt64 = float64(math.MinInt64)
	formatMaxInt64 = float64(math.MaxInt64)
)

var (
	// SchemaErrorDetailsDisabled disables printing of details about schema errors.
	SchemaErrorDetailsDisabled = false

	errSchema = errors.New("input does not match the schema")

	// ErrOneOfConflict is the SchemaError Origin when data matches more than one oneOf schema
	ErrOneOfConflict = errors.New("input matches more than one oneOf schemas")

	// ErrSchemaInputNaN may be returned when validating a number
	ErrSchemaInputNaN = errors.New("floating point NaN is not allowed")
	// ErrSchemaInputInf may be returned when validating a number
	ErrSchemaInputInf = errors.New("floating point Inf is not allowed")

	compiledPatterns sync.Map
)

type SchemaRefs []*SchemaRef

// Schema is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#schema-object
type Schema struct {
	extensions

	OneOf        SchemaRefs    `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf        SchemaRefs    `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	AllOf        SchemaRefs    `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Not          *SchemaRef    `json:"not,omitempty" yaml:"not,omitempty"`
	Type         *Types        `json:"type,omitempty" yaml:"type,omitempty"`
	Title        string        `json:"title,omitempty" yaml:"title,omitempty"`
	Format       string        `json:"format,omitempty" yaml:"format,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	Enum         []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default      interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Example      interface{}   `json:"example,omitempty" yaml:"example,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Array-related, here for struct compactness
	UniqueItems bool `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	// Number-related, here for struct compactness
	ExclusiveMin bool `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	ExclusiveMax bool `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	// Properties
	Nullable        bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	ReadOnly        bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly       bool `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Deprecated      bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	XML             *XML `json:"xml,omitempty" yaml:"xml,omitempty"`

	// Number
	Min        *float64 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Max        *float64 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`

	// String
	MinLength uint64  `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength *uint64 `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Pattern   string  `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// Array
	MinItems uint64     `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxItems *uint64    `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	Items    *SchemaRef `json:"items,omitempty" yaml:"items,omitempty"`

	// Object
	Required             []string             `json:"required,omitempty" yaml:"required,omitempty"`
	Properties           Schemas              `json:"properties,omitempty" yaml:"properties,omitempty"`
	MinProps             uint64               `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	MaxProps             *uint64              `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	AdditionalProperties AdditionalProperties `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Discriminator        *Discriminator       `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
}

type Types []string

func (pTypes *Types) Is(typ string) bool {
	return pTypes != nil && len(*pTypes) == 1 && (*pTypes)[0] == typ
}

func (pTypes *Types) Slice() []string {
	if pTypes == nil {
		return nil
	}
	return *pTypes
}

func (pTypes *Types) Includes(typ string) bool {
	if pTypes == nil {
		return false
	}
	types := *pTypes
	for _, candidate := range types {
		if candidate == typ {
			return true
		}
	}
	return false
}

func (pTypes *Types) Permits(typ string) bool {
	if pTypes == nil {
		return true
	}
	return pTypes.Includes(typ)
}

func (pTypes *Types) marshal() any {

	if pTypes == nil {
		return nil
	}
	types := *pTypes
	switch len(types) {
	case 0:
		return nil
	case 1:
		return types[0]
	default:
		return []string(types)
	}
}

func (pTypes *Types) MarshalYAML() (interface{}, error) {
	return pTypes.marshal(), nil
}

type AdditionalProperties struct {
	Has    *bool
	Schema *SchemaRef
}

// MarshalYAML returns the YAML encoding of AdditionalProperties.
func (addProps *AdditionalProperties) MarshalYAML() (interface{}, error) {
	if x := addProps.Has; x != nil {
		if *x {
			return true, nil
		}
		return false, nil
	}
	if x := addProps.Schema; x != nil {
		return x.Value, nil
	}
	return nil, nil
}

// MarshalJSON returns the JSON encoding of AdditionalProperties.
func (addProps *AdditionalProperties) MarshalJSON() ([]byte, error) {
	if x := addProps.Has; x != nil {
		if *x {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	}
	if x := addProps.Schema; x != nil {
		return json.Marshal(x)
	}
	return nil, nil
}

func NewSchema() *Schema {
	return &Schema{}
}

// MarshalJSON returns the JSON encoding of Schema.
func (schema *Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(schema.marshal())
}
func (schema *Schema) marshal() any {
	m := schema.extensions.Export(36)

	if x := schema.OneOf; len(x) != 0 {
		m["oneOf"] = x
	}
	if x := schema.AnyOf; len(x) != 0 {
		m["anyOf"] = x
	}
	if x := schema.AllOf; len(x) != 0 {
		m["allOf"] = x
	}
	if x := schema.Not; x != nil {
		m["not"] = x
	}
	if x := schema.Type; x != nil {
		m["type"] = x
	}
	if x := schema.Title; len(x) != 0 {
		m["title"] = x
	}
	if x := schema.Format; len(x) != 0 {
		m["format"] = x
	}
	if x := schema.Description; len(x) != 0 {
		m["description"] = x
	}
	if x := schema.Enum; len(x) != 0 {
		m["enum"] = x
	}
	if x := schema.Default; x != nil {
		m["default"] = x
	}
	if x := schema.Example; x != nil {
		m["example"] = x
	}
	if x := schema.ExternalDocs; x != nil {
		m["externalDocs"] = x
	}

	// Array-related
	if x := schema.UniqueItems; x {
		m["uniqueItems"] = x
	}
	// Number-related
	if x := schema.ExclusiveMin; x {
		m["exclusiveMinimum"] = x
	}
	if x := schema.ExclusiveMax; x {
		m["exclusiveMaximum"] = x
	}
	// Properties
	if x := schema.Nullable; x {
		m["nullable"] = x
	}
	if x := schema.ReadOnly; x {
		m["readOnly"] = x
	}
	if x := schema.WriteOnly; x {
		m["writeOnly"] = x
	}
	if x := schema.AllowEmptyValue; x {
		m["allowEmptyValue"] = x
	}
	if x := schema.Deprecated; x {
		m["deprecated"] = x
	}
	if x := schema.XML; x != nil {
		m["xml"] = x
	}

	// Number
	if x := schema.Min; x != nil {
		m["minimum"] = x
	}
	if x := schema.Max; x != nil {
		m["maximum"] = x
	}
	if x := schema.MultipleOf; x != nil {
		m["multipleOf"] = x
	}

	// String
	if x := schema.MinLength; x != 0 {
		m["minLength"] = x
	}
	if x := schema.MaxLength; x != nil {
		m["maxLength"] = x
	}
	if x := schema.Pattern; x != "" {
		m["pattern"] = x
	}

	// Array
	if x := schema.MinItems; x != 0 {
		m["minItems"] = x
	}
	if x := schema.MaxItems; x != nil {
		m["maxItems"] = x
	}
	if x := schema.Items; x != nil {
		m["items"] = x
	}

	// Object
	if x := schema.Required; len(x) != 0 {
		m["required"] = x
	}
	if x := schema.Properties; len(x) != 0 {
		m["properties"] = x
	}
	if x := schema.MinProps; x != 0 {
		m["minProperties"] = x
	}
	if x := schema.MaxProps; x != nil {
		m["maxProperties"] = x
	}
	if x := schema.AdditionalProperties; x.Has != nil || x.Schema != nil {
		m["additionalProperties"] = &x
	}
	if x := schema.Discriminator; x != nil {
		m["discriminator"] = x
	}

	return m
}

func (pTypes *Types) MarshalJSON() ([]byte, error) { return json.Marshal(pTypes.marshal()) }

func (schema *Schema) NewRef() *SchemaRef {
	return &SchemaRef{
		Value: schema,
	}
}

func NewOneOfSchema(schemas ...*Schema) *Schema {
	refs := make(SchemaRefs, 0, len(schemas))
	for _, schema := range schemas {
		refs = append(refs, &SchemaRef{Value: schema})
	}
	return &Schema{
		OneOf: refs,
	}
}

func NewAnyOfSchema(schemas ...*Schema) *Schema {
	refs := make(SchemaRefs, 0, len(schemas))
	for _, schema := range schemas {
		refs = append(refs, &SchemaRef{Value: schema})
	}
	return &Schema{
		AnyOf: refs,
	}
}

func NewAllOfSchema(schemas ...*Schema) *Schema {
	refs := make(SchemaRefs, 0, len(schemas))
	for _, schema := range schemas {
		refs = append(refs, &SchemaRef{Value: schema})
	}
	return &Schema{
		AllOf: refs,
	}
}

func NewBoolSchema() *Schema {
	return &Schema{
		Type: &Types{TypeBoolean},
	}
}

func NewFloat64Schema() *Schema {
	return &Schema{
		Type: &Types{TypeNumber},
	}
}

func NewIntegerSchema() *Schema {
	return &Schema{
		Type: &Types{TypeInteger},
	}
}

func NewInt32Schema() *Schema {
	return &Schema{
		Type:   &Types{TypeInteger},
		Format: "int32",
	}
}

func NewInt64Schema() *Schema {
	return &Schema{
		Type:   &Types{TypeInteger},
		Format: "int64",
	}
}

func NewStringSchema() *Schema {
	return &Schema{
		Type: &Types{TypeString},
	}
}

func NewDateTimeSchema() *Schema {
	return &Schema{
		Type:   &Types{TypeString},
		Format: "date-time",
	}
}

func NewUUIDSchema() *Schema {
	return &Schema{
		Type:   &Types{TypeString},
		Format: "uuid",
	}
}

func NewBytesSchema() *Schema {
	return &Schema{
		Type:   &Types{TypeString},
		Format: "byte",
	}
}

func NewArraySchema() *Schema {
	return &Schema{
		Type: &Types{TypeArray},
	}
}

func NewObjectSchema() *Schema {
	return &Schema{
		Type:       &Types{TypeObject},
		Properties: make(Schemas),
	}
}

func (schema *Schema) WithNullable() *Schema {
	schema.Nullable = true
	return schema
}

func (schema *Schema) WithMin(value float64) *Schema {
	schema.Min = &value
	return schema
}

func (schema *Schema) WithMax(value float64) *Schema {
	schema.Max = &value
	return schema
}

func (schema *Schema) WithExclusiveMin(value bool) *Schema {
	schema.ExclusiveMin = value
	return schema
}

func (schema *Schema) WithExclusiveMax(value bool) *Schema {
	schema.ExclusiveMax = value
	return schema
}

func (schema *Schema) WithEnum(values ...interface{}) *Schema {
	schema.Enum = values
	return schema
}

func (schema *Schema) WithDefault(defaultValue interface{}) *Schema {
	schema.Default = defaultValue
	return schema
}

func (schema *Schema) WithFormat(value string) *Schema {
	schema.Format = value
	return schema
}

func (schema *Schema) WithLength(i int64) *Schema {
	n := uint64(i)
	schema.MinLength = n
	schema.MaxLength = &n
	return schema
}

func (schema *Schema) WithMinLength(i int64) *Schema {
	n := uint64(i)
	schema.MinLength = n
	return schema
}

func (schema *Schema) WithMaxLength(i int64) *Schema {
	n := uint64(i)
	schema.MaxLength = &n
	return schema
}

func (schema *Schema) WithLengthDecodedBase64(i int64) *Schema {
	n := uint64(i)
	v := (n*8 + 5) / 6
	schema.MinLength = v
	schema.MaxLength = &v
	return schema
}

func (schema *Schema) WithMinLengthDecodedBase64(i int64) *Schema {
	n := uint64(i)
	schema.MinLength = (n*8 + 5) / 6
	return schema
}

func (schema *Schema) WithMaxLengthDecodedBase64(i int64) *Schema {
	n := uint64(i)
	schema.MinLength = (n*8 + 5) / 6
	return schema
}

func (schema *Schema) WithPattern(pattern string) *Schema {
	schema.Pattern = pattern
	return schema
}

func (schema *Schema) WithItems(value *Schema) *Schema {
	schema.Items = &SchemaRef{
		Value: value,
	}
	return schema
}

func (schema *Schema) WithMinItems(i int64) *Schema {
	n := uint64(i)
	schema.MinItems = n
	return schema
}

func (schema *Schema) WithMaxItems(i int64) *Schema {
	n := uint64(i)
	schema.MaxItems = &n
	return schema
}

func (schema *Schema) WithUniqueItems(unique bool) *Schema {
	schema.UniqueItems = unique
	return schema
}

func (schema *Schema) WithProperty(name string, propertySchema *Schema) *Schema {
	return schema.WithPropertyRef(name, &SchemaRef{
		Value: propertySchema,
	})
}

func (schema *Schema) WithPropertyRef(name string, ref *SchemaRef) *Schema {
	properties := schema.Properties
	if properties == nil {
		properties = make(Schemas)
		schema.Properties = properties
	}
	properties[name] = ref
	return schema
}

func (schema *Schema) WithProperties(properties map[string]*Schema) *Schema {
	result := make(Schemas, len(properties))
	for k, v := range properties {
		result[k] = &SchemaRef{
			Value: v,
		}
	}
	schema.Properties = result
	return schema
}

func (schema *Schema) WithRequired(required []string) *Schema {
	schema.Required = required
	return schema
}

func (schema *Schema) WithMinProperties(i int64) *Schema {
	n := uint64(i)
	schema.MinProps = n
	return schema
}

func (schema *Schema) WithMaxProperties(i int64) *Schema {
	n := uint64(i)
	schema.MaxProps = &n
	return schema
}

func (schema *Schema) WithAnyAdditionalProperties() *Schema {
	schema.AdditionalProperties = AdditionalProperties{Has: BoolPtr(true)}
	return schema
}

func (schema *Schema) WithoutAdditionalProperties() *Schema {
	schema.AdditionalProperties = AdditionalProperties{Has: BoolPtr(false)}
	return schema
}

func (schema *Schema) WithAdditionalProperties(v *Schema) *Schema {
	schema.AdditionalProperties = AdditionalProperties{}
	if v != nil {
		schema.AdditionalProperties.Schema = &SchemaRef{Value: v}
	}
	return schema
}

func (schema *Schema) PermitsNull() bool {
	return schema.Nullable || schema.Type.Includes("null")
}

// IsEmpty tells whether schema is equivalent to the empty schema `{}`.
func (schema *Schema) IsEmpty() bool {
	if schema.Type != nil || schema.Format != "" || len(schema.Enum) != 0 ||
		schema.UniqueItems || schema.ExclusiveMin || schema.ExclusiveMax ||
		schema.Nullable || schema.ReadOnly || schema.WriteOnly || schema.AllowEmptyValue ||
		schema.Min != nil || schema.Max != nil || schema.MultipleOf != nil ||
		schema.MinLength != 0 || schema.MaxLength != nil || schema.Pattern != "" ||
		schema.MinItems != 0 || schema.MaxItems != nil ||
		len(schema.Required) != 0 ||
		schema.MinProps != 0 || schema.MaxProps != nil {
		return false
	}
	if n := schema.Not; n != nil && n.Value != nil && !n.Value.IsEmpty() {
		return false
	}
	if ap := schema.AdditionalProperties.Schema; ap != nil && ap.Value != nil && !ap.Value.IsEmpty() {
		return false
	}
	if apa := schema.AdditionalProperties.Has; apa != nil && !*apa {
		return false
	}
	if items := schema.Items; items != nil && items.Value != nil && !items.Value.IsEmpty() {
		return false
	}
	for _, s := range schema.Properties {
		if ss := s.Value; ss != nil && !ss.IsEmpty() {
			return false
		}
	}
	for _, s := range schema.OneOf {
		if ss := s.Value; ss != nil && !ss.IsEmpty() {
			return false
		}
	}
	for _, s := range schema.AnyOf {
		if ss := s.Value; ss != nil && !ss.IsEmpty() {
			return false
		}
	}
	for _, s := range schema.AllOf {
		if ss := s.Value; ss != nil && !ss.IsEmpty() {
			return false
		}
	}
	return true
}

// SchemaError is an error that occurs during schema validation.
type SchemaError struct {
	// Value is the value that failed validation.
	Value interface{}
	// reversePath is the path to the value that failed validation.
	reversePath []string
	// Schema is the schema that failed validation.
	Schema *Schema
	// SchemaField is the field of the schema that failed validation.
	SchemaField string
	// Reason is a human-readable message describing the error.
	// The message should never include the original value to prevent leakage of potentially sensitive inputs in error messages.
	Reason string
	// Origin is the original error that caused this error.
	Origin error
	// customizeMessageError is a function that can be used to customize the error message.
	customizeMessageError func(err *SchemaError) string
}

var _ interface{ Unwrap() error } = (*SchemaError)(nil)

func (err *SchemaError) JSONPointer() []string {
	reversePath := err.reversePath
	path := append([]string(nil), reversePath...)
	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}
	return path
}

func (err *SchemaError) Error() string {
	if err.customizeMessageError != nil {
		if msg := err.customizeMessageError(err); msg != "" {
			return msg
		}
	}

	buf := bytes.NewBuffer(make([]byte, 0, 256))

	if len(err.reversePath) > 0 {
		buf.WriteString(`Error at "`)
		reversePath := err.reversePath
		for i := len(reversePath) - 1; i >= 0; i-- {
			buf.WriteByte('/')
			buf.WriteString(reversePath[i])
		}
		buf.WriteString(`": `)
	}

	if err.Origin != nil {
		buf.WriteString(err.Origin.Error())

		return buf.String()
	}

	reason := err.Reason
	if reason == "" {
		buf.WriteString(`Doesn't match schema "`)
		buf.WriteString(err.SchemaField)
		buf.WriteString(`"`)
	} else {
		buf.WriteString(reason)
	}

	if !SchemaErrorDetailsDisabled {
		buf.WriteString("\nSchema:\n  ")
		encoder := json.NewEncoder(buf)
		encoder.SetIndent("  ", "  ")
		if err := encoder.Encode(err.Schema); err != nil {
			panic(err)
		}
		buf.WriteString("\nValue:\n  ")
		if err := encoder.Encode(err.Value); err != nil {
			panic(err)
		}
	}

	return buf.String()
}

func (err *SchemaError) Unwrap() error {
	return err.Origin
}

func isSliceOfUniqueItems(xs []interface{}) bool {
	s := len(xs)
	m := make(map[string]struct{}, s)
	for _, x := range xs {
		// The input slice is converted from a JSON string, there shall
		// have no error when convert it back.
		key, _ := json.Marshal(&x)
		m[string(key)] = struct{}{}
	}
	return s == len(m)
}

// SliceUniqueItemsChecker is a function used to check if a given slice
// have unique items.
type SliceUniqueItemsChecker func(items []interface{}) bool

// By default, using predefined func isSliceOfUniqueItems which make use of
// json.Marshal to generate a key for map used to check if a given slice
// have unique items.
var sliceUniqueItemsChecker SliceUniqueItemsChecker = isSliceOfUniqueItems

// RegisterArrayUniqueItemsChecker is used to register a customized function
// used to check if JSON array have unique items.
func RegisterArrayUniqueItemsChecker(fn SliceUniqueItemsChecker) {
	sliceUniqueItemsChecker = fn
}

func unsupportedFormat(format string) error {
	return fmt.Errorf("unsupported 'format' value %q", format)
}
