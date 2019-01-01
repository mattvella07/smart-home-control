package middleware

import (
	"fmt"
	"net/http"
)

// Method contains the allowed http request methods for an endpoint
type Method struct {
	Allowed []string
}

// MethodChecker validates that the http request method is allowed for the endpoint
func (m Method) MethodChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		for _, a := range m.Allowed {
			if a == r.Method {
				next.ServeHTTP(rw, r)
				return
			}
		}

		// HTTP request method is not allowed for the endpoint
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(fmt.Sprintf("HTTP request method must be one of the following %s for %s", m.Allowed, r.URL.String())))
	})
}
