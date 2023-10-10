#!/bin/bash

set -e
set -x

#download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
#  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
#curl -o /usr/local/bin/swagger -L'#' "$download_url"
#chmod +x /usr/local/bin/swagger

rm -rf pkg/api/kaytu/client pkg/api/kaytu/models

swagger generate client --spec https://app.kaytu.dev/docs/api/1.0/swagger.yaml --target ../pkg/api/kaytu/
go run scripts/fix_api.go
go run scripts/params/populate_parameters.go > ./gen/config/parameters.go

gofmt -s -w ./gen/config/parameters.go
goimports -w ./gen/config/parameters.go

go run gen/*.go

find ./cmd/gen -name "*.go"  | xargs -I{} gofmt -s -w {}
find ./cmd/gen -name "*.go"  | xargs -I{} goimports -w {}