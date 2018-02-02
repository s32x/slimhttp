package slimhttp

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	actual := NewError(
		"test-error-message",
		http.StatusInternalServerError,
		errors.New("This is an error"),
	)
	equal(t, actual.Message, "test-error-message")
	equal(t, actual.StatusCode, 500)
	equal(t, actual.Err, "This is an error")
}

func TestErrorLogging(t *testing.T) {
	actual := NewError(
		"test-error-message",
		http.StatusInternalServerError,
		errors.New("This is an error"),
	)
	equal(t, actual.Error(), "Error 500: 'test-error-message'")
}
