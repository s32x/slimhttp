package slimhttp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"text/template"
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeJSON = "application/json; charset=utf-8"
	contentTypeXML  = "application/xml; charset=utf-8"
	contentTypeText = "text/plain; charset=utf-8"
	contentTypeHTML = "text/html; charset=utf-8"
)

// An Encoder is a function that will write a response to the passed
// ResponseWriter using the provided statuscode and response struct
type Encoder func(w http.ResponseWriter, status int, res interface{})

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
		fmt.Fprint(w, "Failed to XML marshal response struct")
		return
	}
	w.WriteHeader(status)
	fmt.Fprint(w, xml.Header+string(b))
}

// encodeText expects a res of type string, encodes the response as
// Text and writes it to the ResponseWriter
func encodeText(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set(contentTypeKey, contentTypeText)
	body, ok := res.(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Response type should be (string)")
		return
	}
	w.WriteHeader(status)
	fmt.Fprint(w, body)
}

// encodeBytes expects a res interface of type io.Reader and copies
// the bytes read from the reader to the ResponseWriter
func encodeBytes(w http.ResponseWriter, status int, res interface{}) {
	body, ok := res.([]byte)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Response type should be ([]byte)")
		return
	}
	w.WriteHeader(status)
	w.Write(body)
}

// Webpage represents an HTML webpage response. It contains a template
// and data to be written into the template
type Webpage struct {
	Template *template.Template
	Data     interface{}
}

// encodeHTML expects a res struct of type *template.Template, encodes
// the response as HTML and writes it to the ResponseWriter
func encodeHTML(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set(contentTypeKey, contentTypeHTML)
	wp, ok := res.(*Webpage)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Response type should be (*Webpage)")
		return
	}
	w.WriteHeader(status)
	wp.Template.Execute(w, wp.Data)
}
