// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var UsersCmd = &cobra.Command{
	Use: "users",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	UsersCmd.AddCommand(DeleteAuthApiV1UserInviteCmd)
	DeleteAuthApiV1UserInviteCmd.Flags().String("user-id", "", "")
	// DeleteAuthApiV1UserInviteCmd.MarkFlagRequired("user-id")
	DeleteAuthApiV1UserInviteCmd.MarkFlagRequired("user-id")

	UsersCmd.AddCommand(DeleteAuthApiV1UserRoleBindingCmd)
	DeleteAuthApiV1UserRoleBindingCmd.Flags().String("user-id", "", "")
	// DeleteAuthApiV1UserRoleBindingCmd.MarkFlagRequired("user-id")
	DeleteAuthApiV1UserRoleBindingCmd.MarkFlagRequired("user-id")

	UsersCmd.AddCommand(GetAuthApiV1UserUserIdCmd)
	GetAuthApiV1UserUserIdCmd.Flags().String("user-id", "", "")
	// GetAuthApiV1UserUserIdCmd.MarkFlagRequired("user-id")
	GetAuthApiV1UserUserIdCmd.MarkFlagRequired("user-id")

	UsersCmd.AddCommand(PostAuthApiV1UserInviteCmd)
	PostAuthApiV1UserInviteCmd.Flags().String("email", "", "")
	PostAuthApiV1UserInviteCmd.MarkFlagRequired("email")
	PostAuthApiV1UserInviteCmd.Flags().String("role-name", "", "")

	UsersCmd.AddCommand(GetAuthApiV1UserRoleBindingsCmd)

	UsersCmd.AddCommand(GetAuthApiV1UserUserIdWorkspaceMembershipCmd)
	GetAuthApiV1UserUserIdWorkspaceMembershipCmd.Flags().String("user-id", "", "")
	// GetAuthApiV1UserUserIdWorkspaceMembershipCmd.MarkFlagRequired("user-id")
	GetAuthApiV1UserUserIdWorkspaceMembershipCmd.MarkFlagRequired("user-id")

	UsersCmd.AddCommand(GetAuthApiV1UsersCmd)
	GetAuthApiV1UsersCmd.Flags().String("email", "", "")
	// GetAuthApiV1UsersCmd.MarkFlagRequired("email")
	GetAuthApiV1UsersCmd.Flags().Bool("email-verified", false, "")
	// GetAuthApiV1UsersCmd.MarkFlagRequired("email-verified")
	GetAuthApiV1UsersCmd.Flags().String("role-name", "", "")

	UsersCmd.AddCommand(GetAuthApiV1WorkspaceRoleBindingsCmd)

	UsersCmd.AddCommand(PutAuthApiV1UserRoleBindingCmd)
	PutAuthApiV1UserRoleBindingCmd.Flags().String("role-name", "", "")
	PutAuthApiV1UserRoleBindingCmd.MarkFlagRequired("role-name")
	PutAuthApiV1UserRoleBindingCmd.Flags().String("user-id", "", "")
	PutAuthApiV1UserRoleBindingCmd.MarkFlagRequired("user-id")

}
