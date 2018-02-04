package slimhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncodeText(t *testing.T) {
	wr := httptest.NewRecorder()
	encodeText(wr, http.StatusOK, "Here's some text!")

	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, "Here's some text!")
}

func TestEncodeJSON(t *testing.T) {
	wr := httptest.NewRecorder()
	encodeJSON(wr, http.StatusOK, newTestStruct())

	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `{"string_key":"string-val","int_key":5,"float_key":13.37}`+"\n")
}

func TestEncodeXML(t *testing.T) {
	wr := httptest.NewRecorder()
	encodeXML(wr, http.StatusOK, newTestStruct())

	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `<?xml version="1.0" encoding="UTF-8"?>`+"\n"+`<testStruct><stringKey>string-val</stringKey><intKey>5</intKey><floatKey>13.37</floatKey></testStruct>`)
}
