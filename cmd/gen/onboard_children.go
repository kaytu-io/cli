package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/onboard"
	"github.com/spf13/cobra"
)
var PostOnboardApiV1CredentialCredentialIdEnableCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_credential_credential_id_enable",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_enable] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1CredentialCredentialIDEnable(onboard.NewPostOnboardAPIV1CredentialCredentialIDEnableParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_enable] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_enable] : %v", err)
		}

		return nil
	},
}
var PostOnboardApiV1CredentialCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1Credential(onboard.NewPostOnboardAPIV1CredentialParams(), auth)
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
var PostOnboardApiV1SourceAwsCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_source_aws",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_aws] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1SourceAws(onboard.NewPostOnboardAPIV1SourceAwsParams(), auth)
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
var GetOnboardApiV1CatalogMetricsCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_catalog_metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_metrics] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1CatalogMetrics(onboard.NewGetOnboardAPIV1CatalogMetricsParams(), auth)
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
var GetOnboardApiV1ConnectorsConnectorNameCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_connectors_connector_name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connectors_connector_name] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1ConnectorsConnectorName(onboard.NewGetOnboardAPIV1ConnectorsConnectorNameParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connectors_connector_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connectors_connector_name] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ProvidersCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_providers] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1Providers(onboard.NewGetOnboardAPIV1ProvidersParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_providers] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_providers] : %v", err)
		}

		return nil
	},
}
var PostOnboardApiV1CredentialCredentialIdAutoonboardCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_credential_credential_id_autoonboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_autoonboard] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1CredentialCredentialIDAutoonboard(onboard.NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParams(), auth)
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
var PostOnboardApiV1CredentialCredentialIdDisableCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_credential_credential_id_disable",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_disable] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1CredentialCredentialIDDisable(onboard.NewPostOnboardAPIV1CredentialCredentialIDDisableParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_disable] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_credential_credential_id_disable] : %v", err)
		}

		return nil
	},
}
var PostOnboardApiV1SourceSourceIdHealthcheckCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_source_source_id_healthcheck",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_healthcheck] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1SourceSourceIDHealthcheck(onboard.NewPostOnboardAPIV1SourceSourceIDHealthcheckParams(), auth)
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
var PutOnboardApiV1CredentialCredentialIdCmd = &cobra.Command{
	Use: "put_onboard_api_v_1_credential_credential_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		resp, err := client.Onboard.PutOnboardAPIV1CredentialCredentialID(onboard.NewPutOnboardAPIV1CredentialCredentialIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		return nil
	},
}
var DeleteOnboardApiV1SourceSourceIdCmd = &cobra.Command{
	Use: "delete_onboard_api_v_1_source_source_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_source_source_id] : %v", err)
		}

		resp, err := client.Onboard.DeleteOnboardAPIV1SourceSourceID(onboard.NewDeleteOnboardAPIV1SourceSourceIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_source_source_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_source_source_id] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1CredentialSourcesListCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_credential_sources_list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_sources_list] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1CredentialSourcesList(onboard.NewGetOnboardAPIV1CredentialSourcesListParams(), auth)
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
var GetOnboardApiV1SourcesCountCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_sources_count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources_count] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1SourcesCount(onboard.NewGetOnboardAPIV1SourcesCountParams(), auth)
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
var GetOnboardApiV1SourcesCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_sources] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1Sources(onboard.NewGetOnboardAPIV1SourcesParams(), auth)
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
var DeleteOnboardApiV1CredentialCredentialIdCmd = &cobra.Command{
	Use: "delete_onboard_api_v_1_credential_credential_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		resp, err := client.Onboard.DeleteOnboardAPIV1CredentialCredentialID(onboard.NewDeleteOnboardAPIV1CredentialCredentialIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ConnectionsCountCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_connections_count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_count] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1ConnectionsCount(onboard.NewGetOnboardAPIV1ConnectionsCountParams(), auth)
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
var PostOnboardApiV1SourcesCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_sources] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1Sources(onboard.NewPostOnboardAPIV1SourcesParams(), auth)
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
var PostOnboardApiV1SourceSourceIdEnableCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_source_source_id_enable",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_enable] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1SourceSourceIDEnable(onboard.NewPostOnboardAPIV1SourceSourceIDEnableParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_enable] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_enable] : %v", err)
		}

		return nil
	},
}
var PutOnboardApiV1SourceSourceIdCredentialsCmd = &cobra.Command{
	Use: "put_onboard_api_v_1_source_source_id_credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		resp, err := client.Onboard.PutOnboardAPIV1SourceSourceIDCredentials(onboard.NewPutOnboardAPIV1SourceSourceIDCredentialsParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1CredentialCredentialIdCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_credential_credential_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1CredentialCredentialID(onboard.NewGetOnboardAPIV1CredentialCredentialIDParams(), auth)
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
	Use: "get_onboard_api_v_1_credential",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1Credential(onboard.NewGetOnboardAPIV1CredentialParams(), auth)
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
var GetOnboardApiV1SourceSourceIdCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_source_source_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1SourceSourceID(onboard.NewGetOnboardAPIV1SourceSourceIDParams(), auth)
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
var PostOnboardApiV1SourceAzureCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_source_azure",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_azure] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1SourceAzure(onboard.NewPostOnboardAPIV1SourceAzureParams(), auth)
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
var PostOnboardApiV1SourceSourceIdDisableCmd = &cobra.Command{
	Use: "post_onboard_api_v_1_source_source_id_disable",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_disable] : %v", err)
		}

		resp, err := client.Onboard.PostOnboardAPIV1SourceSourceIDDisable(onboard.NewPostOnboardAPIV1SourceSourceIDDisableParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_disable] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_onboard_api_v_1_source_source_id_disable] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1SourceAccountAccountIdCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_source_account_account_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_account_account_id] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1SourceAccountAccountID(onboard.NewGetOnboardAPIV1SourceAccountAccountIDParams(), auth)
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
var GetOnboardApiV1CredentialCredentialIdHealthcheckCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_credential_credential_id_healthcheck",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id_healthcheck] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1CredentialCredentialIDHealthcheck(onboard.NewGetOnboardAPIV1CredentialCredentialIDHealthcheckParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id_healthcheck] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_credential_credential_id_healthcheck] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ProvidersTypesCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_providers_types",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_providers_types] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1ProvidersTypes(onboard.NewGetOnboardAPIV1ProvidersTypesParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_providers_types] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_providers_types] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1SourceSourceIdCredentialsCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_source_source_id_credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_source_source_id_credentials] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1SourceSourceIDCredentials(onboard.NewGetOnboardAPIV1SourceSourceIDCredentialsParams(), auth)
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
var GetOnboardApiV1CatalogConnectorsCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_catalog_connectors",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_connectors] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1CatalogConnectors(onboard.NewGetOnboardAPIV1CatalogConnectorsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_connectors] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_catalog_connectors] : %v", err)
		}

		return nil
	},
}
var GetOnboardApiV1ConnectorsCmd = &cobra.Command{
	Use: "get_onboard_api_v_1_connectors",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connectors] : %v", err)
		}

		resp, err := client.Onboard.GetOnboardAPIV1Connectors(onboard.NewGetOnboardAPIV1ConnectorsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connectors] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connectors] : %v", err)
		}

		return nil
	},
}

}
