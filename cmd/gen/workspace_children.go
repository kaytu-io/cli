package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/spf13/cobra"
)
var PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace_workspace_id_owner",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_owner] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOwner(workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_owner] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_owner] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace_workspace_id_resume",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_resume] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDResume(workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_resume] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_resume] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspaceWorkspaceIdCmd = &cobra.Command{
	Use: "get_workspace_api_v_1_workspace_workspace_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspaceWorkspaceID(workspace.NewGetWorkspaceAPIV1WorkspaceWorkspaceIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd = &cobra.Command{
	Use: "get_workspace_api_v_1_workspaces_byid_workspace_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesByidWorkspaceID(workspace.NewGetWorkspaceAPIV1WorkspacesByidWorkspaceIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd = &cobra.Command{
	Use: "get_workspace_api_v_1_workspaces_limits_byid_workspace_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID(workspace.NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd = &cobra.Command{
	Use: "get_workspace_api_v_1_workspaces_limits_workspace_name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesLimitsWorkspaceName(workspace.NewGetWorkspaceAPIV1WorkspacesLimitsWorkspaceNameParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspacesCmd = &cobra.Command{
	Use: "get_workspace_api_v_1_workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		resp, err := client.Workspace.GetWorkspaceAPIV1Workspaces(workspace.NewGetWorkspaceAPIV1WorkspacesParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace_workspace_id_organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_organization] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOrganization(workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOrganizationParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_organization] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_organization] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace_workspace_id_tier",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDTier(workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDTierParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		return nil
	},
}
var DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd = &cobra.Command{
	Use: "delete_workspace_api_v_1_workspace_workspace_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		resp, err := client.Workspace.DeleteWorkspaceAPIV1WorkspaceWorkspaceID(workspace.NewDeleteWorkspaceAPIV1WorkspaceWorkspaceIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1Workspace(workspace.NewPostWorkspaceAPIV1WorkspaceParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace_workspace_id_name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_name] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDName(workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDNameParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_name] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd = &cobra.Command{
	Use: "post_workspace_api_v_1_workspace_workspace_id_suspend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_suspend] : %v", err)
		}

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspend(workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_suspend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_suspend] : %v", err)
		}

		return nil
	},
}

