package slimhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	r := NewBaseRouter()
	handler := r.endpointWrapper(func(r *http.Request) (interface{}, error) {
		return newTestStruct(), nil
	}, encodeJSON)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `{"string_key":"string-val","int_key":5,"float_key":13.37}`+"\n")
}

func TestXMLEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	r := NewBaseRouter()
	handler := r.endpointWrapper(func(r *http.Request) (interface{}, error) {
		return newTestStruct(), nil
	}, encodeXML)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `<?xml version="1.0" encoding="UTF-8"?>`+"\n"+`<testStruct><stringKey>string-val</stringKey><intKey>5</intKey><floatKey>13.37</floatKey></testStruct>`)
}

func TestTextEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	r := NewBaseRouter()
	handler := r.endpointWrapper(func(r *http.Request) (interface{}, error) {
		return "Here's some text!", nil
	}, encodeText)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, "Here's some text!")
}

func TestHTMLEndpointWrapper(t *testing.T) {
	wr := httptest.NewRecorder()
	r := NewBaseRouter()
	handler := r.endpointWrapper(func(r *http.Request) (interface{}, error) {
		return newTestWebpage(), nil
	}, encodeHTML)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, "<string-test>string-val</string-test><br/><int-test>5</int-test><br/><float-test>13.37</float-test>")
}
