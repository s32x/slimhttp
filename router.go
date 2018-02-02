package slimhttp

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router defines all functionality for our api service router
type Router interface {
	HandleStatic(path string) *mux.Route
	HandleEndpoint(pattern string, endpoint Endpoint) *mux.Route
	ListenAndServe(port string) error
}

// Router contains an embedded router that we can use to bind
// Endpoints to
type router struct {
	router *mux.Router
	encode Encoder
}

// NewRouter generates a new Router that will be used to bind
// handlers to the *mux.Router
func NewRouter(encoder Encoder) Router {
	return &router{
		router: mux.NewRouter(),
		encode: encoder,
	}
}

// NewJSONRouter generates a new mux.Router using the encodeJSON
// encoder for binding endpoints that encode JSON responses
func NewJSONRouter() Router { return NewRouter(encodeJSON) }

// NewXMLRouter generates a new mux.Router using the encodeXML
// encoder for binding endpoints that encode XML responses
func NewXMLRouter() Router { return NewRouter(encodeXML) }

// HandleStatic binds a new fileserver using the passed path to
// the router
func (r *router) HandleStatic(path string) *mux.Route {
	return r.router.PathPrefix("/").Handler(http.FileServer(http.Dir(path)))
}

// HandleEndpoint binds a new Endpoint handler to the router
func (r *router) HandleEndpoint(pattern string, endpoint Endpoint) *mux.Route {
	return r.router.HandleFunc(pattern, r.endpointWrapper(endpoint))
}

// ListenAndServe applies basic CORS headers to an http.Server
// and begins listening for requests
func (r *router) ListenAndServe(port string) error {
	// Create the basic http.Server with base parameters
	srv := &http.Server{
		Handler:      r.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + port,
	}

	srv.Handler = handlers.CORS(
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "GET", "OPTIONS", "HEAD"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r.router)

	// Begin listening for requests
	return srv.ListenAndServe()
}
