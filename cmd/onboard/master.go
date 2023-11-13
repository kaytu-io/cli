package onboard

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func AmITheMaster(cfg aws.Config) (bool, error) {
	ctx := context.Background()

	orgClient := organizations.NewFromConfig(cfg)
	out, err := orgClient.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	if err != nil {
		return false, fmt.Errorf("failure during describing organization: %v", err)
	}

	stsClient := sts.NewFromConfig(cfg)
	caller, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return false, fmt.Errorf("failure during getting caller identity: %v", err)
	}

	callerAccount := *caller.Account
	masterAccount := *out.Organization.MasterAccountId
	return callerAccount == masterAccount, nil
}
