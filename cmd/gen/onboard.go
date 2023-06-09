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
		GetOnboardCmd.AddCommand(GetOnboardApiV1ProvidersTypesCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CatalogConnectorsCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectorsCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ProvidersCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdDisableCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdDisableCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourcesCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CatalogMetricsCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectorsConnectorNameCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdHealthcheckCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourcesCountCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourcesCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdEnableCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceAwsCmd)

		GetOnboardCmd.AddCommand(PutOnboardApiV1CredentialCredentialIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectionsCountCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCredentialsCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdAutoonboardCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceAzureCmd)

		GetOnboardCmd.AddCommand(DeleteOnboardApiV1SourceSourceIdCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdEnableCmd)

		GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdHealthcheckCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1SourceAccountAccountIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialSourcesListCmd)

		GetOnboardCmd.AddCommand(PutOnboardApiV1SourceSourceIdCredentialsCmd)

		GetOnboardCmd.AddCommand(DeleteOnboardApiV1CredentialCredentialIdCmd)

		GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCmd)
}