# go-mux-rest-middleware

This is a Go library that provides middleware that adds HTTP headers for REST APIs when using [gorilla/mux](https://github.com/gorilla/mux). This library follows [OWASP REST security guidelines](https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html#security-headers). It only provides middleware for REST APIs that **do not** return any HTML. If your API returns HTML, please implement your own middleware based on OWASP's guidelines.

## Table of contents

1. [Project directory structure](#project-directory-structure)
2. [How to install](#how-to-install)
3. [How to use](#how-to-use)
4. [Contributing](#contributing)

## Project directory structure

```
go-mux-rest-middleware
├── LICENSE
├── README.md
├── go.mod
├── go.sum
└── pkg
    └── middleware
        └── middleware.go
```

## How to install

`go get github.com/sionpixley/go-mux-rest-middleware`

## How to use

```
package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/sionpixley/go-mux-rest-middleware/pkg/middleware"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/example", getExample).Methods(http.MethodGet, http.MethodOptions)

    router.Use(mux.CORSMethodMiddleware(router))

    // These are the functions provided by the middleware package.
    router.Use(middleware.SetCorsOriginHeader(router, "https://example.com"))
    router.Use(middleware.SetCacheControlHeader(router))
    router.Use(middleware.SetContentTypeHeaders(router, "application/json"))
    router.Use(middleware.SetFrameHeaders(router))
    // Only add this one if you want HSTS.
    router.Use(middleware.SetHstsHeader(router, "max-age=63072000; includeSubDomains; preload"))

    go http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
    log.Fatal(http.ListenAndServeTLS(":443", "certfile", "keyfile", router))
}

func getExample(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        return
    }

    err := json.NewEncoder(w).Encode("example")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://example.com"+r.RequestURI, http.StatusMovedPermanently)
}
```

## Contributing

All contributions are welcome! If you wish to contribute to the project, the best way would be forking this repo and making a pull request from your fork with all of your suggested changes.