package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/metadata"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2MetadataResourcetypeResourceTypeCmd = &cobra.Command{
	Use: "get-resource-type",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype_resource_type] : %v", err)
		}

		req := metadata.NewGetInventoryAPIV2MetadataResourcetypeResourceTypeParams()

		req.SetResourceType(flags.ReadStringFlag(cmd, "ResourceType"))

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
	Use: "list-services",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services] : %v", err)
		}

		req := metadata.NewGetInventoryAPIV2MetadataServicesParams()

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

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
	Use: "get-service",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services_service_name] : %v", err)
		}

		req := metadata.NewGetInventoryAPIV2MetadataServicesServiceNameParams()

		req.SetServiceName(flags.ReadStringFlag(cmd, "ServiceName"))

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
	Use: "get-config-metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_metadata_api_v_1_metadata_key] : %v", err)
		}

		req := metadata.NewGetMetadataAPIV1MetadataKeyParams()

		req.SetKey(flags.ReadStringFlag(cmd, "Key"))

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
	Use: "set-config-metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_metadata_api_v_1_metadata] : %v", err)
		}

		req := metadata.NewPostMetadataAPIV1MetadataParams()

		req.SetReq(&models.GithubComKaytuIoKaytuEnginePkgMetadataAPISetConfigMetadataRequest{
			Key:   flags.ReadStringFlag(cmd, "Key"),
			Value: interface{}(flags.ReadStringFlag(cmd, "Value")),
		})

		_, err = client.Metadata.PostMetadataAPIV1Metadata(req, auth)
		if err != nil {
			return fmt.Errorf("[post_metadata_api_v_1_metadata] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2MetadataResourcetypeCmd = &cobra.Command{
	Use: "list-resource-types",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype] : %v", err)
		}

		req := metadata.NewGetInventoryAPIV2MetadataResourcetypeParams()

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetService(flags.ReadStringArrayFlag(cmd, "Service"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

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
