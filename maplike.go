package openapi3

import (
	"encoding/json"
)

// NewResponsesWithCapacity builds a responses object of the given capacity.
func NewResponsesWithCapacity(cap int) *Responses {
	if cap == 0 {
		return &Responses{m: make(map[string]*ResponseRef)}
	}
	return &Responses{m: make(map[string]*ResponseRef, cap)}
}

// Value returns the responses for key or nil
func (responses *Responses) Value(key string) *ResponseRef {
	if responses.Len() == 0 {
		return nil
	}
	return responses.m[key]
}

// Set adds or replaces key 'key' of 'responses' with 'value'.
// Note: 'responses' MUST be non-nil
func (responses *Responses) Set(key string, value *ResponseRef) {
	if responses.m == nil {
		responses.m = make(map[string]*ResponseRef)
	}
	responses.m[key] = value
}

// Len returns the amount of keys in responses excluding responses.extensions.
func (responses *Responses) Len() int {
	if responses == nil || responses.m == nil {
		return 0
	}
	return len(responses.m)
}

// Delete removes the entry associated with key 'key' from 'responses'.
func (responses *Responses) Delete(key string) {
	if responses != nil && responses.m != nil {
		delete(responses.m, key)
	}
}

// Map returns responses as a 'map'.
// Note: iteration on Go maps is not ordered.
func (responses *Responses) Map() (m map[string]*ResponseRef) {
	if responses == nil || len(responses.m) == 0 {
		return make(map[string]*ResponseRef)
	}
	m = make(map[string]*ResponseRef, len(responses.m))
	for k, v := range responses.m {
		m[k] = v
	}
	return
}

// MarshalJSON returns the JSON encoding of Responses.
func (responses *Responses) marshal() any {

	m := responses.extensions.Export(responses.Len())
	for k, v := range responses.Map() {
		m[k] = v
	}

	return m
}
func (responses *Responses) MarshalYAML() (interface{}, error) {
	return responses.marshal(), nil
}
func (responses *Responses) MarshalJSON() ([]byte, error) { return json.Marshal(responses.marshal()) }

// NewCallbackWithCapacity builds a callback object of the given capacity.
func NewCallbackWithCapacity(cap int) *Callback {
	if cap == 0 {
		return &Callback{m: make(map[string]*PathItem)}
	}
	return &Callback{m: make(map[string]*PathItem, cap)}
}

// Value returns the callback for key or nil
func (callback *Callback) Value(key string) *PathItem {
	if callback.Len() == 0 {
		return nil
	}
	return callback.m[key]
}

// Set adds or replaces key 'key' of 'callback' with 'value'.
// Note: 'callback' MUST be non-nil
func (callback *Callback) Set(key string, value *PathItem) {
	if callback.m == nil {
		callback.m = make(map[string]*PathItem)
	}
	callback.m[key] = value
}

// Len returns the amount of keys in callback excluding callback.extensions.
func (callback *Callback) Len() int {
	if callback == nil || callback.m == nil {
		return 0
	}
	return len(callback.m)
}

// Delete removes the entry associated with key 'key' from 'callback'.
func (callback *Callback) Delete(key string) {
	if callback != nil && callback.m != nil {
		delete(callback.m, key)
	}
}

// Map returns callback as a 'map'.
// Note: iteration on Go maps is not ordered.
func (callback *Callback) Map() (m map[string]*PathItem) {
	if callback == nil || len(callback.m) == 0 {
		return make(map[string]*PathItem)
	}
	m = make(map[string]*PathItem, len(callback.m))
	for k, v := range callback.m {
		m[k] = v
	}
	return
}

// MarshalJSON returns the JSON encoding of Callback.
func (callback *Callback) MarshalJSON() ([]byte, error) {

	return json.Marshal(callback.marshal())
}

// NewPathsWithCapacity builds a paths object of the given capacity.
func NewPathsWithCapacity(cap int) *Paths {
	if cap == 0 {
		return &Paths{m: make(map[string]*PathItem)}
	}
	return &Paths{m: make(map[string]*PathItem, cap)}
}

// Value returns the paths for key or nil
func (paths *Paths) Value(key string) *PathItem {
	if paths.Len() == 0 {
		return nil
	}
	return paths.m[key]
}

// Set adds or replaces key 'key' of 'paths' with 'value'.
// Note: 'paths' MUST be non-nil
func (paths *Paths) Set(key string, value *PathItem) {
	if paths.m == nil {
		paths.m = make(map[string]*PathItem)
	}
	paths.m[key] = value
}

// Len returns the amount of keys in paths excluding paths.extensions.
func (paths *Paths) Len() int {
	if paths == nil || paths.m == nil {
		return 0
	}
	return len(paths.m)
}

// Delete removes the entry associated with key 'key' from 'paths'.
func (paths *Paths) Delete(key string) {
	if paths != nil && paths.m != nil {
		delete(paths.m, key)
	}
}

// Map returns paths as a 'map'.
// Note: iteration on Go maps is not ordered.
func (paths *Paths) Map() (m map[string]*PathItem) {
	if paths == nil || len(paths.m) == 0 {
		return make(map[string]*PathItem)
	}
	m = make(map[string]*PathItem, len(paths.m))
	for k, v := range paths.m {
		m[k] = v
	}
	return
}

// MarshalJSON returns the JSON encoding of Paths.
func (paths *Paths) MarshalJSON() ([]byte, error) {

	return json.Marshal(paths.marshal())
}
