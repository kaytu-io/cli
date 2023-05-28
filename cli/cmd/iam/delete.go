package cmd

import (
	"errors"
	"fmt"
	"github.com/kaytu-io/cli-program/pkg/cli"
	"github.com/spf13/cobra"
)

var Delete = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var deleteCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "it is remove credential",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Lookup("id").Changed {
		} else {
			return errors.New("please enter the id flag")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := cli.GetConfig(cmd, true)
		if err != nil {
			return err
		}
		err = cli.OnboardDeleteCredential(cnf.DefaultWorkspace, cnf.AccessToken, credentialIdGet)
		if err != nil {
			return err
		}
		fmt.Println("removed credential successfully.")
		return nil
	},
}

var deleteSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "it will delete source ",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Lookup("id").Changed {
		} else {
			return errors.New("please enter the id flag")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := cli.GetConfig(cmd, true)
		if err != nil {
			return err
		}
		err = cli.OnboardDeleteSource(cnf.DefaultWorkspace, cnf.AccessToken, sourceIdDelete)
		if err != nil {
			return err
		}
		fmt.Println("removed source successfully")
		return nil
	},
}

var sourceIdDelete string
var credentialIdDelete string
var workspaceNameDelete string

func init() {
	Delete.AddCommand(IamDelete)
	Delete.AddCommand(deleteCredentialCmd)
	Delete.AddCommand(deleteSourceCmd)

	//delete source flag :
	deleteSourceCmd.Flags().StringVar(&sourceIdDelete, "id", "", "it is specifying the source id. ")
	deleteSourceCmd.Flags().StringVar(&workspaceNameDelete, "workspace-name", "", "it is specifying the workspace name[mandatory].")

	//delete credential :
	deleteCredentialCmd.Flags().StringVar(&credentialIdDelete, "id", "", "it is specifying the credentialIdGet[mandatory].")
	deleteCredentialCmd.Flags().StringVar(&workspaceNameDelete, "workspace-name", "", "it is specifying the workspace name[mandatory].")

}
