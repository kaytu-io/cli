package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/roles"
	"github.com/spf13/cobra"
)

var GetAuthApiV1RoleRoleNameKeysCmd = &cobra.Command{
	Use: "roleRoleNameKeys",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_keys] : %v", err)
		}

		resp, err := client.Roles.GetAuthAPIV1RoleRoleNameKeys(roles.NewGetAuthAPIV1RoleRoleNameKeysParams(), auth)
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
	Use: "roleRoleNameUsers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_role_role_name_users] : %v", err)
		}

		resp, err := client.Roles.GetAuthAPIV1RoleRoleNameUsers(roles.NewGetAuthAPIV1RoleRoleNameUsersParams(), auth)
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
	Use: "Roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles] : %v", err)
		}

		resp, err := client.Roles.GetAuthAPIV1Roles(roles.NewGetAuthAPIV1RolesParams(), auth)
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

var GetAuthApiV1RolesRoleNameCmd = &cobra.Command{
	Use: "roleRoleName",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_roles_role_name] : %v", err)
		}

		resp, err := client.Roles.GetAuthAPIV1RolesRoleName(roles.NewGetAuthAPIV1RolesRoleNameParams(), auth)
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
