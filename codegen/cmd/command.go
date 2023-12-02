package cmd

import (
	"github.com/kaytu-io/cli-program/codegen/api"
	"strings"
)

type Command struct {
	UsePath          []string
	ShortDescription string
	LongDescription  string
	API              api.API
}

func FromAPI(api api.API, cmdUse string) Command {
	return Command{
		UsePath:          strings.Split(cmdUse, " "),
		ShortDescription: api.ShortDescription,
		LongDescription:  api.LongDescription,
		API:              api,
	}
}
