package slimhttp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeText = "text/plain; charset=utf-8"
	contentTypeHTML = "text/html; charset=utf-8"
	contentTypeJSON = "application/json; charset=utf-8"
	contentTypeXML  = "application/xml; charset=utf-8"
)

// An Encoder is a function that will write a response to the passed
// ResponseWriter using the provided statuscode and response struct
type Encoder func(w http.ResponseWriter, status int, res interface{})

// encodeText encodes the response as Text and writes it to the ResponseWriter
func encodeText(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set(contentTypeKey, contentTypeText)
	body, ok := res.(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Response type should be string"))
		return
	}
	w.WriteHeader(status)
	w.Write([]byte(body))
}

// encodeJSON encodes the response as JSON and writes it to the ResponseWriter
func encodeJSON(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set(contentTypeKey, contentTypeJSON)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

// encodeXML encodes the response as XML and writes it to the ResponseWriter
func encodeXML(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set(contentTypeKey, contentTypeXML)
	b, err := xml.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to XML marshal response struct"))
		return
	}
	w.WriteHeader(status)
	w.Write(append([]byte(xml.Header), b...))
}
