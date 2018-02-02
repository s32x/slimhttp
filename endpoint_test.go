package slimhttp

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestEndpointWrapper(t *testing.T) {
	r := newTestRouter()
	wr := httptest.NewRecorder()
	handler := r.endpointWrapper(newTestEndpoint)
	handler(wr, nil)
	body, status, err := result(wr)
	equal(t, err, nil)
	equal(t, status, 200)
	equal(t, body, fmt.Sprintf("%s\n", `{"string_key":"string-val","int_key":5,"float_key":13.37}`))
}

// func TestEndpointStandardError(t *testing.T) {
// 	wr := httptest.NewRecorder()

// 	encodeJSON(wr, http.StatusOK, TestStruct{
// 		StringKey: "string-val",
// 		IntKey:    5,
// 		FloatKey:  13.37,
// 	})

// 	body, status, err := result(wr)
// 	equal(t, err, nil)
// 	equal(t, status, 200)
// 	equal(t, body, fmt.Sprintf("%s\n", `{"string_key":"string-val","int_key":5,"float_key":13.37}`))
// }

// func TestEndpointDetailedError(t *testing.T) {
// 	wr := httptest.NewRecorder()

// 	encodeJSON(wr, http.StatusOK, TestStruct{
// 		StringKey: "string-val",
// 		IntKey:    5,
// 		FloatKey:  13.37,
// 	})

// 	body, status, err := result(wr)
// 	equal(t, err, nil)
// 	equal(t, status, 200)
// 	equal(t, body, fmt.Sprintf("%s\n", `{"string_key":"string-val","int_key":5,"float_key":13.37}`))
// }
