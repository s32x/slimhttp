package main

import (
	"log"

	"github.com/s32x/slimhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new router
	r := slimhttp.NewBaseRouter()

	logger := logrus.New()
	s := NewHelloService(logger)
	h := slimhttp.NewHealthcheckService(logger, "api.example.com")

	// Bind an Endpoint to the router at the specified path
	r.HandleJSONEndpoint("/healtcheck", h.Healthcheck)
	r.HandleJSONEndpoint("/hello/{name}", s.Hello)

	// Start the service!
	log.Fatal(r.ListenAndServe(8080))
}
