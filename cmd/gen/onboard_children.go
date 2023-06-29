package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/onboard"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var PostOnboardApiV1ConnectionsConnectionIdStateCmd = &cobra.Command{
	Use: "update-connection-state",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_connections_connection_id_state] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1ConnectionsConnectionIDStateParams()

		req.SetConnectionID(flags.ReadInt64Flag(cmd, "ConnectionID"))
		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPIChangeConnectionLifecycleStateRequest{
			State: models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionLifecycleState(flags.ReadStringFlag(cmd, "State")),
		})

		_, err = client.Onboard.PostOnboardAPIV1ConnectionsConnectionIDState(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_connections_connection_id_state] : %v", err)
		}

		return nil
	},
}
var PostOnboardApiV1CredentialCredentialIdDisableCmd = &cobra.Command{
	Use: "disable-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_disable] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1CredentialCredentialIDDisableParams()

		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		_, err = client.Onboard.PostOnboardAPIV1CredentialCredentialIDDisable(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_disable] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ConnectionsCountCmd = &cobra.Command{
	Use: "get-connections-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_count] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1ConnectionsCountParams()

		req.SetType(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionCountRequest{
			Connectors: flags.ReadStringArrayFlag(cmd, "Connectors"),
			State:      models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionLifecycleState(flags.ReadStringFlag(cmd, "State")),
		})

		resp, err := client.Onboard.GetOnboardAPIV1ConnectionsCount(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_count] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1SourcesCmd = &cobra.Command{
	Use: "list-sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1SourcesParams()

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))

		resp, err := client.Onboard.GetOnboardAPIV1Sources(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1CredentialCredentialIdHealthcheckCmd = &cobra.Command{
	Use: "healthcheck-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id_healthcheck] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1CredentialCredentialIDHealthcheckParams()

		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		_, err = client.Onboard.GetOnboardAPIV1CredentialCredentialIDHealthcheck(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id_healthcheck] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1SourceSourceIdCredentialsCmd = &cobra.Command{
	Use: "source-credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1SourceSourceIDCredentialsParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		resp, err := client.Onboard.GetOnboardAPIV1SourceSourceIDCredentials(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1SourceSourceIdCmd = &cobra.Command{
	Use: "get-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1SourceSourceIDParams()

		req.SetSourceID(flags.ReadInt64Flag(cmd, "SourceID"))

		resp, err := client.Onboard.GetOnboardAPIV1SourceSourceID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1SourcesCountCmd = &cobra.Command{
	Use: "sources-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources_count] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1SourcesCountParams()

		req.SetConnector(flags.ReadStringOptionalFlag(cmd, "Connector"))

		resp, err := client.Onboard.GetOnboardAPIV1SourcesCount(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources_count] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1CredentialCredentialIdAutoonboardCmd = &cobra.Command{
	Use: "auto-onboard-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_autoonboard] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParams()

		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		resp, err := client.Onboard.PostOnboardAPIV1CredentialCredentialIDAutoonboard(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_autoonboard] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_autoonboard] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1CredentialCredentialIdEnableCmd = &cobra.Command{
	Use: "enable-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_enable] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1CredentialCredentialIDEnableParams()

		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		_, err = client.Onboard.PostOnboardAPIV1CredentialCredentialIDEnable(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_enable] : %v", err)
		}

		return nil
	},
}
var DeleteOnboardApiV1CredentialCredentialIdCmd = &cobra.Command{
	Use: "delete-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		req := onboard.NewDeleteOnboardAPIV1CredentialCredentialIDParams()

		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		_, err = client.Onboard.DeleteOnboardAPIV1CredentialCredentialID(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ConnectorCmd = &cobra.Command{
	Use: "get-connectors",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connector] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1ConnectorParams()

		resp, err := client.Onboard.GetOnboardAPIV1Connector(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connector] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connector] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1SourcesCmd = &cobra.Command{
	Use: "get-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_sources] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1SourcesParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPIGetSourcesRequest{
			SourceIds: flags.ReadStringArrayFlag(cmd, "SourceIds"),
		})
		req.SetType(flags.ReadStringOptionalFlag(cmd, "Type"))

		resp, err := client.Onboard.PostOnboardAPIV1Sources(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_sources] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_sources] : %v", err)
		}

		return nil
	},
}

var PutOnboardApiV1CredentialCredentialIdCmd = &cobra.Command{
	Use: "update-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		req := onboard.NewPutOnboardAPIV1CredentialCredentialIDParams()

		req.SetConfig(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest{
			Config:    interface{}(flags.ReadStringFlag(cmd, "Config")),
			Connector: models.SourceType(flags.ReadStringFlag(cmd, "Connector")),
			Name:      flags.ReadStringFlag(cmd, "Name"),
		})
		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		_, err = client.Onboard.PutOnboardAPIV1CredentialCredentialID(req, auth)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1CredentialCredentialIdCmd = &cobra.Command{
	Use: "get-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1CredentialCredentialIDParams()

		req.SetCredentialID(flags.ReadStringFlag(cmd, "CredentialID"))

		resp, err := client.Onboard.GetOnboardAPIV1CredentialCredentialID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1CredentialCmd = &cobra.Command{
	Use: "list-credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1CredentialParams()

		req.SetConnector(flags.ReadStringOptionalFlag(cmd, "Connector"))
		req.SetHealth(flags.ReadStringOptionalFlag(cmd, "Health"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))

		resp, err := client.Onboard.GetOnboardAPIV1Credential(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1CredentialSourcesListCmd = &cobra.Command{
	Use: "list-credentials-sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_sources_list] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1CredentialSourcesListParams()

		req.SetConnector(flags.ReadStringOptionalFlag(cmd, "Connector"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))

		resp, err := client.Onboard.GetOnboardAPIV1CredentialSourcesList(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_sources_list] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_sources_list] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1SourceAwsCmd = &cobra.Command{
	Use: "create-aws-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_aws] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1SourceAwsParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPISourceAwsRequest{
			Config: &models.GithubComKaytuIoKaytuEnginePkgOnboardAPISourceConfigAWS{
				AccessKey: flags.ReadStringOptionalFlag(cmd, "AccessKey"),
				AccountID: flags.ReadStringFlag(cmd, "AccountID"),
				Regions:   flags.ReadStringArrayFlag(cmd, "Regions"),
				SecretKey: flags.ReadStringOptionalFlag(cmd, "SecretKey"),
			},
			Description: flags.ReadStringFlag(cmd, "Description"),
			Email:       flags.ReadStringFlag(cmd, "Email"),
			Name:        flags.ReadStringFlag(cmd, "Name"),
		})

		resp, err := client.Onboard.PostOnboardAPIV1SourceAws(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_aws] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_aws] : %v", err)
		}

		return nil
	},
}

var PutOnboardApiV1SourceSourceIdCredentialsCmd = &cobra.Command{
	Use: "put-source-credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		req := onboard.NewPutOnboardAPIV1SourceSourceIDCredentialsParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		_, err = client.Onboard.PutOnboardAPIV1SourceSourceIDCredentials(req, auth)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		return nil
	},
}
var DeleteOnboardApiV1SourceSourceIdCmd = &cobra.Command{
	Use: "delete-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_source_source_id] : %v", err)
		}

		req := onboard.NewDeleteOnboardAPIV1SourceSourceIDParams()

		req.SetSourceID(flags.ReadInt64Flag(cmd, "SourceID"))

		_, err = client.Onboard.DeleteOnboardAPIV1SourceSourceID(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_source_source_id] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ConnectorConnectorNameCmd = &cobra.Command{
	Use: "get-connector",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connector_connector_name] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1ConnectorConnectorNameParams()

		req.SetConnectorName(flags.ReadStringFlag(cmd, "ConnectorName"))

		resp, err := client.Onboard.GetOnboardAPIV1ConnectorConnectorName(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connector_connector_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connector_connector_name] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1CredentialCmd = &cobra.Command{
	Use: "create-credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1CredentialParams()

		req.SetConfig(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest{
			Config:     interface{}(flags.ReadStringFlag(cmd, "Config")),
			Name:       flags.ReadStringFlag(cmd, "Name"),
			SourceType: models.SourceType(flags.ReadStringFlag(cmd, "SourceType")),
		})

		resp, err := client.Onboard.PostOnboardAPIV1Credential(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1SourceAzureCmd = &cobra.Command{
	Use: "create-azure-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_azure] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1SourceAzureParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgOnboardAPISourceAzureRequest{
			Config: &models.GithubComKaytuIoKaytuEnginePkgOnboardAPISourceConfigAzure{
				ClientID:       flags.ReadStringOptionalFlag(cmd, "ClientID"),
				ClientSecret:   flags.ReadStringOptionalFlag(cmd, "ClientSecret"),
				ObjectID:       flags.ReadStringOptionalFlag(cmd, "ObjectID"),
				SecretID:       flags.ReadStringOptionalFlag(cmd, "SecretID"),
				SubscriptionID: flags.ReadStringFlag(cmd, "SubscriptionID"),
				TenantID:       flags.ReadStringOptionalFlag(cmd, "TenantID"),
			},
			Description: flags.ReadStringFlag(cmd, "Description"),
			Name:        flags.ReadStringFlag(cmd, "Name"),
		})

		resp, err := client.Onboard.PostOnboardAPIV1SourceAzure(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_azure] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_azure] : %v", err)
		}

		return nil
	},
}

var PostOnboardApiV1SourceSourceIdHealthcheckCmd = &cobra.Command{
	Use: "healthchek-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_healthcheck] : %v", err)
		}

		req := onboard.NewPostOnboardAPIV1SourceSourceIDHealthcheckParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		resp, err := client.Onboard.PostOnboardAPIV1SourceSourceIDHealthcheck(req, auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_healthcheck] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_healthcheck] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1CatalogMetricsCmd = &cobra.Command{
	Use: "get-catalog-metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_metrics] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1CatalogMetricsParams()

		resp, err := client.Onboard.GetOnboardAPIV1CatalogMetrics(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_metrics] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_metrics] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1SourceAccountAccountIdCmd = &cobra.Command{
	Use: "get-account-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_account_account_id] : %v", err)
		}

		req := onboard.NewGetOnboardAPIV1SourceAccountAccountIDParams()

		req.SetAccountID(flags.ReadInt64Flag(cmd, "AccountID"))

		resp, err := client.Onboard.GetOnboardAPIV1SourceAccountAccountID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_account_account_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_account_account_id] : %v", err)
		}

		return nil
	},
}
