package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/spf13/cobra"
)
var PostAuthApiV1KeyCreateCmd = &cobra.Command{
	Use: "post_auth_api_v_1_key_create",
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
	Use: "post_auth_api_v_1_key_id_activate",
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
	Use: "post_auth_api_v_1_key_id_suspend",
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
	Use: "post_auth_api_v_1_key_role",
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
var DeleteAuthApiV1KeyIdDeleteCmd = &cobra.Command{
	Use: "delete_auth_api_v_1_key_id_delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

		resp, err := client.Keys.DeleteAuthAPIV1KeyIDDelete(keys.NewDeleteAuthAPIV1KeyIDDeleteParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1KeyIdCmd = &cobra.Command{
	Use: "get_auth_api_v_1_key_id",
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
	Use: "get_auth_api_v_1_keys",
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

},
}
