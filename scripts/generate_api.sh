#!/bin/bash

set -e

swagger generate client --spec ../keibi-engine/pkg/docs/swagger.yaml --target ./pkg/api/kaytu/
go run scripts/fix_api.go
go run scripts/params/populate_parameters.go > ./gen/parameters.go

gofmt -s -w ./gen/parameters.go
goimports -w ./gen/parameters.go

go run gen/*.go

find ./cmd/gen -name "*.go"  | xargs -I{} gofmt -s -w {}
find ./cmd/gen -name "*.go"  | xargs -I{} goimports -w {}