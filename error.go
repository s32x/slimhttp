package slimhttp

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var (
	ErrorBadRequest          = NewError("Bad Request", http.StatusBadRequest, nil)
	ErrorBadGateway          = NewError("Bad Gateway", http.StatusBadGateway, nil)
	ErrorForbidden           = NewError("Forbidden", http.StatusForbidden, nil)
	ErrorMovedPermanently    = NewError("Moved Permanently", http.StatusMovedPermanently, nil)
	ErrorNotFound            = NewError("Not Found", http.StatusNotFound, nil)
	ErrorInternalServerError = NewError("Internal Server Error", http.StatusInternalServerError, nil)
)

// Error is a standard API error that should be used for
// all API error responses
type Error struct {
	XMLName    xml.Name `json:"-" xml:"error"`
	Message    string   `json:"message,omitempty" xml:"message,omitempty"`
	StatusCode int      `json:"status,omitempty" xml:"status,omitempty"`
	Err        string   `json:"err,omitempty" xml:"err,omitempty"`
}

// NewError is a function that generates a new Error
func NewError(message string, statusCode int, err error) *Error {
	var errString string
	if err != nil {
		errString = err.Error()
	}
	return &Error{
		Message:    message,
		StatusCode: statusCode,
		Err:        errString,
	}
}

// Error returns a string representation of the error and
// helps to satisfy the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Error %d: '%s'", e.StatusCode, e.Message)
}

// Log will log the Error before returning it
func (e *Error) Log(log *logrus.Entry) *Error {
	log.WithFields(logrus.Fields{"error": e.Err, "status": e.StatusCode}).Error(e.Message)
	return e
}
