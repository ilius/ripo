package ripo

import "net/http"

type Response struct {
	// Data: map or struct with json tags
	Data any

	Header http.Header

	RedirectPath       string
	RedirectStatusCode int
}
