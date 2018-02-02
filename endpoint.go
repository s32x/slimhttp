package slimhttp

import "net/http"

// ErrorStandard is a basic error message that will be returned if no other
// Error message is defined
const ErrorStandard = "An error has occurred"

// Endpoint is a service endpoint that receives a request and returns either
// a successfully processed response body or an Error. In either case both
// responses are encoded and returned to the requestor with an appropriately
// defined status code
type Endpoint func(*http.Request) (interface{}, error)

// endpointWrapper transforms an Endpoint into a standard http.Handlerfunc
func (r *router) endpointWrapper(e Endpoint) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Handle the request and respond appropriately
		res, err := e(req)
		if err != nil {
			if e, ok := err.(*Error); ok {
				r.encode(rw, e.StatusCode, e)
			} else {
				r.encode(rw, http.StatusInternalServerError,
					NewError(ErrorStandard, http.StatusInternalServerError, err))
			}
			return
		}
		r.encode(rw, http.StatusOK, res)
	}
}
