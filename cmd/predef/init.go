package predef

import (
	"github.com/kaytu-io/cli-program/cmd/genv2"
	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu"
	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu/get"
)

func Init() {
	kaytu.OnboardCmd.AddCommand(AWSCmd)
	get.SpendCmd.AddCommand(GetCostsCmd)
	//update.Rule.AddCommand(PUTAlertingApiV1RuleUpdateCmd)
	//kaytu.AlertingCmd.AddCommand(POSTAlertingApiV1RuleCreateCmd)

	genv2.KaytuCmd.AddCommand(AboutCmd)
	genv2.KaytuCmd.AddCommand(LoginCmd)
	genv2.KaytuCmd.AddCommand(LogoutCmd)
	genv2.KaytuCmd.AddCommand(VersionCmd)
}
