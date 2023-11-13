package onboard

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetConfig(ctx context.Context, awsAccessKey, awsSecretKey, awsSessionToken, assumeRoleArn string, externalId *string) (aws.Config, error) {
	var opts []func(*config.LoadOptions) error

	if awsAccessKey != "" {
		opts = append(opts, config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, awsSessionToken)))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load AWS config: %w", err)
	}
	if cfg.Region == "" {
		cfg.Region = "us-east-1"
	}

	if assumeRoleArn != "" {
		cfg, err = config.LoadDefaultConfig(
			ctx,
			config.WithCredentialsProvider(
				stscreds.NewAssumeRoleProvider(
					sts.NewFromConfig(cfg),
					assumeRoleArn,
					func(o *stscreds.AssumeRoleOptions) {
						o.ExternalID = externalId
					},
				),
			),
		)
		if err != nil {
			return aws.Config{}, fmt.Errorf("failed to assume role: %w", err)
		}
	}

	return cfg, nil
}
