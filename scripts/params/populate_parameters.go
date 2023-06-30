package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	keyNames := map[string]interface{}{}
	fmt.Println("package config")
	fmt.Println("var ParamModels = map[string]interface{}{")
	err := filepath.Walk("./pkg/api/kaytu/client", func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, "_parameters.go") {
			return nil
		}
		idx := strings.LastIndex(path, "/")
		tmp := path[:idx]
		idx = strings.LastIndex(tmp, "/")
		serviceName := tmp[idx+1:]

		c, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		lines := strings.Split(string(c), "\n")
		for _, line := range lines {
			if strings.Contains(line, "struct ") {
				arr := strings.Split(strings.TrimSpace(line), " ")
				name := arr[1]
				keyName := name[:len(name)-6]

				if _, ok := keyNames[keyName]; ok {
					continue
				}
				keyNames[keyName] = struct{}{}
				fmt.Printf("\"%s\": \t%s.%s{},\n", keyName, serviceName, name)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("}")
}
