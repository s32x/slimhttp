package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/sdwolfe32/slimhttp"
)

// Helloer defines all functionality for a Helloer
// service
type Helloer interface {
	Hello(r *http.Request) (interface{}, error)
}

// helloer contains all dependencies for a Helloer
type helloer struct{ log *logrus.Entry }

// NewHelloer generates a new Helloer service
func NewHelloer(log *logrus.Logger) Helloer {
	return &helloer{log: log.WithField("service", "helloer")}
}

// Output is an example output struct that will be
// encoded to JSON on the response
type Output struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// Hello is an example Endpoint method. It receives a
// request so that you have access to everything on the
// request and returns a successful body or error
func (h *helloer) Hello(r *http.Request) (interface{}, error) {
	h.log.Debug("New Hello request received")
	name := mux.Vars(r)["name"] // The name passed on the request
	l := h.log.WithField("name", name)

	// "fancy-error" as the name invokes and returns a fully
	// encoded slimhttp.Error which is created here
	if name == "fancy-error" {
		err := errors.New("This is a fancy error")
		return nil, slimhttp.NewError("There's a very bad error!", http.StatusBadRequest, err).Log(l)
	}

	// "basic-error" as the name invokes and returns a fully
	// encoded slimhttp.Error that is generated in the wrapper
	if name == "basic-error" {
		err := errors.New("This is a basic error")
		l.WithError(err).Error(err)
		return nil, err
	}

	// All other names will pass through and return a fully
	// encoded Output
	l.Debug("Returning new Output response")
	return &Output{
		Message: fmt.Sprintf("Hello %s!", name),
		Success: true,
	}, nil
}
