package slimhttp

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	cache "github.com/patrickmn/go-cache"
	"golang.org/x/crypto/acme/autocert"
)

// Router contains an mux.Router that we can use to bind Endpoints
type Router struct {
	Router    *mux.Router
	Server    *http.Server
	Cache     *cache.Cache
	URLSigner *URLSigner
}

// NewDefaultRouter generates a basic default router with some
// standard parameters
func NewDefaultRouter() *Router {
	r := NewBaseRouter()
	r.SetTimeout(30*time.Second, 30*time.Second)
	r.SetCache(5*time.Minute, 10*time.Minute)
	r.SetURLSigner("")
	return r
}

// NewBaseRouter generates a new Router instance with a basic Router
func NewBaseRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
		Server: &http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

// SetTimeout sets a custom read and write timeout on the routers
// http.Server
func (r *Router) SetTimeout(read, write time.Duration) {
	r.Server.ReadTimeout = read
	r.Server.WriteTimeout = write
}

// SetCache sets a custom default expiration and cleanup interval
// in a new cache on the Router
func (r *Router) SetCache(defaultExpiration, cleanupInterval time.Duration) {
	r.Cache = cache.New(defaultExpiration, cleanupInterval)
}

// SetURLSigner sets a new custom URL Signer with a custom secret
func (r *Router) SetURLSigner(secret string) {
	if secret == "" {
		r.URLSigner = NewURLSigner()
		return
	}
	r.URLSigner = NewURLSignerFromSecret(secret)
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
