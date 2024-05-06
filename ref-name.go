package openapi3

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	refNameConfig = map[string]string{
		typeName[RequestBody]():    "requestBodies",
		typeName[SecurityScheme](): "securitySchemes",
	}
)

func refNames[T any](vs ...*T) string {

	tn := typeName[T]()
	names, has := refNameConfig[tn]
	if has {
		return names
	}
	return fmt.Sprintf("%ss", strings.ToLower(names))
}
func typeName[T any]() string {
	var v *T = nil
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()

}
