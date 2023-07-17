// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var DeleteAuthApiV1UserInviteCmd = &cobra.Command{
	Use:   "delete-user-invite",
	Short: `Revokes user's access to the workspace.`,
	Long:  `Revokes user's access to the workspace.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

		req := users.NewDeleteAuthAPIV1UserInviteParams()

		req.SetUserID(flags.ReadStringFlag(cmd, "UserID"))

		_, err = client.Users.DeleteAuthAPIV1UserInvite(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_invite] : %v", err)
		}

		return nil
	},
}

var DeleteAuthApiV1UserRoleBindingCmd = &cobra.Command{
	Use:   "delete-user-role-binding",
	Short: `Revokes a user's access to the workspace`,
	Long:  `Revokes a user's access to the workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

		req := users.NewDeleteAuthAPIV1UserRoleBindingParams()

		req.SetUserID(flags.ReadStringFlag(cmd, "UserID"))

		_, err = client.Users.DeleteAuthAPIV1UserRoleBinding(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_auth_api_v_1_user_role_binding] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1UserUserIdCmd = &cobra.Command{
	Use:   "get-user",
	Short: `Returns user details by specified user id.`,
	Long:  `Returns user details by specified user id.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

		req := users.NewGetAuthAPIV1UserUserIDParams()

		req.SetUserID(flags.ReadStringFlag(cmd, "UserID"))

		resp, err := client.Users.GetAuthAPIV1UserUserID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-auth-api-v1-user-user-id", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id] : %v", err)
		}

		return nil
	},
}

var PostAuthApiV1UserInviteCmd = &cobra.Command{
	Use:   "invite-user",
	Short: `Sends an invitation to a user to join the workspace with a designated role.`,
	Long:  `Sends an invitation to a user to join the workspace with a designated role.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

		req := users.NewPostAuthAPIV1UserInviteParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgAuthAPIInviteRequest{
			Email:    flags.ReadStringOptionalFlag(cmd, "Email"),
			RoleName: models.GithubComKaytuIoKaytuEnginePkgAuthAPIRole(flags.ReadStringFlag(cmd, "RoleName")),
		})

		_, err = client.Users.PostAuthAPIV1UserInvite(req, auth)
		if err != nil {
			return fmt.Errorf("[post_auth_api_v_1_user_invite] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1UserRoleBindingsCmd = &cobra.Command{
	Use:   "list-user-role-bindings",
	Short: `Retrieves the roles that the user who sent the request has in all workspaces they are a member of.`,
	Long:  `Retrieves the roles that the user who sent the request has in all workspaces they are a member of.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

		req := users.NewGetAuthAPIV1UserRoleBindingsParams()

		resp, err := client.Users.GetAuthAPIV1UserRoleBindings(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-auth-api-v1-user-role-bindings", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_role_bindings] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1UserUserIdWorkspaceMembershipCmd = &cobra.Command{
	Use:   "list-user-workspace-memberships",
	Short: `Returns a list of workspaces that the specified user belongs to, along with their role in each workspace.`,
	Long:  `Returns a list of workspaces that the specified user belongs to, along with their role in each workspace.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

		req := users.NewGetAuthAPIV1UserUserIDWorkspaceMembershipParams()

		req.SetUserID(flags.ReadStringFlag(cmd, "UserID"))

		resp, err := client.Users.GetAuthAPIV1UserUserIDWorkspaceMembership(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-auth-api-v1-user-user-id-workspace-membership", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_user_user_id_workspace_membership] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1UsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: `Retrieves a list of users who are members of the workspace.`,
	Long:  `Retrieves a list of users who are members of the workspace.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

		req := users.NewGetAuthAPIV1UsersParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgAuthAPIGetUsersRequest{
			Email:         flags.ReadStringFlag(cmd, "Email"),
			EmailVerified: flags.ReadBooleanFlag(cmd, "EmailVerified"),
			RoleName:      models.GithubComKaytuIoKaytuEnginePkgAuthAPIRole(flags.ReadStringFlag(cmd, "RoleName")),
		})

		resp, err := client.Users.GetAuthAPIV1Users(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-auth-api-v1-users", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_users] : %v", err)
		}

		return nil
	},
}

var GetAuthApiV1WorkspaceRoleBindingsCmd = &cobra.Command{
	Use:   "list-workspace-role-bindings",
	Short: `Get all the RoleBindings of the workspace. RoleBinding defines the roles and actions a user can perform. There are currently three roles (admin, editor, viewer). The workspace path is based on the DNS such as (workspace1.app.keibi.io)`,
	Long:  `Get all the RoleBindings of the workspace. RoleBinding defines the roles and actions a user can perform. There are currently three roles (admin, editor, viewer). The workspace path is based on the DNS such as (workspace1.app.keibi.io)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

		req := users.NewGetAuthAPIV1WorkspaceRoleBindingsParams()

		resp, err := client.Users.GetAuthAPIV1WorkspaceRoleBindings(req, auth)
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-auth-api-v1-workspace-role-bindings", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_auth_api_v_1_workspace_role_bindings] : %v", err)
		}

		return nil
	},
}

var PutAuthApiV1UserRoleBindingCmd = &cobra.Command{
	Use:   "update-user-role-binding",
	Short: `Updates the role of a user in the workspace.`,
	Long:  `Updates the role of a user in the workspace.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

		req := users.NewPutAuthAPIV1UserRoleBindingParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgAuthAPIPutRoleBindingRequest{
			RoleName: models.GithubComKaytuIoKaytuEnginePkgAuthAPIRole(flags.ReadStringFlag(cmd, "RoleName")),
			UserID:   flags.ReadStringOptionalFlag(cmd, "UserID"),
		})

		_, err = client.Users.PutAuthAPIV1UserRoleBinding(req, auth)
		if err != nil {
			return fmt.Errorf("[put_auth_api_v_1_user_role_binding] : %v", err)
		}

		return nil
	},
}
