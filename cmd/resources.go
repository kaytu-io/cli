
package cmd

import (
	"github.com/kaytu-io/cli-program/cmd/gen"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var createCmd = &cobra.Command{
	Use: "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
var updateCmd = &cobra.Command{
	Use: "update",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	// ========== DO NOT EDIT THIS PART! AUTO GENERATED!!! =========

	getCmd.AddCommand(gen.GetWorkspaceCmd)
	updateCmd.AddCommand(gen.UpdateWorkspaceCmd)
	deleteCmd.AddCommand(gen.DeleteWorkspaceCmd)
	createCmd.AddCommand(gen.CreateWorkspaceCmd)


	getCmd.AddCommand(gen.GetUsersCmd)
	updateCmd.AddCommand(gen.UpdateUsersCmd)
	deleteCmd.AddCommand(gen.DeleteUsersCmd)
	createCmd.AddCommand(gen.CreateUsersCmd)


	getCmd.AddCommand(gen.GetStackCmd)
	updateCmd.AddCommand(gen.UpdateStackCmd)
	deleteCmd.AddCommand(gen.DeleteStackCmd)
	createCmd.AddCommand(gen.CreateStackCmd)


	getCmd.AddCommand(gen.GetSmartQueryCmd)
	updateCmd.AddCommand(gen.UpdateSmartQueryCmd)
	deleteCmd.AddCommand(gen.DeleteSmartQueryCmd)
	createCmd.AddCommand(gen.CreateSmartQueryCmd)


	getCmd.AddCommand(gen.GetServicesCmd)
	updateCmd.AddCommand(gen.UpdateServicesCmd)
	deleteCmd.AddCommand(gen.DeleteServicesCmd)
	createCmd.AddCommand(gen.CreateServicesCmd)


	getCmd.AddCommand(gen.GetScheduleCmd)
	updateCmd.AddCommand(gen.UpdateScheduleCmd)
	deleteCmd.AddCommand(gen.DeleteScheduleCmd)
	createCmd.AddCommand(gen.CreateScheduleCmd)


	getCmd.AddCommand(gen.GetRolesCmd)
	updateCmd.AddCommand(gen.UpdateRolesCmd)
	deleteCmd.AddCommand(gen.DeleteRolesCmd)
	createCmd.AddCommand(gen.CreateRolesCmd)


	getCmd.AddCommand(gen.GetResourceCmd)
	updateCmd.AddCommand(gen.UpdateResourceCmd)
	deleteCmd.AddCommand(gen.DeleteResourceCmd)
	createCmd.AddCommand(gen.CreateResourceCmd)


	getCmd.AddCommand(gen.GetOnboardCmd)
	updateCmd.AddCommand(gen.UpdateOnboardCmd)
	deleteCmd.AddCommand(gen.DeleteOnboardCmd)
	createCmd.AddCommand(gen.CreateOnboardCmd)


	getCmd.AddCommand(gen.GetMetadataCmd)
	updateCmd.AddCommand(gen.UpdateMetadataCmd)
	deleteCmd.AddCommand(gen.DeleteMetadataCmd)
	createCmd.AddCommand(gen.CreateMetadataCmd)


	getCmd.AddCommand(gen.GetLocationCmd)
	updateCmd.AddCommand(gen.UpdateLocationCmd)
	deleteCmd.AddCommand(gen.DeleteLocationCmd)
	createCmd.AddCommand(gen.CreateLocationCmd)


	getCmd.AddCommand(gen.GetKeysCmd)
	updateCmd.AddCommand(gen.UpdateKeysCmd)
	deleteCmd.AddCommand(gen.DeleteKeysCmd)
	createCmd.AddCommand(gen.CreateKeysCmd)


	getCmd.AddCommand(gen.GetInventoryCmd)
	updateCmd.AddCommand(gen.UpdateInventoryCmd)
	deleteCmd.AddCommand(gen.DeleteInventoryCmd)
	createCmd.AddCommand(gen.CreateInventoryCmd)


	getCmd.AddCommand(gen.GetInsightsCmd)
	updateCmd.AddCommand(gen.UpdateInsightsCmd)
	deleteCmd.AddCommand(gen.DeleteInsightsCmd)
	createCmd.AddCommand(gen.CreateInsightsCmd)


	getCmd.AddCommand(gen.GetDescribeCmd)
	updateCmd.AddCommand(gen.UpdateDescribeCmd)
	deleteCmd.AddCommand(gen.DeleteDescribeCmd)
	createCmd.AddCommand(gen.CreateDescribeCmd)


	getCmd.AddCommand(gen.GetConnectionsCmd)
	updateCmd.AddCommand(gen.UpdateConnectionsCmd)
	deleteCmd.AddCommand(gen.DeleteConnectionsCmd)
	createCmd.AddCommand(gen.CreateConnectionsCmd)


	getCmd.AddCommand(gen.GetConnectionCmd)
	updateCmd.AddCommand(gen.UpdateConnectionCmd)
	deleteCmd.AddCommand(gen.DeleteConnectionCmd)
	createCmd.AddCommand(gen.CreateConnectionCmd)


	getCmd.AddCommand(gen.GetComplianceCmd)
	updateCmd.AddCommand(gen.UpdateComplianceCmd)
	deleteCmd.AddCommand(gen.DeleteComplianceCmd)
	createCmd.AddCommand(gen.CreateComplianceCmd)


	getCmd.AddCommand(gen.GetBenchmarksAssignmentCmd)
	updateCmd.AddCommand(gen.UpdateBenchmarksAssignmentCmd)
	deleteCmd.AddCommand(gen.DeleteBenchmarksAssignmentCmd)
	createCmd.AddCommand(gen.CreateBenchmarksAssignmentCmd)

	// =============================================================

}
