#!/bin/bash

set -e

swagger generate client --spec ../keibi-engine/pkg/docs/swagger.yaml --target ./pkg/api/kaytu/
go run scripts/fix_api.go