package restpc

import "net/http"

type Response struct {
	Data   map[string]interface{}
	Header http.Header
}
