package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/keibiengine/keibi-engine/pkg/cli"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into Kaytu",
	Long:  `Logging into Kaytu using device authentication mechanism`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceCode, err := cli.RequestDeviceCode()
		if err != nil {
			return fmt.Errorf("[login] : %v", err)
		}

		accessToken, err := cli.AccessToken(deviceCode)
		if err != nil {
			return fmt.Errorf("[login] : %v", err)
		}

		err = cli.AddConfig(accessToken)
		if err != nil {
			return fmt.Errorf("[login] : %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
