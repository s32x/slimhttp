package slimhttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandleStatic binds a new fileserver using the passed prefix and
// path to the router
func (r *Router) HandleStatic(prefix, path string) *mux.Route {
	fs := http.FileServer(http.Dir(path))
	return r.Router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, fs))
}

// HandleEndpoint binds a new Endpoint and Encoder handler to the router
func (r *Router) HandleEndpoint(pattern string, endpoint Endpoint, encoder Encoder, opts ...Option) *mux.Route {
	return r.Router.HandleFunc(pattern, r.endpointWrapper(endpoint, encoder, opts...))
}

// HandleJSONEndpoint binds a new JSON Endpoint handler to the router
func (r *Router) HandleJSONEndpoint(pattern string, endpoint Endpoint, opts ...Option) *mux.Route {
	return r.Router.HandleFunc(pattern, r.endpointWrapper(endpoint, encodeJSON, opts...))
}

// HandleXMLEndpoint binds a new XML Endpoint handler to the router
func (r *Router) HandleXMLEndpoint(pattern string, endpoint Endpoint, opts ...Option) *mux.Route {
	return r.Router.HandleFunc(pattern, r.endpointWrapper(endpoint, encodeXML, opts...))
}

// HandleTextEndpoint binds a new Text Endpoint handler to the router
func (r *Router) HandleTextEndpoint(pattern string, endpoint Endpoint, opts ...Option) *mux.Route {
	return r.Router.HandleFunc(pattern, r.endpointWrapper(endpoint, encodeText, opts...))
}

// HandleBytesEndpoint binds a new Bytes Endpoint handler to the router
func (r *Router) HandleBytesEndpoint(pattern string, endpoint Endpoint, opts ...Option) *mux.Route {
	return r.Router.HandleFunc(pattern, r.endpointWrapper(endpoint, encodeBytes, opts...))
}

// HandleHTMLEndpoint binds a new HTML Endpoint handler to the router
func (r *Router) HandleHTMLEndpoint(pattern string, endpoint Endpoint, opts ...Option) *mux.Route {
	return r.Router.HandleFunc(pattern, r.endpointWrapper(endpoint, encodeHTML, opts...))
}
