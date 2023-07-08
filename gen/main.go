package main

import (
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/kaytu-io/cli-program/gen/crud"
	"github.com/kaytu-io/cli-program/gen/services"
	"github.com/kaytu-io/cli-program/gen/swagger"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var swaggerFileName = "https://app.kaytu.dev/docs/api/1.0/swagger.yaml"

func main() {
	swag, err := swagger.New(swaggerFileName)
	if err != nil {
		panic(err)
	}

	root := "."
	_, err = os.Stat(root + "/pkg")
	if err != nil {
		root = ".."
	}
	os.RemoveAll(root + "/cmd/gen")
	os.Mkdir(root+"/cmd/gen", os.ModePerm)

	crudGen := crud.Generator{}
	err = filepath.Walk(root+"/pkg/api/kaytu/client", func(path string, info fs.FileInfo, err error) error {
		if info == nil || !info.IsDir() || strings.HasSuffix(path, "/kaytu/client") {
			return nil
		}

		serviceName := info.Name()
		serviceGen := services.Generator{
			ServiceFileName:  serviceName,
			ServiceNameCamel: strcase.ToCamel(serviceName),
			ServiceNameSnake: strcase.ToSnake(serviceName),
			Swagger:          swag,
		}

		err = serviceGen.ExtractChildren(root, path)
		if err != nil {
			return err
		}

		err = serviceGen.Generate(root)
		if err != nil {
			return err
		}

		crudGen.Services = append(crudGen.Services, strcase.ToCamel(info.Name()))
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = crudGen.Generate(root)
	if err != nil {
		panic(err)
	}
}
