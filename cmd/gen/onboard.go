package gen

import (
	"github.com/spf13/cobra"
)

var GetOnboardCmd = &cobra.Command{
	Use: "onboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateOnboardCmd = &cobra.Command{
	Use: "onboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteOnboardCmd = &cobra.Command{
	Use: "onboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateOnboardCmd = &cobra.Command{
	Use: "onboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetOnboardCmd.AddCommand(DeleteOnboardApiV1SourceSourceIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCmd)

		GetOnboardCmd.AddCommand(PutOnboardApiV1CredentialCredentialIdCmd)

		GetOnboardCmd.AddCommand(PutOnboardApiV1SourceSourceIdCredentialsCmd)

		GetOnboardCmd.AddCommand(DeleteOnboardApiV1CredentialCredentialIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectionsCountCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CatalogMetricsCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourceAccountAccountIdCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdDisableCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceAwsCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdDisableCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdHealthcheckCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdEnableCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ProvidersCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourcesCountCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceAzureCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdHealthcheckCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourcesCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectorsCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdAutoonboardCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectorsConnectorNameCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdEnableCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourcesCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialSourcesListCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ProvidersTypesCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCredentialsCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CatalogConnectorsCmd)
}