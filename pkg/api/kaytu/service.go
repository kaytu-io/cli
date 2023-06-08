package kaytu

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kaytu-io/cli-program/pkg"
	apiclient "github.com/kaytu-io/cli-program/pkg/api/kaytu/client"
	"github.com/spf13/cobra"
)

func GetKaytuClient(cmd *cobra.Command) (*apiclient.KeibiServiceAPI, func(r *runtime.ClientOperation), error) {
	cnf, err := pkg.GetConfig(cmd, false)
	if err != nil {
		return nil, nil, err
	}

	client := apiclient.New(httptransport.New(pkg.KaytuHostname, "/"+cnf.DefaultWorkspace, []string{"https"}), strfmt.Default)
	bearerTokenAuth := httptransport.BearerToken(cnf.AccessToken)
	opt := func(r *runtime.ClientOperation) {
		r.AuthInfo = bearerTokenAuth
	}

	return client, opt, err
}

func GetKaytuAuthClient(cmd *cobra.Command) (*apiclient.KeibiServiceAPI, runtime.ClientAuthInfoWriter, error) {
	cnf, err := pkg.GetConfig(cmd, false)
	if err != nil {
		return nil, nil, err
	}

	client := apiclient.New(httptransport.New(pkg.KaytuHostname, "/"+cnf.DefaultWorkspace, []string{"https"}), strfmt.Default)
	bearerTokenAuth := httptransport.BearerToken(cnf.AccessToken)
	return client, bearerTokenAuth, err
}
