package slimhttp

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router defines all functionality for our api service router
type Router interface {
	HandleStatic(prefix, path string) *mux.Route
	HandleEndpoint(pattern string, endpoint Endpoint, encoder Encoder) *mux.Route
	HandleJSONEndpoint(pattern string, endpoint Endpoint) *mux.Route
	HandleXMLEndpoint(pattern string, endpoint Endpoint) *mux.Route
	ListenAndServe(port string) error
}

// Router contains an mux.Router that we can use to bind Endpoints to
type router struct{ router *mux.Router }

// NewRouter generates a new Router that will be used to bind
// handlers to the *mux.Router
func NewRouter() Router {
	return &router{router: mux.NewRouter()}
}

// HandleStatic binds a new fileserver using the passed prefix and
// path to the router
func (r *router) HandleStatic(prefix, path string) *mux.Route {
	fs := http.FileServer(http.Dir(path))
	return r.router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, fs))
}

// HandleEndpoint binds a new Endpoint and Encoder handler to the router
func (r *router) HandleEndpoint(pattern string, endpoint Endpoint, encoder Encoder) *mux.Route {
	return r.router.HandleFunc(pattern, endpointWrapper(endpoint, encoder))
}

// HandleJSONEndpoint binds a new JSON Endpoint handler to the router
func (r *router) HandleJSONEndpoint(pattern string, endpoint Endpoint) *mux.Route {
	return r.router.HandleFunc(pattern, endpointWrapper(endpoint, encodeJSON))
}

// HandleXMLEndpoint binds a new XML Endpoint handler to the router
func (r *router) HandleXMLEndpoint(pattern string, endpoint Endpoint) *mux.Route {
	return r.router.HandleFunc(pattern, endpointWrapper(endpoint, encodeXML))
}

// ListenAndServe applies basic CORS headers to a new http.Server
// and begins listening for requests
func (r *router) ListenAndServe(port string) error {
	// Create the basic http.Server with base parameters
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Apply CORS headers
	srv.Handler = handlers.CORS(
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "GET", "OPTIONS", "HEAD"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r.router)

	// Begin listening for requests
	return srv.ListenAndServe()
}
