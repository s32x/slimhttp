package slimhttp

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

// Router contains an mux.Router that we can use to bind Endpoints to
type Router struct {
	Router *mux.Router
	Server *http.Server
}

// NewRouter generates a new router containing a mux.Router and
// an http.Server containing a default timeout of 30 seconds for
// reading/writing
func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
		Server: &http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

// WithTimeout sets a custom read and write timeout on the routers
// http.Server
func (r *Router) WithTimeout(read, write time.Duration) *Router {
	r.Server.ReadTimeout = read
	r.Server.WriteTimeout = write
	return r
}

// HandleStatic binds a new fileserver using the passed prefix and
// path to the router
func (r *Router) HandleStatic(prefix, path string) *mux.Route {
	fs := http.FileServer(http.Dir(path))
	return r.Router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, fs))
}

// HandleEndpoint binds a new Endpoint and Encoder handler to the router
func (r *Router) HandleEndpoint(pattern string, endpoint Endpoint, encoder Encoder) *mux.Route {
	return r.Router.HandleFunc(pattern, endpointWrapper(endpoint, encoder))
}

// HandleTextEndpoint binds a new Text Endpoint handler to the router
func (r *Router) HandleTextEndpoint(pattern string, endpoint Endpoint) *mux.Route {
	return r.Router.HandleFunc(pattern, endpointWrapper(endpoint, encodeText))
}

// HandleJSONEndpoint binds a new JSON Endpoint handler to the router
func (r *Router) HandleJSONEndpoint(pattern string, endpoint Endpoint) *mux.Route {
	return r.Router.HandleFunc(pattern, endpointWrapper(endpoint, encodeJSON))
}

// HandleXMLEndpoint binds a new XML Endpoint handler to the router
func (r *Router) HandleXMLEndpoint(pattern string, endpoint Endpoint) *mux.Route {
	return r.Router.HandleFunc(pattern, endpointWrapper(endpoint, encodeXML))
}

// ListenAndServe assigns the router and address to the routers
// http.Server and begins listening for requests
func (r *Router) ListenAndServe(port int) error {
	// Assign the router to the server handler
	r.Server.Handler = r.Router
	r.Server.Addr = ":" + strconv.Itoa(port)

	// Begin listening for requests
	return r.Server.ListenAndServe()
}

// ListenAndServeTLS assigns the router and address to the routers
// http.Server and begins listening for TLS requests using the
// passed cert and key files
func (r *Router) ListenAndServeTLS(port int, certFile, keyFile string) error {
	// Assign the router to the server handler
	r.Server.Handler = r.Router
	r.Server.Addr = ":" + strconv.Itoa(port)

	// Begin listening for TLS requests
	return r.Server.ListenAndServeTLS(certFile, keyFile)
}

// ListenAndServeFreeTLS will attempt to configure and retrieve
// free SSL certificates from LetsEncrypt using the passed domain
// name and certificate cache directory. It will then listen for
// requests on the standard 443/80 ports.
func (r *Router) ListenAndServeFreeTLS(domain, certDir string) error {
	// Assign the router to the server handler
	r.Server.Handler = r.Router
	r.Server.Addr = ":https"

	// Set up a cert manager that will retrieve configured SSL certs
	// for your domain
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain), // Your domain here
		Cache:      autocert.DirCache("certs"),     // Folder for storing certificates
	}

	// Set the GetCertificate func on the TLS config
	r.Server.TLSConfig = &tls.Config{
		GetCertificate: certManager.GetCertificate,
	}

	// HTTP needed for LetsEncrypt security issue. HTTP should however
	// redirect to HTTPS
	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	// Begin listening for TLS requests
	return r.Server.ListenAndServeTLS("", "")
}
