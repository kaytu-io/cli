// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var GetWorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateWorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteWorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateWorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd)
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.Flags().Bool("ignore-usage", false, "")
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.Flags().String("workspace-name", "", "")
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.MarkFlagRequired("workspace-name")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesCmd)

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesWorkspaceIdCmd.Flags().String("workspace-id", "", "")
	GetWorkspaceApiV1WorkspacesWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd.Flags().String("workspace-id", "", "")
	GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd.Flags().String("workspace-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd.MarkFlagRequired("workspace-id")

	UpdateWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.Flags().String("new-name", "", "")

	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.Flags().String("workspace-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.MarkFlagRequired("workspace-id")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspaceCurrentCmd)

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd.Flags().String("workspace-id", "", "")
	GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	CreateWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceCmd)
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("description", "", "")
	PostWorkspaceApiV1WorkspaceCmd.MarkFlagRequired("description")
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("name", "", "")
	PostWorkspaceApiV1WorkspaceCmd.MarkFlagRequired("name")
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("tier", "", "")
	PostWorkspaceApiV1WorkspaceCmd.MarkFlagRequired("tier")

	UpdateWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.Flags().String("new-owner-user-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.MarkFlagRequired("new-owner-user-id")

	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.Flags().String("workspace-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.MarkFlagRequired("workspace-id")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd.Flags().String("workspace-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd.MarkFlagRequired("workspace-id")

	DeleteWorkspaceCmd.AddCommand(DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd)
	DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd.Flags().String("workspace-id", "", "")
	DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	UpdateWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.Flags().String("new-name", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.MarkFlagRequired("new-name")

	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.Flags().String("workspace-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.MarkFlagRequired("workspace-id")

	UpdateWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.Flags().Int64("new-org-id", 0, "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.MarkFlagRequired("new-org-id")

	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.Flags().String("workspace-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.MarkFlagRequired("workspace-id")

}
