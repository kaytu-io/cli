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
	GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectorsConnectorNameCmd)
	GetOnboardApiV1ConnectorsConnectorNameCmd.Flags().String("connector-name", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1SourcesCountCmd)
	GetOnboardApiV1SourcesCountCmd.Flags().String("connector", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1SourceAzureCmd)
	PostOnboardApiV1SourceAzureCmd.Flags().String("client-id", "", "")
	PostOnboardApiV1SourceAzureCmd.Flags().String("client-secret", "", "")
	PostOnboardApiV1SourceAzureCmd.Flags().String("object-id", "", "")
	PostOnboardApiV1SourceAzureCmd.Flags().String("secret-id", "", "")
	PostOnboardApiV1SourceAzureCmd.Flags().String("subscription-id", "", "")
	PostOnboardApiV1SourceAzureCmd.Flags().String("tenant-id", "", "")

	PostOnboardApiV1SourceAzureCmd.Flags().String("description", "", "")
	PostOnboardApiV1SourceAzureCmd.Flags().String("name", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1SourcesCmd)
	PostOnboardApiV1SourcesCmd.Flags().StringArray("source-ids", nil, "")

	PostOnboardApiV1SourcesCmd.Flags().String("type", "", "")

	GetOnboardCmd.AddCommand(DeleteOnboardApiV1CredentialCredentialIdCmd)
	DeleteOnboardApiV1CredentialCredentialIdCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCredentialsCmd)
	GetOnboardApiV1SourceSourceIdCredentialsCmd.Flags().String("source-id", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdDisableCmd)
	PostOnboardApiV1CredentialCredentialIdDisableCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1SourceAwsCmd)
	PostOnboardApiV1SourceAwsCmd.Flags().String("access-key", "", "")
	PostOnboardApiV1SourceAwsCmd.Flags().String("account-id", "", "")
	PostOnboardApiV1SourceAwsCmd.Flags().StringArray("regions", nil, "")
	PostOnboardApiV1SourceAwsCmd.Flags().String("secret-key", "", "")

	PostOnboardApiV1SourceAwsCmd.Flags().String("description", "", "")
	PostOnboardApiV1SourceAwsCmd.Flags().String("email", "", "")
	PostOnboardApiV1SourceAwsCmd.Flags().String("name", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1CatalogConnectorsCmd)
	GetOnboardApiV1CatalogConnectorsCmd.Flags().String("category", "", "")
	GetOnboardApiV1CatalogConnectorsCmd.Flags().String("id", "", "")
	GetOnboardApiV1CatalogConnectorsCmd.Flags().String("min-connection", "", "")
	GetOnboardApiV1CatalogConnectorsCmd.Flags().String("state", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialSourcesListCmd)
	GetOnboardApiV1CredentialSourcesListCmd.Flags().String("connector", "", "")
	GetOnboardApiV1CredentialSourcesListCmd.Flags().Int64("page-number", 0, "")
	GetOnboardApiV1CredentialSourcesListCmd.Flags().Int64("page-size", 0, "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1SourcesCmd)
	GetOnboardApiV1SourcesCmd.Flags().String("connector", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCmd)
	PostOnboardApiV1CredentialCmd.Flags().String("config", "", "")
	PostOnboardApiV1CredentialCmd.Flags().String("name", "", "")
	PostOnboardApiV1CredentialCmd.Flags().String("source-type", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectionsCountCmd)
	GetOnboardApiV1ConnectionsCountCmd.Flags().StringArray("connectors", nil, "")
	GetOnboardApiV1ConnectionsCountCmd.Flags().String("health", "", "")
	GetOnboardApiV1ConnectionsCountCmd.Flags().String("state", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdHealthcheckCmd)
	GetOnboardApiV1CredentialCredentialIdHealthcheckCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1ProvidersCmd)

	GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdDisableCmd)
	PostOnboardApiV1SourceSourceIdDisableCmd.Flags().Int64("source-id", 0, "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdEnableCmd)
	PostOnboardApiV1SourceSourceIdEnableCmd.Flags().Int64("source-id", 0, "")

	GetOnboardCmd.AddCommand(PutOnboardApiV1CredentialCredentialIdCmd)
	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("config", "", "")
	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("connector", "", "")
	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("name", "", "")

	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(PutOnboardApiV1SourceSourceIdCredentialsCmd)
	PutOnboardApiV1SourceSourceIdCredentialsCmd.Flags().String("source-id", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1SourceAccountAccountIdCmd)
	GetOnboardApiV1SourceAccountAccountIdCmd.Flags().Int64("account-id", 0, "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1ConnectionsConnectionIdStateCmd)
	PostOnboardApiV1ConnectionsConnectionIdStateCmd.Flags().Int64("connection-id", 0, "")
	PostOnboardApiV1ConnectionsConnectionIdStateCmd.Flags().String("state", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdAutoonboardCmd)
	PostOnboardApiV1CredentialCredentialIdAutoonboardCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdEnableCmd)
	PostOnboardApiV1CredentialCredentialIdEnableCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1ConnectorsCmd)

	GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdCmd)
	GetOnboardApiV1CredentialCredentialIdCmd.Flags().String("credential-id", "", "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1CredentialCmd)
	GetOnboardApiV1CredentialCmd.Flags().String("connector", "", "")
	GetOnboardApiV1CredentialCmd.Flags().String("health", "", "")
	GetOnboardApiV1CredentialCmd.Flags().Int64("page-number", 0, "")
	GetOnboardApiV1CredentialCmd.Flags().Int64("page-size", 0, "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1CatalogMetricsCmd)

	GetOnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdHealthcheckCmd)
	PostOnboardApiV1SourceSourceIdHealthcheckCmd.Flags().String("source-id", "", "")

	GetOnboardCmd.AddCommand(DeleteOnboardApiV1SourceSourceIdCmd)
	DeleteOnboardApiV1SourceSourceIdCmd.Flags().Int64("source-id", 0, "")

	GetOnboardCmd.AddCommand(GetOnboardApiV1ProvidersTypesCmd)

	GetOnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCmd)
	GetOnboardApiV1SourceSourceIdCmd.Flags().Int64("source-id", 0, "")
}
