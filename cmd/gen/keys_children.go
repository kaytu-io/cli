package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/spf13/cobra"
)

var DeleteAuthApiV1KeyIdDeleteCmd = &cobra.Command{
	Use: "idIdDelete",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

		_, err = client.Keys.DeleteAuthAPIV1KeyIDDelete(keys.NewDeleteAuthAPIV1KeyIDDeleteParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1KeyIdCmd = &cobra.Command{
	Use: "idId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_key_id] : %v", err)
		}

		resp, err := client.Keys.GetAuthAPIV1KeyID(keys.NewGetAuthAPIV1KeyIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_key_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_key_id] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1KeysCmd = &cobra.Command{
	Use: "Keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_keys] : %v", err)
		}

		resp, err := client.Keys.GetAuthAPIV1Keys(keys.NewGetAuthAPIV1KeysParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_keys] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_keys] : %v", err)
		}

		return nil
	},
}

var PostAuthApiV1KeyCreateCmd = &cobra.Command{
	Use: "createCreate",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_create] : %v", err)
		}

		resp, err := client.Keys.PostAuthAPIV1KeyCreate(keys.NewPostAuthAPIV1KeyCreateParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_create] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_create] : %v", err)
		}

		return nil
	},
}

var PostAuthApiV1KeyIdActivateCmd = &cobra.Command{
	Use: "idIdActivate",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_activate] : %v", err)
		}

		resp, err := client.Keys.PostAuthAPIV1KeyIDActivate(keys.NewPostAuthAPIV1KeyIDActivateParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_activate] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_activate] : %v", err)
		}

		return nil
	},
}

var PostAuthApiV1KeyIdSuspendCmd = &cobra.Command{
	Use: "idIdSuspend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_suspend] : %v", err)
		}

		resp, err := client.Keys.PostAuthAPIV1KeyIDSuspend(keys.NewPostAuthAPIV1KeyIDSuspendParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_suspend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_suspend] : %v", err)
		}

		return nil
	},
}

var PostAuthApiV1KeyRoleCmd = &cobra.Command{
	Use: "roleRole",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_role] : %v", err)
		}

		resp, err := client.Keys.PostAuthAPIV1KeyRole(keys.NewPostAuthAPIV1KeyRoleParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_role] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_role] : %v", err)
		}

		return nil
	},
}
