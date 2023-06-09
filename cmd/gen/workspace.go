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
		GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd)

		GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspaceWorkspaceIdCmd)

		GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd)

		GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd)

		GetWorkspaceCmd.AddCommand(GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd)

		GetWorkspaceCmd.AddCommand(DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd)

		GetWorkspaceCmd.AddCommand(PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd)
}