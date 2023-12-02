package main

import (
	_ "embed"
	"github.com/kaytu-io/cli-program/codegen/api"
	"github.com/kaytu-io/cli-program/codegen/cmd"
	"github.com/kaytu-io/cli-program/codegen/cobra"
	"github.com/kaytu-io/cli-program/codegen/config"
	"github.com/kaytu-io/cli-program/codegen/swagger"
	"os"
)

func cleanup(root string) error {
	return os.RemoveAll(root + "/cmd/genv2")
}

var swaggerFileName = "https://app.kaytu.io/docs/api/1.0/swagger.yaml"

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

	err = cleanup(root)
	if err != nil {
		panic(err)
	}

	var commands []cmd.Command
	apis := api.ExtractAPIs(root)
	for _, item := range apis {
		cmdUse := config.GetCommandDef(item)
		if cmdUse == "" {
			continue
		}

		cmdItem := cmd.FromAPI(item, cmdUse)
		commands = append(commands, cmdItem)
	}

	rootCmd := cobra.BuildCobraCommands(commands, swag)
	cobra.GenerateCobraCommands(rootCmd, root+"/cmd/genv2")
}
