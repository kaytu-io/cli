package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var GetAuthApiV1WorkspaceRoleBindingsCmd = &cobra.Command{
	Use: "workspace-role-bindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

        req := users.NewGetAuthAPIV1WorkspaceRoleBindingsParams()

        

		resp, err := client.Users.GetAuthAPIV1WorkspaceRoleBindings(req, auth)
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
	Use: "user-role-binding",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

        req := users.NewDeleteAuthAPIV1UserRoleBindingParams()

        req.SetUserID(flags.ReadStringFlag("UserID"))


		_, err = client.Users.DeleteAuthAPIV1UserRoleBinding(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1UserRoleBindingsCmd = &cobra.Command{
	Use: "user-role-bindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

        req := users.NewGetAuthAPIV1UserRoleBindingsParams()

        

		resp, err := client.Users.GetAuthAPIV1UserRoleBindings(req, auth)
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

var GetAuthApiV1UserUserIdCmd = &cobra.Command{
	Use: "user-user-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

        req := users.NewGetAuthAPIV1UserUserIDParams()

        req.SetUserID(flags.ReadStringFlag("UserID"))


		resp, err := client.Users.GetAuthAPIV1UserUserID(req, auth)
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
	Use: "users",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

        req := users.NewGetAuthAPIV1UsersParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUsersRequest{
Email: flags.ReadStringFlag("Email"),
EmailVerified: flags.ReadBooleanFlag("EmailVerified"),
RoleName: models.GitlabComKeibiengineKeibiEnginePkgAuthAPIRole(flags.ReadStringFlag("RoleName")),

})


		resp, err := client.Users.GetAuthAPIV1Users(req, auth)
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

var DeleteAuthApiV1UserInviteCmd = &cobra.Command{
	Use: "user-invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

        req := users.NewDeleteAuthAPIV1UserInviteParams()

        req.SetUserID(flags.ReadStringFlag("UserID"))


		_, err = client.Users.DeleteAuthAPIV1UserInvite(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

		return nil
	},
}
var GetAuthApiV1UserUserIdWorkspaceMembershipCmd = &cobra.Command{
	Use: "user-user-id-workspace-membership",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

        req := users.NewGetAuthAPIV1UserUserIDWorkspaceMembershipParams()

        req.SetUserID(flags.ReadStringFlag("UserID"))


		resp, err := client.Users.GetAuthAPIV1UserUserIDWorkspaceMembership(req, auth)
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
	Use: "user-invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

        req := users.NewPostAuthAPIV1UserInviteParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgAuthAPIInviteRequest{
Email: flags.ReadStringOptionalFlag("Email"),
RoleName: models.GitlabComKeibiengineKeibiEnginePkgAuthAPIRole(flags.ReadStringFlag("RoleName")),

})


		_, err = client.Users.PostAuthAPIV1UserInvite(req, auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

		return nil
	},
}
var PutAuthApiV1UserRoleBindingCmd = &cobra.Command{
	Use: "user-role-binding",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

        req := users.NewPutAuthAPIV1UserRoleBindingParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest{
RoleName: models.GitlabComKeibiengineKeibiEnginePkgAuthAPIRole(flags.ReadStringFlag("RoleName")),
UserID: flags.ReadStringOptionalFlag("UserID"),

})


		_, err = client.Users.PutAuthAPIV1UserRoleBinding(req, auth)
		if err != nil {
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

		return nil
	},
}