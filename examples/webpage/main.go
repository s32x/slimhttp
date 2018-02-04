package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sdwolfe32/slimhttp"
)

func main() {
	r := slimhttp.NewRouter()                    // Create a new router
	r.HandleHTMLEndpoint("/hello/{name}", Hello) // Bind an Endpoint to the router at the specified path
	log.Fatal(r.ListenAndServe(8080))            // Start the service!
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
func Hello(r *http.Request) (interface{}, error) {
	name := mux.Vars(r)["name"] // The name passed on the request

	switch name {
	case "basic-error":
		// An example of returning a raw error
		err := errors.New("This is a basic error")
		return nil, err
	case "standard-error":
		// An example of returning a predefined Error
		return nil, slimhttp.ErrorBadRequest
	case "fancy-error":
		// An example of returning a fully self-defined Error
		err := errors.New("This is a fancy error")
		return nil, slimhttp.NewError("This is a fancy error!", http.StatusBadRequest, err)
	}

	// Parse a basic template for using as the output
	temp, err := template.New("hello").Parse("<b>{{.Message}}</b><br/>{{.Success}}")
	if err != nil {
		return nil, err
	}

	// All other names will be returned on a HelloResponse
	return &slimhttp.Webpage{
		Template: temp,
		Data: &HelloResponse{
			Message: fmt.Sprintf("Hello %s!", name),
			Success: true,
		},
	}, nil
}
