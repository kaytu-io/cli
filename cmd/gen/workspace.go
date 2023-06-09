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
	GetWorkspaceCmd.AddCommand(DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd)
	DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceCmd)
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("description", "", "")
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("name", "", "")
	PostWorkspaceApiV1WorkspaceCmd.Flags().String("tier", "", "")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.Flags().Int64("new-org-id", 0, "")

	PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.Flags().String("new-owner-user-id", "", "")

	PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.Flags().String("new-name", "", "")

	PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspaceWorkspaceIdCmd)
	GetWorkspaceApiV1WorkspaceWorkspaceIdCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd)
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.Flags().Bool("ignore-usage", false, "")
	GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd.Flags().String("workspace-name", "", "")

	GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesCmd)

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.Flags().String("new-name", "", "")

	PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd.Flags().String("workspace-id", "", "")

	GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd)
	PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd.Flags().String("workspace-id", "", "")
}
