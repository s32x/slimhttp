package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s32x/slimhttp"
	"github.com/sirupsen/logrus"
)

// HelloService defines all functionality for a helloService
type HelloService interface {
	Hello(r *http.Request) (interface{}, error)
}

// helloService contains all dependencies for a HelloService
type helloService struct{ log *logrus.Entry }

// NewHelloService generates a new HelloService
func NewHelloService(log *logrus.Logger) HelloService {
	return &helloService{log: log.WithField("service", "hello")}
}

// HelloResponse is an example response struct that will be
// encoded to JSON on a Hello request
type HelloResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// Hello is an example Endpoint method. It receives a request
// so that you have access to everything on the request and
// returns a successful body or error
func (h *helloService) Hello(r *http.Request) (interface{}, error) {
	h.log.Debug("New Hello request received")
	name := mux.Vars(r)["name"] // The name passed on the request
	l := h.log.WithField("name", name)

	switch name {
	case "basic-error":
		// An example of returning a raw error (no logging)
		err := errors.New("This is a basic error")
		return nil, err
	case "standard-error":
		// An example of logging and returning a predefined Error
		return nil, slimhttp.ErrorBadRequest.Log(l)
	case "fancy-error":
		// An example of logging and returning a fully self-defined Error
		err := errors.New("This is a fancy error")
		return nil, slimhttp.NewError("This is a fancy error!", http.StatusBadRequest, err).Log(l)
	}

	// All other names will pass through and return a fully
	// encoded Output
	l.Debug("Returning new Output response")
	return &HelloResponse{
		Message: fmt.Sprintf("Hello %s!", name),
		Success: true,
	}, nil
}
