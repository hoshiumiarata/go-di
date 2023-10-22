# di

This package is a runtime dependency injection framework. It is PoC for now.

It has notable overhead, so it should be used with caution.

Check out [documentation](https://pkg.go.dev/github.com/hoshiumiarata/go-di) for more information.

## Installation

```bash
go get github.com/hoshiumiarata/go-di
```

## Example

This example shows how to use this package with `net/http` package.

```go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    "github.com/hoshiumiarata/go-di"
)

// Example of a handler. Arguments can be in any order.
func hello(w http.ResponseWriter, req *http.Request, sql *sql.DB, logger *log.Logger, container *di.Container) {
    logger.Println("Hello!")
    logger.Printf("SQL: %#v\n", sql)
    logger.Printf("Container: %#v\n", container)
}

// Example of a handler. Arguments can be in any order.
func world(w http.ResponseWriter, req *http.Request, container *di.Container, logger *log.Logger, sql *sql.DB) {
    logger.Println("World!")
    logger.Printf("SQL: %#v\n", sql)
    logger.Printf("Container: %#v\n", container)
}

// Wrap a handler to inject values.
func wrapHandler(container *di.Container, f any) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        // Create a child container.
        child := container.New()
        // Register values to the child container.
        di.RegisterByPositionTo(child, 0, w)
        di.RegisterByPositionTo(child, 1, req)
        // Call the handler.
        child.Call(f)
    }
}

func main() {
    // Create values.
    sql := &sql.DB{}
    logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

    // Create a container.
    container := di.New()
    // Register values to the container.
    di.RegisterValueTo(container, sql)
    di.RegisterValueTo(container, logger)

    // Register handlers.
    http.HandleFunc("/hello", wrapHandler(container, hello))
    http.HandleFunc("/world", wrapHandler(container, world))

    http.ListenAndServe(":8090", nil)
}
```
