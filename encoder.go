package slimhttp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// An Encoder is a function that will write a response to the passed
// ResponseWriter using the provided statuscode and response struct
type Encoder func(w http.ResponseWriter, status int, res interface{})

// encodeJSON encodes the response to JSON and writes it to the ResponseWriter
func encodeJSON(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

// encodeXML encodes the response to XML and writes it to the ResponseWriter
func encodeXML(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(status)
	xml.NewEncoder(w).Encode(res)
}
