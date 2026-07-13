# cognition-go-sdk

A Go client for the [Devin External API](https://docs.devin.ai), generated from
the official OpenAPI specification using
[`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen).

## Install

```sh
go get github.com/shinji62/cognition-go-sdk/api
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    devin "github.com/shinji62/cognition-go-sdk/api"
)

func main() {
    client, err := devin.NewBearerClient(os.Getenv("DEVIN_API_KEY"))
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.PostV1SessionsWithResponse(context.Background(), devin.PostV1SessionsJSONRequestBody{
        Prompt: "Review the pull request at https://github.com/example/repo/pull/123",
    })
    if err != nil {
        log.Fatal(err)
    }
    if resp.JSON200 == nil {
        log.Fatalf("unexpected status %d: %s", resp.StatusCode(), string(resp.Body))
    }

    fmt.Println("Session:", resp.JSON200.SessionId)
    fmt.Println("URL:", resp.JSON200.Url)
}
```

`NewBearerClient` targets the production server (`https://api.devin.ai`) and adds
the `Authorization: Bearer <token>` header to every request. Use
`NewBearerClientWithServer` to point at a different base URL.

## Project layout

| Path                     | Description                                             |
| ------------------------ | ------------------------------------------------------- |
| `openapi/openapi.yaml`   | Vendored copy of the Devin OpenAPI spec.                |
| `api/config.yaml`        | `oapi-codegen` configuration.                           |
| `api/devin.gen.go`       | Generated client + models (**do not edit by hand**).    |
| `api/client.go`          | Hand-written helpers (bearer auth, default server).     |
| `api/generate.go`        | `go:generate` directive.                                |

## Regenerating the client

The client is regenerated from the vendored spec:

```sh
make generate   # runs `go generate ./...` + `go mod tidy`
```

To pull the latest spec first:

```sh
make fetch-spec generate build
```

### On demand via GitHub Actions

The [Generate SDK workflow](.github/workflows/generate.yml) can be run manually
from the **Actions** tab (**Run workflow**). It:

1. Fetches the latest OpenAPI spec (URL overridable as a workflow input).
2. Regenerates the client and verifies it compiles.
3. Opens a pull request if anything changed.

## Requirements

- Go 1.24.3 or newer (see [go.mod](go.mod)).
