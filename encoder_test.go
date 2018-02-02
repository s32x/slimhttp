package slimhttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncodeJSON(t *testing.T) {
	wr := httptest.NewRecorder()
	encodeJSON(wr, http.StatusOK, newTestStruct())

	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, fmt.Sprintf("%s\n", `{"string_key":"string-val","int_key":5,"float_key":13.37}`))
}

func TestEncodeXML(t *testing.T) {
	wr := httptest.NewRecorder()
	encodeXML(wr, http.StatusOK, newTestStruct())

	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, `<testStruct><stringKey>string-val</stringKey><intKey>5</intKey><floatKey>13.37</floatKey></testStruct>`)
}
