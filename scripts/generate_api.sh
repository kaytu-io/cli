#!/bin/bash

set -e
set -x

#download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
#  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
#curl -o /usr/local/bin/swagger -L'#' "$download_url"
#chmod +x /usr/local/bin/swagger

rm -rf pkg/api/kaytu/client pkg/api/kaytu/models

swagger generate client --spec /home/saleh/projects/keibi/kaytu-engine/pkg/docs/swagger.yaml --target ./pkg/api/kaytu/
#swagger generate client --spec https://app.kaytu.io/docs/api/1.0/swagger.yaml --target ./pkg/api/kaytu/

go run scripts/fix_api.go
go run codegen/*.go

find ./cmd/genv2 -name "*.go"  | xargs -I{} gofmt -s -w {}
find ./cmd/genv2 -name "*.go"  | xargs -I{} goimports -w {}
