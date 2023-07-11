// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var WorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceCmd)
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("description", "", "")
	PostWorkspaceApiV1WorkspaceCmd.MarkFlagRequired("description")
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("name", "", "")
	PostWorkspaceApiV1WorkspaceCmd.MarkFlagRequired("name")
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("tier", "", "")
	PostWorkspaceApiV1WorkspaceCmd.MarkFlagRequired("tier")

	WorkspaceCmd.AddCommand(DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd)
	DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd.Flags().String("workspace-id", "", "Workspace ID")
	DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspaceCurrentCmd)

	WorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesWorkspaceIdCmd.Flags().String("workspace-id", "", "Workspace ID")
	GetWorkspaceApiV1WorkspacesWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd.Flags().String("workspace-id", "", "Workspace ID")
	GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd)
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.Flags().Bool("ignore-usage", false, "Ignore usage")
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.Flags().String("workspace-name", "", "Workspace Name")
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.MarkFlagRequired("workspace-name")

	WorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd.Flags().String("workspace-id", "", "Workspace ID")
	GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesCmd)

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd.Flags().String("workspace-id", "", "Workspace ID")
	PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd.Flags().String("workspace-id", "", "Workspace ID")
	PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.Flags().String("new-name", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.MarkFlagRequired("new-name")

	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.Flags().String("workspace-id", "", "WorkspaceID")
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.Flags().Int64("new-org-id", 0, "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.MarkFlagRequired("new-org-id")

	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.Flags().String("workspace-id", "", "WorkspaceID")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.Flags().String("new-owner-user-id", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.MarkFlagRequired("new-owner-user-id")

	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.Flags().String("workspace-id", "", "WorkspaceID")
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.MarkFlagRequired("workspace-id")

	WorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.Flags().String("new-name", "", "")
	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.MarkFlagRequired("new-name")

	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.Flags().String("workspace-id", "", "WorkspaceID")
	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.MarkFlagRequired("workspace-id")

}
