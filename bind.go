package HttpHeader

import (
	"net/http"
	"reflect"
)

// An InvalidBindError describes an invalid argument passed to Bind.
// (The argument to Bind must be a non-nil pointer.)
type InvalidBindError struct {
	Type reflect.Type
}

func (e *InvalidBindError) Error() string {
	if e.Type == nil {
		return "header: Bind(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "header: Bind(non-pointer " + e.Type.String() + ")"
	}

	return "header: Bind(nil " + e.Type.String() + ")"
}

// Bind processes the HTTP header fields and stores the result in the value pointed to by v.
func Bind(r http.Request, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidBindError{reflect.TypeOf(v)}
	}

	return nil
}
