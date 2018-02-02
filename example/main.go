package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdwolfe32/slimhttp"
)

// Output is an example output struct that will be
// encoded to JSON on the response
type Output struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func main() {
	// Create a new router
	r := slimhttp.NewJSONRouter()

	// Bind an Endpoint to the router at the specified path
	r.HandleEndpoint("/{name}/", Hello)

	// Start the service!
	r.ListenAndServe("8080")
}

// Hello is an example Endpoint method. It receives a
// request so that you have access to everything on the
// request and returns a successful body or error
func Hello(r *http.Request) (interface{}, error) {
	name := mux.Vars(r)["name"] // The name passed on the request

	// "fancy-error" as the name invokes and returns a fully
	// encoded slimhttp.Error which is created here
	if name == "fancy-error" {
		err := errors.New("This is a fancy error")
		return nil, slimhttp.NewError("There's a very bad error!", http.StatusBadRequest, err)
	}

	// "basic-error" as the name invokes and returns a fully
	// encoded slimhttp.Error that is generated in the wrapper
	if name == "basic-error" {
		err := errors.New("This is a basic error")
		return nil, err
	}

	// All other names will pass through and return a fully
	// encoded Output
	return &Output{
		Message: fmt.Sprintf("Hello %s!", name),
		Success: true,
	}, nil
}
