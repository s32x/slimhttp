package slimhttp

import (
	"net/http"
)

// genericErrorMsg is a generic error message that will be returned if no
// other Error message is defined
const genericErrorMsg = "An error has occurred"

// Endpoint is a service endpoint that receives a request and returns either
// a successfully processed response body or an Error. In either case both
// responses are encoded and returned to the requestor with an appropriately
// defined status code
type Endpoint func(*http.Request) (interface{}, error)

// endpointWrapper transforms an Endpoint into a standard http.Handlerfunc
func (r *Router) endpointWrapper(e Endpoint, n Encoder, opts ...Option) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		for _, opt := range opts {
			switch opt {
			case OptionURLSigning:
				// Verify URL Signature
				if err := r.URLSigner.ValidateURL(req.URL); err != nil {
					returnError(rw, n, err)
					return
				}
			case OptionCache:
				// If the data request is in our cache return it
				if res, found := r.Cache.Get(req.URL.String()); found {
					n(rw, http.StatusOK, res)
					return
				}
			}
		}

		// Handle the request and respond appropriately
		res, err := e(req)
		if err != nil {
			returnError(rw, n, err)
			return
		}
		n(rw, http.StatusOK, res)

		for _, opt := range opts {
			switch opt {
			case OptionCache:
				// Cache the response
				r.Cache.SetDefault(req.URL.String(), res)
			}
		}
	}
}

// returnError with either attempt to encode our own Error or it will
// generate a new one and write that to the ResponseWriter
func returnError(rw http.ResponseWriter, n Encoder, err error) {
	if e, ok := err.(*Error); ok {
		n(rw, e.StatusCode, e)
	}
	n(rw, http.StatusInternalServerError,
		NewError(genericErrorMsg, http.StatusInternalServerError, err))
}
