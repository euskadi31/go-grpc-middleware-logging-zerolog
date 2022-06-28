# Zerolog logger for grpc middleware [![Last release](https://img.shields.io/github/release/euskadi31/go-grpc-middleware-logging-zerolog.svg)](https://github.com/euskadi31/go-grpc-middleware-logging-zerolog/releases/latest) [![Documentation](https://godoc.org/github.com/euskadi31/go-grpc-middleware-logging-zerolog?status.svg)](https://godoc.org/github.com/euskadi31/go-grpc-middleware-logging-zerolog)

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-grpc-middleware-logging-zerolog)](https://goreportcard.com/report/github.com/euskadi31/go-grpc-middleware-logging-zerolog)

| Branch | Status                                                                                                                                                                                                | Coverage                                                                                                                                                                                       |
| ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| main   | [![Go](https://github.com/euskadi31/go-grpc-middleware-logging-zerolog/actions/workflows/go.yml/badge.svg)](https://github.com/euskadi31/go-grpc-middleware-logging-zerolog/actions/workflows/go.yml) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-grpc-middleware-logging-zerolog/master.svg)](https://coveralls.io/github/euskadi31/go-grpc-middleware-logging-zerolog?branch=main) |

## Example

```go
package main

import (
    grpczerolog "github.com/euskadi31/go-grpc-middleware-logging-zerolog"
    "github.com/rs/zerolog"
    grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
    "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func main() {
    logger := zerolog.New()

    opts := []grpc.ServerOption{
        grpcmiddleware.WithUnaryServerChain(
			logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger, grpczerolog.WithFieldPrefix("foo"))),
        ),
        grpcmiddleware.WithStreamServerChain(
            logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger), grpczerolog.WithFieldPrefix("foo")),
        ),
    }

    //...
}
```
