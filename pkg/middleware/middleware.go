/*
Package middleware provides functions that add HTTP headers to REST APIs based on [OWASP guidelines].

It only works with the [gorilla/mux] library, and it assumes that your API does not return any HTML.
If your API returns HTML, please implement custom middleware.

[gorilla/mux]: https://pkg.go.dev/github.com/gorilla/mux
[OWASP guidelines]: https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html#security-headers
*/
package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Adds the 'Cache-Control' header with the value 'no-store'.
func SetCacheControlHeader(router *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
		})
	}
}

// Adds the 'Content-Type' header with the value of the 'contentType' argument.
// Also adds the 'X-Content-Type-Options' header with the value 'nosniff'.
func SetContentTypeHeaders(router *mux.Router, contentType string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("X-Content-Type-Options", "nosniff")
			next.ServeHTTP(w, r)
		})
	}
}

// Adds the 'Access-Control-Allow-Origin' header with the value of the 'origin' argument.
func SetCorsOriginHeader(router *mux.Router, origin string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}
}

// Adds the 'Content-Security-Policy' header with the value "frame-ancestors 'none'".
// Also adds the 'X-Frame-Options' with the value 'DENY'.
func SetFrameHeaders(router *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Security-Policy", "frame-ancestors 'none'")
			w.Header().Set("X-Frame-Options", "DENY")
			next.ServeHTTP(w, r)
		})
	}
}

// Adds the 'Strict-Transport-Security' header with the value of the 'value' argument.
func SetHstsHeader(router *mux.Router, value string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Strict-Transport-Security", value)
			next.ServeHTTP(w, r)
		})
	}
}
