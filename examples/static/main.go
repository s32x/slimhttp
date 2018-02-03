package main

import (
	"log"

	"github.com/sdwolfe32/slimhttp"
)

func main() {
	r := slimhttp.NewRouter()                // Create a new router
	r.HandleStatic("/path-in-url", "assets") // Bind a static path to a static directory
	log.Fatal(r.ListenAndServe(8080))        // Start the service!
}
