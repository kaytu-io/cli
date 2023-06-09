package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/metadata"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var GetInventoryApiV2MetadataResourcetypeCmd = &cobra.Command{
	Use: "metadata-resourcetype",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype] : %v", err)
		}

        req := metadata.NewGetInventoryAPIV2MetadataResourcetypeParams()

        req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetPageNumber(flags.ReadInt64OptionalFlag("PageNumber"))
req.SetPageSize(flags.ReadInt64OptionalFlag("PageSize"))
req.SetService(flags.ReadStringArrayFlag("Service"))
req.SetTag(flags.ReadStringArrayFlag("Tag"))


		resp, err := client.Metadata.GetInventoryAPIV2MetadataResourcetype(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2MetadataResourcetypeResourceTypeCmd = &cobra.Command{
	Use: "metadata-resourcetype-resource-type",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype_resource_type] : %v", err)
		}

        req := metadata.NewGetInventoryAPIV2MetadataResourcetypeResourceTypeParams()

        req.SetResourceType(flags.ReadStringFlag("ResourceType"))


		resp, err := client.Metadata.GetInventoryAPIV2MetadataResourcetypeResourceType(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype_resource_type] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype_resource_type] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2MetadataServicesCmd = &cobra.Command{
	Use: "metadata-services",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services] : %v", err)
		}

        req := metadata.NewGetInventoryAPIV2MetadataServicesParams()

        req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetCostSupport(flags.ReadBooleanOptionalFlag("CostSupport"))
req.SetPageNumber(flags.ReadInt64OptionalFlag("PageNumber"))
req.SetPageSize(flags.ReadInt64OptionalFlag("PageSize"))
req.SetTag(flags.ReadStringArrayFlag("Tag"))


		resp, err := client.Metadata.GetInventoryAPIV2MetadataServices(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2MetadataServicesServiceNameCmd = &cobra.Command{
	Use: "metadata-services-service-name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services_service_name] : %v", err)
		}

        req := metadata.NewGetInventoryAPIV2MetadataServicesServiceNameParams()

        req.SetServiceName(flags.ReadStringFlag("ServiceName"))


		resp, err := client.Metadata.GetInventoryAPIV2MetadataServicesServiceName(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services_service_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services_service_name] : %v", err)
		}

		return nil
	},
}

var GetMetadataApiV1MetadataKeyCmd = &cobra.Command{
	Use: "metadata-key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_metadata_api_v_1_metadata_key] : %v", err)
		}

        req := metadata.NewGetMetadataAPIV1MetadataKeyParams()

        req.SetKey(flags.ReadStringFlag("Key"))


		resp, err := client.Metadata.GetMetadataAPIV1MetadataKey(req, auth)
		if err != nil {
			return fmt.Errorf("[get_metadata_api_v_1_metadata_key] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_metadata_api_v_1_metadata_key] : %v", err)
		}

		return nil
	},
}

var PostMetadataApiV1MetadataCmd = &cobra.Command{
	Use: "metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_metadata_api_v_1_metadata] : %v", err)
		}

        req := metadata.NewPostMetadataAPIV1MetadataParams()

        

		_, err = client.Metadata.PostMetadataAPIV1Metadata(req, auth)
		if err != nil {
			return fmt.Errorf("[post_metadata_api_v_1_metadata] : %v", err)
		}

		return nil
	},
}