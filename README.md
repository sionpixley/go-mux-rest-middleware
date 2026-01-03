# go-mux-rest-middleware

gmrm is a Go library that provides middleware that adds HTTP response headers for REST APIs. This library follows [OWASP REST security guidelines](https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html#security-headers). It only provides middleware for REST APIs that **do not** return any HTML. If your API returns HTML, please implement your own middleware based on OWASP's guidelines.

## Table of contents

1. [How to install](#how-to-install)
2. [How to use](#how-to-use)
3. [Contributing](#contributing)

## How to install

`go get github.com/sionpixley/go-mux-rest-middleware`

## How to use

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    
    "github.com/sionpixley/go-mux-rest-middleware/pkg/gmrm"
)

func main() {
    router := http.NewServeMux()
    router.HandleFunc("GET /api/example", getExample)

    go log.Fatalln(http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps)))
    log.Fatalln(http.ListenAndServeTLS(":443", "certfile", "keyfile", useMiddleware(router)))
}

func getExample(w http.ResponseWriter, r *http.Request) {
    err := json.NewEncoder(w).Encode("example")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://example.com"+r.RequestURI, http.StatusMovedPermanently)
}

func useMiddleware(h http.Handler) http.Handler {
    h = gmrm.CORSMiddleware(h, "https://example", "DELETE, GET, OPTIONS, PATCH, POST", "*")
    h = gmrm.CacheControlMiddleware(h)
    h = gmrm.ContentTypeMiddleware(h, "application/json")
    h = gmrm.FrameMiddleware(h)
    // Only add this one if you want HSTS.
    h = gmrm.HSTSMiddleware(h, "max-age=63072000; includeSubDomains; preload")
    return h
}
```

## Contributing

All contributions are welcome! If you wish to contribute to the project, the best way would be forking this repo and making a pull request from your fork with all of your suggested changes.
