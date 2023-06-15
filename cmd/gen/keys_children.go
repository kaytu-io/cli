package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var PostAuthApiV1KeyRoleCmd = &cobra.Command{
	Use: "key-role",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_role] : %v", err)
		}

        req := keys.NewPostAuthAPIV1KeyRoleParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgAuthAPIUpdateKeyRoleRequest{
ID: flags.ReadInt64Flag(cmd, "ID"),
RoleName: models.GitlabComKeibiengineKeibiEnginePkgAuthAPIRole(flags.ReadStringFlag(cmd, "RoleName")),

})


		resp, err := client.Keys.PostAuthAPIV1KeyRole(req, auth)
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
	Use: "key-id-delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

        req := keys.NewDeleteAuthAPIV1KeyIDDeleteParams()

        req.SetID(flags.ReadStringFlag(cmd, "ID"))


		_, err = client.Keys.DeleteAuthAPIV1KeyIDDelete(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_key_id_delete] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1KeyIdCmd = &cobra.Command{
	Use: "key-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_key_id] : %v", err)
		}

        req := keys.NewGetAuthAPIV1KeyIDParams()

        req.SetID(flags.ReadStringFlag(cmd, "ID"))


		resp, err := client.Keys.GetAuthAPIV1KeyID(req, auth)
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
	Use: "keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_keys] : %v", err)
		}

        req := keys.NewGetAuthAPIV1KeysParams()

        

		resp, err := client.Keys.GetAuthAPIV1Keys(req, auth)
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
	Use: "key-create",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_create] : %v", err)
		}

        req := keys.NewPostAuthAPIV1KeyCreateParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgAuthAPICreateAPIKeyRequest{
Name: flags.ReadStringFlag(cmd, "Name"),
RoleName: models.GitlabComKeibiengineKeibiEnginePkgAuthAPIRole(flags.ReadStringFlag(cmd, "RoleName")),

})


		resp, err := client.Keys.PostAuthAPIV1KeyCreate(req, auth)
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
	Use: "key-id-activate",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_activate] : %v", err)
		}

        req := keys.NewPostAuthAPIV1KeyIDActivateParams()

        req.SetID(flags.ReadStringFlag(cmd, "ID"))


		resp, err := client.Keys.PostAuthAPIV1KeyIDActivate(req, auth)
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
	Use: "key-id-suspend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_key_id_suspend] : %v", err)
		}

        req := keys.NewPostAuthAPIV1KeyIDSuspendParams()

        req.SetID(flags.ReadStringFlag(cmd, "ID"))


		resp, err := client.Keys.PostAuthAPIV1KeyIDSuspend(req, auth)
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
