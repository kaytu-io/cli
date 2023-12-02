/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/kaytu-io/cli-program/cmd/genv2"
	"github.com/kaytu-io/cli-program/cmd/predef"
	"os"
)

func main() {
	predef.Init()

	err := genv2.KaytuCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
