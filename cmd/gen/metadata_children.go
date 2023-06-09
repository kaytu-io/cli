package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/metadata"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2MetadataResourcetypeCmd = &cobra.Command{
	Use: "resourcetypeResourcetype",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype] : %v", err)
		}

		resp, err := client.Metadata.GetInventoryAPIV2MetadataResourcetype(metadata.NewGetInventoryAPIV2MetadataResourcetypeParams(), auth)
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
	Use: "resourcetypeResourcetypeResourceType",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_resourcetype_resource_type] : %v", err)
		}

		resp, err := client.Metadata.GetInventoryAPIV2MetadataResourcetypeResourceType(metadata.NewGetInventoryAPIV2MetadataResourcetypeResourceTypeParams(), auth)
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
	Use: "servicesServices",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services] : %v", err)
		}

		resp, err := client.Metadata.GetInventoryAPIV2MetadataServices(metadata.NewGetInventoryAPIV2MetadataServicesParams(), auth)
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
	Use: "servicesServicesServiceName",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_metadata_services_service_name] : %v", err)
		}

		resp, err := client.Metadata.GetInventoryAPIV2MetadataServicesServiceName(metadata.NewGetInventoryAPIV2MetadataServicesServiceNameParams(), auth)
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
	Use: "keyKey",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_metadata_api_v_1_metadata_key] : %v", err)
		}

		resp, err := client.Metadata.GetMetadataAPIV1MetadataKey(metadata.NewGetMetadataAPIV1MetadataKeyParams(), auth)
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
	Use: "Metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_metadata_api_v_1_metadata] : %v", err)
		}

		_, err = client.Metadata.PostMetadataAPIV1Metadata(metadata.NewPostMetadataAPIV1MetadataParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_metadata_api_v_1_metadata] : %v", err)
		}

		return nil
	},
}