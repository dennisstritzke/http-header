package HttpHeader

import (
	"net/http"
	"testing"
)

func TestBind_nil(t *testing.T) {
	err := Bind(http.Request{}, nil)
	checkForInvalidBindError(t, err)
}

func TestBind_nonPointerType(t *testing.T) {
	err := Bind(http.Request{}, "")
	checkForInvalidBindError(t, err)
}

func checkForInvalidBindError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an error but got nil.")
	}

	if _, ok := err.(*InvalidBindError); !ok {
		t.Error("Expected InvalidBindError, but got something different")
	}
}
