package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/spf13/cobra"
)
var GetAuthApiV1WorkspaceRoleBindingsCmd = &cobra.Command{
	Use: "get_auth_api_v_1_workspace_role_bindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

		resp, err := client.Users.GetAuthAPIV1WorkspaceRoleBindings(users.NewGetAuthAPIV1WorkspaceRoleBindingsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

		return nil
	},
}
var DeleteAuthApiV1UserRoleBindingCmd = &cobra.Command{
	Use: "delete_auth_api_v_1_user_role_binding",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

		resp, err := client.Users.DeleteAuthAPIV1UserRoleBinding(users.NewDeleteAuthAPIV1UserRoleBindingParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1UserRoleBindingsCmd = &cobra.Command{
	Use: "get_auth_api_v_1_user_role_bindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

		resp, err := client.Users.GetAuthAPIV1UserRoleBindings(users.NewGetAuthAPIV1UserRoleBindingsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1UserUserIdWorkspaceMembershipCmd = &cobra.Command{
	Use: "get_auth_api_v_1_user_user_id_workspace_membership",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

		resp, err := client.Users.GetAuthAPIV1UserUserIDWorkspaceMembership(users.NewGetAuthAPIV1UserUserIDWorkspaceMembershipParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

		return nil
	},
}
var PostAuthApiV1UserInviteCmd = &cobra.Command{
	Use: "post_auth_api_v_1_user_invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

		resp, err := client.Users.PostAuthAPIV1UserInvite(users.NewPostAuthAPIV1UserInviteParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

		return nil
	},
}
var PutAuthApiV1UserRoleBindingCmd = &cobra.Command{
	Use: "put_auth_api_v_1_user_role_binding",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

		resp, err := client.Users.PutAuthAPIV1UserRoleBinding(users.NewPutAuthAPIV1UserRoleBindingParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

		return nil
	},
}
var DeleteAuthApiV1UserInviteCmd = &cobra.Command{
	Use: "delete_auth_api_v_1_user_invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

		resp, err := client.Users.DeleteAuthAPIV1UserInvite(users.NewDeleteAuthAPIV1UserInviteParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1UserUserIdCmd = &cobra.Command{
	Use: "get_auth_api_v_1_user_user_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

		resp, err := client.Users.GetAuthAPIV1UserUserID(users.NewGetAuthAPIV1UserUserIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1UsersCmd = &cobra.Command{
	Use: "get_auth_api_v_1_users",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

		resp, err := client.Users.GetAuthAPIV1Users(users.NewGetAuthAPIV1UsersParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

		return nil
	},
}
