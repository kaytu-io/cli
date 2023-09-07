package main

import (
	_ "embed"
	"fmt"
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
	err = filepath.Walk(root+"/cmd/gen", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "manual_commands" {
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), "manual") {
			return nil
		}
		fmt.Println("PATHHH:", path)
		if path != "./cmd/gen" {
			return os.RemoveAll(path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	crudGen := crud.Generator{}
	var miss bool
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
			if strings.Contains(err.Error(), "url name not found for cmd:") {
				miss = true
				fmt.Println(err.Error())
				return nil
			}
			return err
		}

		err = serviceGen.Generate(root)
		if err != nil {
			return err
		}

		for _, service := range serviceGen.PreferredServices {
			has := false
			for _, s := range crudGen.Services {
				if s == service.NameCamel {
					has = true
				}
			}
			if !has {
				crudGen.Services = append(crudGen.Services, service.NameCamel)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	if miss {
		panic("miss some services")
	}

	err = crudGen.Generate(root)
	if err != nil {
		panic(err)
	}
}
