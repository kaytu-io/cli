package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/roles"
	"github.com/spf13/cobra"
)

var GetAuthApiV1RolesRoleNameCmd = &cobra.Command{
	Use: "roles-role-name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles_role_name] : %v", err)
		}

		req := roles.NewGetAuthAPIV1RolesRoleNameParams()

		req.SetRoleName(flags.ReadStringFlag("RoleName"))

		resp, err := client.Roles.GetAuthAPIV1RolesRoleName(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles_role_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles_role_name] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1RoleRoleNameKeysCmd = &cobra.Command{
	Use: "role-role-name-keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_keys] : %v", err)
		}

		req := roles.NewGetAuthAPIV1RoleRoleNameKeysParams()

		req.SetRoleName(flags.ReadStringFlag("RoleName"))

		resp, err := client.Roles.GetAuthAPIV1RoleRoleNameKeys(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_keys] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_keys] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1RoleRoleNameUsersCmd = &cobra.Command{
	Use: "role-role-name-users",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_users] : %v", err)
		}

		req := roles.NewGetAuthAPIV1RoleRoleNameUsersParams()

		req.SetRoleName(flags.ReadStringFlag("RoleName"))

		resp, err := client.Roles.GetAuthAPIV1RoleRoleNameUsers(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_users] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_users] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1RolesCmd = &cobra.Command{
	Use: "roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles] : %v", err)
		}

		req := roles.NewGetAuthAPIV1RolesParams()

		resp, err := client.Roles.GetAuthAPIV1Roles(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles] : %v", err)
		}

		return nil
	},
}
