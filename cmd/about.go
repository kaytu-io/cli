package cmd

import (
	"errors"
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/auth0"
	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use: "about",
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := pkg.GetConfig(cmd, false)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[about]: %v", err)
		}

		bodyResponse, err := auth0.RequestAbout(cnf.AccessToken)
		if err != nil {
			return fmt.Errorf("[about]: %v", err)
		}

		err = pkg.PrintOutput(cmd, "about", bodyResponse)
		if err != nil {
			return fmt.Errorf("[about]: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
