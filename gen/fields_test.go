package main

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/compliance"
	"reflect"
	"testing"
)

func TestExtractFields(t *testing.T) {
	//&models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest
	f := ExtractFields(reflect.TypeOf(compliance.PostComplianceAPIV1FindingsParams{}))
	fmt.Println(f)
}
