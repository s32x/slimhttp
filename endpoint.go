package slimhttp

import "net/http"

// genericErrorMsg is a generic error message that will be returned if no other
// Error message is defined
const genericErrorMsg = "An error has occurred"

// Endpoint is a service endpoint that receives a request and returns either
// a successfully processed response body or an Error. In either case both
// responses are encoded and returned to the requestor with an appropriately
// defined status code
type Endpoint func(*http.Request) (interface{}, error)

// endpointWrapper transforms an Endpoint into a standard http.Handlerfunc
func endpointWrapper(e Endpoint, n Encoder) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Handle the request and respond appropriately
		res, err := e(req)
		if err != nil {
			if e, ok := err.(*Error); ok {
				n(rw, e.StatusCode, e)
			} else {
				n(rw, http.StatusInternalServerError,
					NewError(genericErrorMsg, http.StatusInternalServerError, err))
			}
			return
		}
		n(rw, http.StatusOK, res)
	}
}
