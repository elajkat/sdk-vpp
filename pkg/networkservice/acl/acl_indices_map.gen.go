// Code generated by "-output acl_indices_map.gen.go -type aclIndicesMap<string,[]uint32> -output acl_indices_map.gen.go -type aclIndicesMap<string,[]uint32>"; DO NOT EDIT.
// Install -output acl_indices_map.gen.go -type aclIndicesMap<string,[]uint32> by "go get -u github.com/searKing/golang/tools/-output acl_indices_map.gen.go -type aclIndicesMap<string,[]uint32>"

package acl

import (
	"sync" // Used by sync.Map.
)

// Generate code that will fail if the constants change value.
func _() {
	// An "cannot convert aclIndicesMap literal (type aclIndicesMap) to type sync.Map" compiler error signifies that the base type have changed.
	// Re-run the go-syncmap command to generate them again.
	_ = (sync.Map)(aclIndicesMap{})
}

var _nil_aclIndicesMap_uint32_value = func() (val []uint32) { return }()

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *aclIndicesMap) Load(key string) ([]uint32, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if value == nil {
		return _nil_aclIndicesMap_uint32_value, ok
	}
	return value.([]uint32), ok
}

// Store sets the value for a key.
func (m *aclIndicesMap) Store(key string, value []uint32) {
	(*sync.Map)(m).Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *aclIndicesMap) LoadOrStore(key string, value []uint32) ([]uint32, bool) {
	actual, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if actual == nil {
		return _nil_aclIndicesMap_uint32_value, loaded
	}
	return actual.([]uint32), loaded
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *aclIndicesMap) LoadAndDelete(key string) (value []uint32, loaded bool) {
	actual, loaded := (*sync.Map)(m).LoadAndDelete(key)
	if actual == nil {
		return _nil_aclIndicesMap_uint32_value, loaded
	}
	return actual.([]uint32), loaded
}

// Delete deletes the value for a key.
func (m *aclIndicesMap) Delete(key string) {
	(*sync.Map)(m).Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *aclIndicesMap) Range(f func(key string, value []uint32) bool) {
	(*sync.Map)(m).Range(func(key, value interface{}) bool {
		return f(key.(string), value.([]uint32))
	})
}
