// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var OnboardCmd = &cobra.Command{
	Use: "onboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	OnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdAutoonboardCmd)
	PostOnboardApiV1CredentialCredentialIdAutoonboardCmd.Flags().String("credential-id", "", "")
	// PostOnboardApiV1CredentialCredentialIdAutoonboardCmd.MarkFlagRequired("credential-id")
	PostOnboardApiV1CredentialCredentialIdAutoonboardCmd.MarkFlagRequired("credential-id")

	OnboardCmd.AddCommand(PostOnboardApiV1SourceAwsCmd)
	PostOnboardApiV1SourceAwsCmd.Flags().String("access-key", "", "")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("access-key")
	PostOnboardApiV1SourceAwsCmd.Flags().String("account-id", "", "")
	// PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("account-id")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("account-id")
	PostOnboardApiV1SourceAwsCmd.Flags().String("assume-role-name", "", "")
	// PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("assume-role-name")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("assume-role-name")
	PostOnboardApiV1SourceAwsCmd.Flags().String("external-id", "", "")
	// PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("external-id")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("external-id")
	PostOnboardApiV1SourceAwsCmd.Flags().StringArray("regions", nil, "")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("regions")
	PostOnboardApiV1SourceAwsCmd.Flags().String("secret-key", "", "")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("secret-key")

	PostOnboardApiV1SourceAwsCmd.Flags().String("description", "", "")
	// PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("description")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("description")
	PostOnboardApiV1SourceAwsCmd.Flags().String("email", "", "")
	// PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("email")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("email")
	PostOnboardApiV1SourceAwsCmd.Flags().String("name", "", "")
	// PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("name")
	PostOnboardApiV1SourceAwsCmd.MarkFlagRequired("name")

	OnboardCmd.AddCommand(PostOnboardApiV1SourceAzureCmd)
	PostOnboardApiV1SourceAzureCmd.Flags().String("client-id", "", "")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("client-id")
	PostOnboardApiV1SourceAzureCmd.Flags().String("client-secret", "", "")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("client-secret")
	PostOnboardApiV1SourceAzureCmd.Flags().String("object-id", "", "")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("object-id")
	PostOnboardApiV1SourceAzureCmd.Flags().String("secret-id", "", "")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("secret-id")
	PostOnboardApiV1SourceAzureCmd.Flags().String("subscription-id", "", "")
	// PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("subscription-id")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("subscription-id")
	PostOnboardApiV1SourceAzureCmd.Flags().String("tenant-id", "", "")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("tenant-id")

	PostOnboardApiV1SourceAzureCmd.Flags().String("description", "", "")
	// PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("description")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("description")
	PostOnboardApiV1SourceAzureCmd.Flags().String("name", "", "")
	// PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("name")
	PostOnboardApiV1SourceAzureCmd.MarkFlagRequired("name")

	OnboardCmd.AddCommand(PostOnboardApiV1CredentialCmd)
	PostOnboardApiV1CredentialCmd.Flags().String("config", "", "")
	PostOnboardApiV1CredentialCmd.MarkFlagRequired("config")
	PostOnboardApiV1CredentialCmd.Flags().String("name", "", "")
	// PostOnboardApiV1CredentialCmd.MarkFlagRequired("name")
	PostOnboardApiV1CredentialCmd.MarkFlagRequired("name")
	PostOnboardApiV1CredentialCmd.Flags().String("source-type", "", "")
	PostOnboardApiV1CredentialCmd.MarkFlagRequired("source-type")

	OnboardCmd.AddCommand(DeleteOnboardApiV1CredentialCredentialIdCmd)
	DeleteOnboardApiV1CredentialCredentialIdCmd.Flags().String("credential-id", "", "")
	// DeleteOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("credential-id")
	DeleteOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("credential-id")

	OnboardCmd.AddCommand(DeleteOnboardApiV1SourceSourceIdCmd)
	DeleteOnboardApiV1SourceSourceIdCmd.Flags().Int64("source-id", 0, "")
	// DeleteOnboardApiV1SourceSourceIdCmd.MarkFlagRequired("source-id")
	DeleteOnboardApiV1SourceSourceIdCmd.MarkFlagRequired("source-id")

	OnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdDisableCmd)
	PostOnboardApiV1CredentialCredentialIdDisableCmd.Flags().String("credential-id", "", "")
	// PostOnboardApiV1CredentialCredentialIdDisableCmd.MarkFlagRequired("credential-id")
	PostOnboardApiV1CredentialCredentialIdDisableCmd.MarkFlagRequired("credential-id")

	OnboardCmd.AddCommand(PostOnboardApiV1CredentialCredentialIdEnableCmd)
	PostOnboardApiV1CredentialCredentialIdEnableCmd.Flags().String("credential-id", "", "")
	// PostOnboardApiV1CredentialCredentialIdEnableCmd.MarkFlagRequired("credential-id")
	PostOnboardApiV1CredentialCredentialIdEnableCmd.MarkFlagRequired("credential-id")

	OnboardCmd.AddCommand(GetOnboardApiV1SourceAccountAccountIdCmd)
	GetOnboardApiV1SourceAccountAccountIdCmd.Flags().Int64("account-id", 0, "")
	// GetOnboardApiV1SourceAccountAccountIdCmd.MarkFlagRequired("account-id")
	GetOnboardApiV1SourceAccountAccountIdCmd.MarkFlagRequired("account-id")

	OnboardCmd.AddCommand(GetOnboardApiV1CatalogMetricsCmd)

	OnboardCmd.AddCommand(GetOnboardApiV1ConnectionsCountCmd)
	GetOnboardApiV1ConnectionsCountCmd.Flags().StringArray("connectors", nil, "")
	GetOnboardApiV1ConnectionsCountCmd.MarkFlagRequired("connectors")
	GetOnboardApiV1ConnectionsCountCmd.Flags().String("state", "", "")
	GetOnboardApiV1ConnectionsCountCmd.MarkFlagRequired("state")

	OnboardCmd.AddCommand(GetOnboardApiV1ConnectorConnectorNameCmd)
	GetOnboardApiV1ConnectorConnectorNameCmd.Flags().String("connector-name", "", "")
	// GetOnboardApiV1ConnectorConnectorNameCmd.MarkFlagRequired("connector-name")
	GetOnboardApiV1ConnectorConnectorNameCmd.MarkFlagRequired("connector-name")

	OnboardCmd.AddCommand(GetOnboardApiV1ConnectorCmd)

	OnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdCmd)
	GetOnboardApiV1CredentialCredentialIdCmd.Flags().String("credential-id", "", "")
	// GetOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("credential-id")
	GetOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("credential-id")

	OnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCmd)
	GetOnboardApiV1SourceSourceIdCmd.Flags().Int64("source-id", 0, "")
	// GetOnboardApiV1SourceSourceIdCmd.MarkFlagRequired("source-id")
	GetOnboardApiV1SourceSourceIdCmd.MarkFlagRequired("source-id")

	OnboardCmd.AddCommand(PostOnboardApiV1SourcesCmd)
	PostOnboardApiV1SourcesCmd.Flags().StringArray("source-ids", nil, "")

	PostOnboardApiV1SourcesCmd.Flags().String("type", "", "")

	OnboardCmd.AddCommand(GetOnboardApiV1CredentialCredentialIdHealthcheckCmd)
	GetOnboardApiV1CredentialCredentialIdHealthcheckCmd.Flags().String("credential-id", "", "")
	// GetOnboardApiV1CredentialCredentialIdHealthcheckCmd.MarkFlagRequired("credential-id")
	GetOnboardApiV1CredentialCredentialIdHealthcheckCmd.MarkFlagRequired("credential-id")

	OnboardCmd.AddCommand(PostOnboardApiV1SourceSourceIdHealthcheckCmd)
	PostOnboardApiV1SourceSourceIdHealthcheckCmd.Flags().String("source-id", "", "")
	// PostOnboardApiV1SourceSourceIdHealthcheckCmd.MarkFlagRequired("source-id")
	PostOnboardApiV1SourceSourceIdHealthcheckCmd.MarkFlagRequired("source-id")

	OnboardCmd.AddCommand(GetOnboardApiV1CredentialCmd)
	GetOnboardApiV1CredentialCmd.Flags().String("connector", "", "")
	GetOnboardApiV1CredentialCmd.Flags().String("credential-type", "", "")
	GetOnboardApiV1CredentialCmd.Flags().String("health", "", "")
	GetOnboardApiV1CredentialCmd.Flags().Int64("page-number", 0, "")
	GetOnboardApiV1CredentialCmd.Flags().Int64("page-size", 0, "")

	OnboardCmd.AddCommand(GetOnboardApiV1CredentialSourcesListCmd)
	GetOnboardApiV1CredentialSourcesListCmd.Flags().String("connector", "", "")
	GetOnboardApiV1CredentialSourcesListCmd.Flags().String("credential-type", "", "")
	GetOnboardApiV1CredentialSourcesListCmd.Flags().Int64("page-number", 0, "")
	GetOnboardApiV1CredentialSourcesListCmd.Flags().Int64("page-size", 0, "")

	OnboardCmd.AddCommand(GetOnboardApiV1SourcesCmd)
	GetOnboardApiV1SourcesCmd.Flags().StringArray("connector", nil, "")

	OnboardCmd.AddCommand(PutOnboardApiV1SourceSourceIdCredentialsCmd)
	PutOnboardApiV1SourceSourceIdCredentialsCmd.Flags().String("source-id", "", "")
	// PutOnboardApiV1SourceSourceIdCredentialsCmd.MarkFlagRequired("source-id")
	PutOnboardApiV1SourceSourceIdCredentialsCmd.MarkFlagRequired("source-id")

	OnboardCmd.AddCommand(GetOnboardApiV1SourceSourceIdCredentialsCmd)
	GetOnboardApiV1SourceSourceIdCredentialsCmd.Flags().String("source-id", "", "")
	// GetOnboardApiV1SourceSourceIdCredentialsCmd.MarkFlagRequired("source-id")
	GetOnboardApiV1SourceSourceIdCredentialsCmd.MarkFlagRequired("source-id")

	OnboardCmd.AddCommand(GetOnboardApiV1SourcesCountCmd)
	GetOnboardApiV1SourcesCountCmd.Flags().String("connector", "", "")

	OnboardCmd.AddCommand(PostOnboardApiV1ConnectionsConnectionIdStateCmd)
	PostOnboardApiV1ConnectionsConnectionIdStateCmd.Flags().Int64("connection-id", 0, "")
	// PostOnboardApiV1ConnectionsConnectionIdStateCmd.MarkFlagRequired("connection-id")
	PostOnboardApiV1ConnectionsConnectionIdStateCmd.MarkFlagRequired("connection-id")
	PostOnboardApiV1ConnectionsConnectionIdStateCmd.Flags().String("state", "", "")
	PostOnboardApiV1ConnectionsConnectionIdStateCmd.MarkFlagRequired("state")

	OnboardCmd.AddCommand(PutOnboardApiV1CredentialCredentialIdCmd)
	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("config", "", "")
	PutOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("config")
	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("connector", "", "")
	PutOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("connector")
	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("name", "", "")
	// PutOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("name")
	PutOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("name")

	PutOnboardApiV1CredentialCredentialIdCmd.Flags().String("credential-id", "", "")
	// PutOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("credential-id")
	PutOnboardApiV1CredentialCredentialIdCmd.MarkFlagRequired("credential-id")

}
