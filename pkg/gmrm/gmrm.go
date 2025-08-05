/*
Package gmrm provides functions that add HTTP response headers to REST APIs based on [OWASP guidelines].

It only works with the [gorilla/mux] library, and it assumes that your API does not return any HTML.
If your API returns HTML, please implement custom middleware.

[gorilla/mux]: https://pkg.go.dev/github.com/gorilla/mux
[OWASP guidelines]: https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html#security-headers
*/
package gmrm

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CacheControlMiddleware function adds the 'Cache-Control' response header with the value 'no-store'.
func CacheControlMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
		})
	}
}

// ContentTypeMiddleware function adds the 'Content-Type' response header with the value of the 'contentType' argument.
// Also adds the 'X-Content-Type-Options' response header with the value 'nosniff'.
func ContentTypeMiddleware(contentType string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("X-Content-Type-Options", "nosniff")
			next.ServeHTTP(w, r)
		})
	}
}

// CorsOriginMiddleware function adds the 'Access-Control-Allow-Origin' response header with the value of the 'origin' argument.
func CorsOriginMiddleware(origin string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			next.ServeHTTP(w, r)
		})
	}
}

// FrameMiddleware function adds the 'Content-Security-Policy' response header with the value "frame-ancestors 'none'".
// Also adds the 'X-Frame-Options' response header with the value 'DENY'.
func FrameMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Security-Policy", "frame-ancestors 'none'")
			w.Header().Set("X-Frame-Options", "DENY")
			next.ServeHTTP(w, r)
		})
	}
}

// HstsMiddleware function adds the 'Strict-Transport-Security' response header with the value of the 'value' argument.
func HstsMiddleware(value string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Strict-Transport-Security", value)
			next.ServeHTTP(w, r)
		})
	}
}
