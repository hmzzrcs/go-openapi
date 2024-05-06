package openapi3

import (
	"strings"
)

var (
	_ baseMarshaller = (*Paths)(nil)
)

// Paths is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#paths-object
type Paths struct {
	extensions
	m map[string]*PathItem
}

// NewPaths builds a paths object with path items in insertion order.
func NewPaths(opts ...NewPathsOption) *Paths {
	paths := NewPathsWithCapacity(len(opts))
	for _, opt := range opts {
		opt(paths)
	}
	return paths
}

// NewPathsOption describes options to NewPaths func
type NewPathsOption func(*Paths)

// WithPath adds a named path item
func WithPath(path string, pathItem *PathItem) NewPathsOption {
	return func(paths *Paths) {
		if p := pathItem; p != nil && path != "" {
			paths.Set(path, p)
		}
	}
}

// Support YAML Marshaler interface for gopkg.in/yaml
func (paths *Paths) MarshalYAML() (any, error) {
	return paths.marshal(), nil
}
func (paths *Paths) marshal() any {
	res := paths.extensions.Export(len(paths.m))

	for k, v := range paths.m {
		res[k] = v
	}

	return res
}

func normalizeTemplatedPath(path string) (string, uint, map[string]struct{}) {
	if strings.IndexByte(path, '{') < 0 {
		return path, 0, nil
	}

	var buffTpl strings.Builder
	buffTpl.Grow(len(path))

	var (
		cc         rune
		count      uint
		isVariable bool
		vars       = make(map[string]struct{})
		buffVar    strings.Builder
	)
	for i, c := range path {
		if isVariable {
			if c == '}' {
				// End path variable
				isVariable = false

				vars[buffVar.String()] = struct{}{}
				buffVar = strings.Builder{}

				// First append possible '*' before this character
				// The character '}' will be appended
				if i > 0 && cc == '*' {
					buffTpl.WriteRune(cc)
				}
			} else {
				buffVar.WriteRune(c)
				continue
			}

		} else if c == '{' {
			// Begin path variable
			isVariable = true

			// The character '{' will be appended
			count++
		}

		// Append the character
		buffTpl.WriteRune(c)
		cc = c
	}
	return buffTpl.String(), count, vars
}
