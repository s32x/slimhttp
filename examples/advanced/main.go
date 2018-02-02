package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/sdwolfe32/slimhttp"
)

func main() {
	// Create a new router
	r := slimhttp.NewJSONRouter()

	logger := logrus.New()
	s := NewHelloer(logger)
	h := slimhttp.NewHealthchecker(logger, "api.example.com")

	// Bind an Endpoint to the router at the specified path
	r.HandleEndpoint("/healtcheck", h.Healthcheck)
	r.HandleEndpoint("/hello/{name}/", s.Hello)

	// Start the service!
	r.ListenAndServe("8080")
}
