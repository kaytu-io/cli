// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var PostWorkspaceApiV1WorkspaceCmd = &cobra.Command{
	Use:   "create-workspace",
	Short: `Returns workspace created`,
	Long:  `Returns workspace created`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPICreateWorkspaceRequest{
			Description: flags.ReadStringFlag(cmd, "Description"),
			Name:        flags.ReadStringFlag(cmd, "Name"),
			Tier:        flags.ReadStringFlag(cmd, "Tier"),
		})

		resp, err := client.Workspace.PostWorkspaceAPIV1Workspace(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "post-workspace-api-v1-workspace", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace] : %v", err)
		}

		return nil
	},
}

var DeleteWorkspaceApiV1WorkspaceWorkspaceIdCmd = &cobra.Command{
	Use:   "delete-workspace",
	Short: `Delete workspace with workspace id`,
	Long:  `Delete workspace with workspace id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
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

var GetWorkspaceApiV1WorkspaceCurrentCmd = &cobra.Command{
	Use:   "get-current-workspace",
	Short: `Returns all workspaces with owner id`,
	Long:  `Returns all workspaces with owner id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_workspace_api_v_1_workspace_current] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspaceCurrentParams()

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspaceCurrent(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_current] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-workspace-api-v1-workspace-current", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspace_current] : %v", err)
		}

		return nil
	},
}

var GetWorkspaceApiV1WorkspacesWorkspaceIdCmd = &cobra.Command{
	Use:   "get-workspace",
	Short: `Get workspace with workspace id`,
	Long:  `Get workspace with workspace id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_workspace_id] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.GetWorkspaceAPIV1WorkspacesWorkspaceID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_workspace_id] : %v", err)
		}

		return nil
	},
}

var GetWorkspaceApiV1WorkspacesByidWorkspaceIdCmd = &cobra.Command{
	Use:   "get-workspace-by-id",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesByidWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesByidWorkspaceID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-workspace-api-v1-workspaces-byid-workspace-id", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_byid_workspace_id] : %v", err)
		}

		return nil
	},
}

var GetWorkspaceApiV1WorkspacesLimitsWorkspaceNameCmd = &cobra.Command{
	Use:   "get-workspace-limits",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesLimitsWorkspaceNameParams()

		req.SetIgnoreUsage(flags.ReadBooleanOptionalFlag(cmd, "IgnoreUsage"))
		req.SetWorkspaceName(flags.ReadStringFlag(cmd, "WorkspaceName"))

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesLimitsWorkspaceName(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-workspace-api-v1-workspaces-limits-workspace-name", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_workspace_name] : %v", err)
		}

		return nil
	},
}

var GetWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdCmd = &cobra.Command{
	Use:   "get-workspace-limits-by-id",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams()

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		resp, err := client.Workspace.GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-workspace-api-v1-workspaces-limits-byid-workspace-id", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces_limits_byid_workspace_id] : %v", err)
		}

		return nil
	},
}

var GetWorkspaceApiV1WorkspacesCmd = &cobra.Command{
	Use:   "list-workspaces",
	Short: `Returns all workspaces with owner id`,
	Long:  `Returns all workspaces with owner id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		req := workspace.NewGetWorkspaceAPIV1WorkspacesParams()

		resp, err := client.Workspace.GetWorkspaceAPIV1Workspaces(req, auth)
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-workspace-api-v1-workspaces", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_workspace_api_v_1_workspaces] : %v", err)
		}

		return nil
	},
}

var PostWorkspaceApiV1WorkspaceWorkspaceIdResumeCmd = &cobra.Command{
	Use:   "resume-workspace",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
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

var PostWorkspaceApiV1WorkspaceWorkspaceIdSuspendCmd = &cobra.Command{
	Use:   "suspend-workspace",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
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

var PostWorkspaceApiV1WorkspaceWorkspaceIdNameCmd = &cobra.Command{
	Use:   "update-workspace-name",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_name] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDNameParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceNameRequest{
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

var PostWorkspaceApiV1WorkspaceWorkspaceIdOrganizationCmd = &cobra.Command{
	Use:   "update-workspace-organization",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_organization] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOrganizationParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest{
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
	Use:   "update-workspace-owner",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_owner] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOwnershipRequest{
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

var PostWorkspaceApiV1WorkspaceWorkspaceIdTierCmd = &cobra.Command{
	Use:   "update-workspace-tier",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDTierParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceTierRequest{
			NewName: models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPITier(flags.ReadStringFlag(cmd, "NewName")),
		})
		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		_, err = client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDTier(req, auth)
		if err != nil {
			return fmt.Errorf("[post_workspace_api_v_1_workspace_workspace_id_tier] : %v", err)
		}

		return nil
	},
}
