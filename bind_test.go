package HttpHeader

import (
	"net/http"
	"testing"
)

func createRequest(omit ...string) http.Request {
	headers := map[string][]string{
		"Foo": {"Bar"},
		"Bar": {"Foo", "42"},
	}

	for _, headerToOmit := range omit {
		delete(headers, headerToOmit)
	}

	return http.Request{
		Header: headers,
	}
}

func TestBind_nil(t *testing.T) {
	err := Bind(http.Request{}, nil)
	checkForInvalidBindError(t, err)
}

func TestBind_nonPointerType(t *testing.T) {
	err := Bind(http.Request{}, "")
	checkForInvalidBindError(t, err)
}

func TestBind_String(t *testing.T) {
	type Test struct {
		Foo string `header:"Foo"`
	}

	var test Test

	err := Bind(createRequest(), &test)
	if err != nil {
		t.Error(err)
		return
	}

	expected := "Bar"
	actual := test.Foo
	if actual != expected {
		t.Errorf("Expected field to have value '%s', but got '%s'...", expected, actual)
	}
}

func TestBind_StringSlice(t *testing.T) {
	type Test struct {
		Bar []string `header:"Bar"`
	}

	var test Test

	err := Bind(createRequest(), &test)
	if err != nil {
		t.Error(err)
		return
	}

	if len(test.Bar) != 2 {
		t.Errorf("Expected two values, but got %d: %s", len(test.Bar), test.Bar)
		return
	}

	if test.Bar[0] != "Foo" {
		t.Errorf("Expected first element to be 'Foo', but was '%s'...", test.Bar[0])
	}

	if test.Bar[1] != "42" {
		t.Errorf("Expected first element to be '42', but was '%s'...", test.Bar[1])
	}
}

func checkForInvalidBindError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an error but got nil.")
	}

	if _, ok := err.(*InvalidBindError); !ok {
		t.Error("Expected InvalidBindError, but got something different")
	}
}
