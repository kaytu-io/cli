// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package workspace

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var NameCmd = &cobra.Command{
	Use: "name",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := workspace.NewPostWorkspaceAPIV1WorkspaceWorkspaceIDNameParams()

		var config *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceNameRequest
		flags.ReadStructFlag(cmd, "Request", &config)
		req.SetRequest(config)

		req.SetWorkspaceID(flags.ReadStringFlag(cmd, "WorkspaceID"))

		resp, err := client.Workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDName(req, auth)
		if err != nil {
			return fmt.Errorf("[PostWorkspaceAPIV1WorkspaceWorkspaceIDName] : %v", err)
		}

		_ = resp

		return nil

	},
}

func init() {

	NameCmd.Flags().String("request", "", "Change name request")

	NameCmd.MarkFlagRequired("request")

	NameCmd.Flags().String("workspace-id", "", "WorkspaceID")

	NameCmd.MarkFlagRequired("workspace-id")

}