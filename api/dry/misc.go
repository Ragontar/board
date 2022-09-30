package dry

import (
	"reflect"
)

// CheckNonNil returns True if none of the arguments are nil-pointer. Returns False if at least one argument
// is nil or non-pointer arg was passed.
func CheckNonNil(ptrs ...interface{}) bool {
	for _, ptr := range ptrs {
		if reflect.ValueOf(ptr).Kind() != reflect.Ptr {
			return false
		}
		if ptr == nil {
			return false
		}
	}

	return true
}
