# Anakin Logger Middleware (anakin-logger-mw)

Anakin Logger Middleware (**anakin-logger-mw**) is a flexible logging middleware for Go that works seamlessly with multiple frameworks including **HTTP**, **Fiber**, **Echo**, and **gRPC**. It provides a simple way to log request and response information, helping developers with debugging and monitoring.

## Purpose

The purpose of this middleware is to:
- Log incoming HTTP/gRPC requests, including method and URL/path.
- Log response status and time taken for the request to complete.
- Support multiple frameworks such as HTTP, Fiber, Echo, and gRPC.
  
This library helps simplify the process of logging and provides an easy-to-integrate solution for capturing key information about requests and responses.

## Features

- Log requests (method, URL, headers).
- Log responses (status code, time taken).
- Support for multiple frameworks:
  - **Standard HTTP** (net/http)
  - **Fiber** (gofiber/fiber)
  - **Echo** (labstack/echo)
  - **gRPC** (google.golang.org/grpc)

## Installation

To install the package, run:

```bash
go get github.com/buzz-jedsada/anakin-logger-mw
```

## Usage
HTTP (Standard net/http)
To use the logging middleware in a standard Go HTTP server:
```go
package main

import (
    "github.com/buzz-jedsada/anakin-logger-mw"
    "net/http"
    "fmt"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // Use the HTTP logging middleware
    loggedMux := middleware.HTTPLogger(mux)

    http.ListenAndServe(":8080", loggedMux)
}
```

Fiber
For Fiber applications, the middleware can be used as follows:
```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/buzz-jedsada/anakin-logger-mw"
)

func main() {
    app := fiber.New()

    // Use the Fiber logging middleware
    app.Use(middleware.FiberLogger)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":8080")
}
```

Echo
For Echo applications, the setup is quite simple:
```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/buzz-jedsada/anakin-logger-mw"
)

func main() {
    e := echo.New()

    // Use the Echo logging middleware
    e.Use(middleware.EchoLogger)

    e.GET("/", func(c echo.Context) error {
        return c.String(200, "Hello, World!")
    })

    e.Start(":8080")
}
```

gRPC
To use the logging middleware with a gRPC service:
```go
package main

import (
    "google.golang.org/grpc"
    "github.com/buzz-jedsada/anakin-logger-mw"
)

func main() {
    // Create a new gRPC server and attach the logger middleware
    server := grpc.NewServer(
        grpc.UnaryInterceptor(middleware.GRPCLogger),
    )

    // Register your gRPC services here
    // ...

    server.Serve(":8080")
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.