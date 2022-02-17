# slimhttp

slimhttp is a simple API library used for writing JSON/XML services quickly and easily. It was written with the aim of providing a go-kit like service definition (slimhttp.Endpoint) while avoiding all the extra RPC logic and encoder/decoder interfaces. The purpose of this project is to implement of a lot of the basic boilerplate associated with writing API services so that you can focus on writing business logic.

## Why?

The heart of slimhttp is the Endpoint type. All your endpoints from now on will take the form of the below function signature.

```
type Endpoint func(*http.Request) (interface{}, error)
```

The use of an http.HandlerFunc was the driving force behind writing this library. Satisfying the http.HandlerFunc type (including all encoding and error checking in the same function) was not something I was a fan of and thus came up with above type to make things a little more straightforward. Using the new Endpoint type above now gives you the ability to offload all encoding and error handling to this library, making the process of implementing business logic a little cleaner.

## A Note

This is only the beginning for this project. I understand it is extremely basic and that is sort of the point. If there's something you'd really like to see implemented which you use normally in your API logic, I'd love to hear about it! - Post an issue or a pull-request.

## Usage Example

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s32x/slimhttp"
)

func main() {
	r := slimhttp.NewRouter()                     // Create a new router
	r.HandleJSONEndpoint("/hello/{name}/", Hello) // Bind an Endpoint to the router at the specified path
	log.Fatal(r.ListenAndServe(8080))             // Start the service!
}

// HelloResponse is an example response struct that will be
// encoded to JSON on a Hello request
type HelloResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// Hello is an example Endpoint method. It receives a request
// so that you have access to everything on the request and
// returns a successful body or error
func Hello(r *http.Request) (interface{}, error) {
	name := mux.Vars(r)["name"] // The name passed on the request

	switch name {
	case "basic-error":
		// An example of returning a raw error
		err := errors.New("This is a basic error")
		return nil, err
	case "standard-error":
		// An example of returning a predefined Error
		return nil, slimhttp.ErrorBadRequest
	case "fancy-error":
		// An example of returning a fully self-defined Error
		err := errors.New("This is a fancy error")
		return nil, slimhttp.NewError("This is a fancy error!", http.StatusBadRequest, err)
	}

	// All other names will be returned on a HelloResponse
	return &HelloResponse{
		Message: fmt.Sprintf("Hello %s!", name),
		Success: true,
	}, nil
}

```

The MIT License (MIT)
=====================

Copyright © 2022 s32x

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the “Software”), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
