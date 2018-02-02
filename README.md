# slimhttp 

[![CircleCI](https://circleci.com/gh/sdwolfe32/slimhttp.svg?style=svg)](https://circleci.com/gh/sdwolfe32/slimhttp)
[![GoDoc](https://godoc.org/github.com/sdwolfe32/slimhttp?status.svg)](https://godoc.org/github.com/sdwolfe32/slimhttp)

slimhttp is a simple API library used for writing JSON/XML services quickly and easily. It was written with the aim of providing a go-kit like service definition (slimhttp.Endpoint) while avoiding all the extra RPC logic and encoder/decoder interfaces. The purpose of this project is to implement of a lot of the basic boilerplate associated with writing API services so that you can focus on writing business logic.

## A Note

This is only the beginning for this project. I understand it is extremely basic and that is sort of the point. If there's something you'd really like to see implemented which you use normally in your API logic, I'd love to hear about it! - Post an issue or a pull-request.

## Usage Example

```
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdwolfe32/slimhttp"
)

// Output is an example output struct that will be
// encoded to JSON on the response
type Output struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func main() {
	// Create a new router
	r := slimhttp.NewJSONRouter()

	// Bind an Endpoint to the router at the specified path
	r.HandleEndpoint("/hello/{name}/", Hello)

	// Start the service!
	r.ListenAndServe("8080")
}

// Hello is an example Endpoint method. It receives a
// request so that you have access to everything on the
// request and returns a successful body or error
func Hello(r *http.Request) (interface{}, error) {
	name := mux.Vars(r)["name"] // The name passed on the request

	// "fancy-error" as the name invokes and returns a fully
	// encoded slimhttp.Error which is created here
	if name == "fancy-error" {
		err := errors.New("This is a fancy error")
		return nil, slimhttp.NewError("There's a very bad error!", http.StatusBadRequest, err)
	}

	// "basic-error" as the name invokes and returns a fully
	// encoded slimhttp.Error that is generated in the wrapper
	if name == "basic-error" {
		err := errors.New("This is a basic error")
		return nil, err
	}

	// All other names will pass through and return a fully
	// encoded Output
	return &Output{
		Message: fmt.Sprintf("Hello %s!", name),
		Success: true,
	}, nil
}

```

The MIT License (MIT)
=====================

Copyright © 2018 Steven Wolfe

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