package restpc

import "net/http"

type Response struct {
	// Data: map or struct with json tags
	Data interface{}

	Header http.Header

	RedirectPath       string
	RedirectStatusCode int
}
