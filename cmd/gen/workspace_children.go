package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd = &cobra.Command{
	Use: "workspace-workspace-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		req := workspace.NewDeleteWorkspaceAPIV1WorkspaceWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.DeleteWorkspaceAPIV1WorkspaceWorkspaceID(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd = &cobra.Command{
	Use: "workspaces-byid-workspace-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesByidWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesByidWorkspaceID(req, auth)
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
	Use: "workspaces-limits-byid-workspace-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID(req, auth)
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

var GetWorkspaceApiV1WorkspacesCmd = &cobra.Command{
	Use: "workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesParams()

		resp, err := client.Workspace.GetWorkspaceAPIV1Workspaces(req, auth)
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
	Use: "workspace-workspace-id-organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_organization] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOrganizationParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest{
			NewOrgID: flags.ReadInt64Flag(cmd, "NewOrgID"),
		})
		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOrganization(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_organization] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdOwnerCmd = &cobra.Command{
	Use: "workspace-workspace-id-owner",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_owner] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOwnershipRequest{
			NewOwnerUserID: flags.ReadStringFlag(cmd, "NewOwnerUserID"),
		})
		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOwner(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_owner] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd = &cobra.Command{
	Use: "workspace-workspace-id-resume",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_resume] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDResume(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_resume] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd = &cobra.Command{
	Use: "workspace-workspace-id-tier",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDTierParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest{
			NewName: models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPITier(flags.ReadStringFlag(cmd, "NewName")),
		})
		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDTier(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspaceWorkspaceIdCmd = &cobra.Command{
	Use: "workspace-workspace-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspaceWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.GetWorkspaceAPIV1WorkspaceWorkspaceID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_workspace_id] : %v", err)
		}

		return nil
	},
}
var GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd = &cobra.Command{
	Use: "workspaces-limits-workspace-name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesLimitsWorkspaceNameParams()

		req.SetIgnoreUsage(flags.ReadBooleanOptionalFlag(cmd, "IgnoreUsage"))
		req.SetWorkspaceName(flags.ReadStringFlag(cmd, "WorkspaceName"))

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesLimitsWorkspaceName(req, auth)
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

var PostWorkspaceApiV1WorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPICreateWorkspaceRequest{
			Description: flags.ReadStringFlag(cmd, "Description"),
			Name:        flags.ReadStringFlag(cmd, "Name"),
			Tier:        flags.ReadStringFlag(cmd, "Tier"),
		})

		resp, err := client.Workspace.PostWorkspaceAPIV1Workspace(req, auth)
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
	Use: "workspace-workspace-id-name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_name] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDNameParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceNameRequest{
			NewName: flags.ReadStringFlag(cmd, "NewName"),
		})
		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDName(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_name] : %v", err)
		}

		return nil
	},
}
var PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd = &cobra.Command{
	Use: "workspace-workspace-id-suspend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_suspend] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspend(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_suspend] : %v", err)
		}

		return nil
	},
}
