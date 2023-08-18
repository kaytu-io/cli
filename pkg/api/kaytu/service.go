package kaytu

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kaytu-io/cli-program/pkg"
	apiclient "github.com/kaytu-io/cli-program/pkg/api/kaytu/client"
	"github.com/spf13/cobra"
)

func GetKaytuAuthClient(cmd *cobra.Command) (*apiclient.KaytuServiceAPI, runtime.ClientAuthInfoWriter, error) {
	cnf, err := pkg.GetConfig(cmd, false)
	if err != nil {
		return nil, nil, err
	}

	client, bearerTokenAuth := GetKaytuAuthClientWithConfig(cnf.DefaultWorkspace, cnf.AccessToken)
	return client, bearerTokenAuth, err
}

func GetKaytuAuthClientWithConfig(workspace, accessToken string) (*apiclient.KaytuServiceAPI, runtime.ClientAuthInfoWriter) {
	client := apiclient.New(httptransport.New(pkg.KaytuHostname, "/"+workspace, []string{"https"}), strfmt.Default)
	bearerTokenAuth := httptransport.BearerToken(accessToken)
	return client, bearerTokenAuth
}
