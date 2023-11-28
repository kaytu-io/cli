package cmd

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/gen"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

//go:embed onboard/template.yaml
var templateBody string

func GetConfig(ctx context.Context, awsAccessKey, awsSecretKey, awsSessionToken, assumeRoleArn, profile string, externalId *string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}

	if awsAccessKey != "" {
		opts = append(opts, config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, awsSessionToken)))
	}

	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
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
			context.Background(),
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

func CreateStack(cfg aws.Config, userARN, handshakeID string) (string, error) {
	client := cloudformation.NewFromConfig(cfg)
	stackName := "Kaytu-Deploy"
	_, err := client.CreateStack(context.Background(), &cloudformation.CreateStackInput{
		StackName: aws.String(stackName),
		Parameters: []types.Parameter{
			{
				ParameterKey:   aws.String("KaytuManagementIAMUser"),
				ParameterValue: aws.String(userARN),
			},
			{
				ParameterKey:   aws.String("KaytuHandshakeID"),
				ParameterValue: aws.String(handshakeID),
			},
		},
		Capabilities: []types.Capability{types.CapabilityCapabilityNamedIam},
		TemplateBody: aws.String(templateBody),
	})
	if err != nil {
		return "", err
	}

	for {
		stacks, err := client.DescribeStacks(context.Background(), &cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		})
		if err != nil {
			return "", err
		}

		if len(stacks.Stacks) > 0 {
			var roleName string
			for _, output := range stacks.Stacks[0].Outputs {
				if *output.OutputKey == "KaytuReadOnly" {
					roleName = *output.OutputValue
				}
			}
			if roleName != "" {
				return roleName, nil
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func GetRootID(cfg aws.Config) (string, error) {
	orgClient := organizations.NewFromConfig(cfg)
	roots, err := orgClient.ListRoots(context.Background(), &organizations.ListRootsInput{})
	if err != nil {
		return "", err
	}

	if len(roots.Roots) > 1 {
		return "", errors.New("multiple roots detected")
	}

	if len(roots.Roots) == 0 {
		return "", errors.New("no root found")
	}

	return *roots.Roots[0].Id, nil
}

func CreateStackSet(cfg aws.Config, userARN, handshakeID string) error {
	rootID, err := GetRootID(cfg)
	if err != nil {
		return err
	}

	client := cloudformation.NewFromConfig(cfg)
	stackName := "KaytuEnrollOrganization"
	_, err = client.CreateStackSet(context.Background(), &cloudformation.CreateStackSetInput{
		StackSetName: aws.String(stackName),
		AutoDeployment: &types.AutoDeployment{
			Enabled:                      aws.Bool(true),
			RetainStacksOnAccountRemoval: aws.Bool(true),
		},
		PermissionModel: types.PermissionModelsServiceManaged,
		Parameters: []types.Parameter{
			{
				ParameterKey:   aws.String("KaytuManagementIAMUser"),
				ParameterValue: aws.String(userARN),
			},
			{
				ParameterKey:   aws.String("KaytuHandshakeID"),
				ParameterValue: aws.String(handshakeID),
			},
		},
		Capabilities: []types.Capability{types.CapabilityCapabilityNamedIam},
		TemplateBody: aws.String(templateBody),
	})
	if err != nil {
		return err
	}

	_, err = client.CreateStackInstances(context.Background(), &cloudformation.CreateStackInstancesInput{
		Regions:      []string{"us-east-1"},
		StackSetName: aws.String(stackName),
		DeploymentTargets: &types.DeploymentTargets{
			OrganizationalUnitIds: []string{rootID},
		},
		OperationPreferences: &types.StackSetOperationPreferences{
			MaxConcurrentCount:    aws.Int32(10),
			FailureToleranceCount: aws.Int32(9),
		},
	})
	if err != nil {
		return err
	}

	for {
		instances, err := client.ListStackInstances(context.Background(), &cloudformation.ListStackInstancesInput{
			StackSetName: aws.String(stackName),
		})
		if err != nil {
			return err
		}

		if len(instances.Summaries) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		isRunning := false
		for _, instance := range instances.Summaries {
			if instance.StackInstanceStatus.DetailedStatus == types.StackInstanceDetailedStatusPending ||
				instance.StackInstanceStatus.DetailedStatus == types.StackInstanceDetailedStatusRunning {
				isRunning = true
			}
		}

		if isRunning {
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	// one success
	return nil
}

func CheckAccessToMasterAccount(cfg aws.Config) (bool, error) {
	orgClient := organizations.NewFromConfig(cfg)
	out, err := orgClient.DescribeOrganization(context.Background(), &organizations.DescribeOrganizationInput{})
	if err != nil {
		if strings.Contains(err.Error(), "AWSOrganizationsNotInUseException") {
			return false, nil
		}
		return false, err
	}

	stsClient := sts.NewFromConfig(cfg)
	caller, err := stsClient.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		return false, err
	}

	callerAccount := *caller.Account
	masterAccount := *out.Organization.MasterAccountId
	return callerAccount == masterAccount, nil
}

var awsCmd = &cobra.Command{
	Use: "aws",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := pkg.GetConfig(cmd, false)
		if err != nil {
			return err
		}
		kaytuClient, auth := kaytu.GetKaytuAuthClientWithConfig("kaytu", config.AccessToken)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return err
		}

		resp, err := kaytuClient.Workspace.GetWorkspaceAPIV1Workspaces(workspace.NewGetWorkspaceAPIV1WorkspacesParams(), auth)
		if err != nil {
			return fmt.Errorf("[bootstrap-aws]: %v", err)
		}
		response := resp.GetPayload()
		var workspaceName string
		var ws *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceResponse
		var items []string
		for _, r := range response {
			if r.Status == models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusBOOTSTRAPPING {
				items = append(items, r.Name)
			}
		}

		if len(items) > 1 {
			prompt := promptui.Select{
				Label: "Please select the workspace you want to onboard into",
				Items: items,
			}
			_, result, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("[bootstrap-aws]: %v", err)
			}
			workspaceName = result
		} else if len(response) == 0 {
			return fmt.Errorf("no workspace found")
		} else {
			workspaceName = response[0].Name
		}
		for _, r := range response {
			if r.Name == workspaceName {
				ws = r
			}
		}

		fmt.Println("* Checking IAM Login")
		var cfg aws.Config
		cfg, err = GetConfig(context.Background(), "", "", "", "", cmd.Flag("profile").Value.String(), nil)
		if err != nil {
			return err
		}

		fmt.Println("* Checking Access")
		isMaster, err := CheckAccessToMasterAccount(cfg)
		if err != nil {
			return err
		}

		fmt.Println("* Configuring Organization Master Account")
		arn, err := CreateStack(cfg, ws.AwsUserArn, ws.AwsUniqueID)
		if err != nil {
			return err
		}
		fmt.Println("* Completed configuration of Organization Master Account")

		arnPart1, roleName, ok := strings.Cut(arn, "/")
		if !ok {
			return errors.New("invalid arn: " + arn)
		}
		parts := strings.Split(arnPart1, ":")
		accountID := parts[4]

		if isMaster {
			fmt.Println("* Configuring Organization Members Account")
			err := CreateStackSet(cfg, ws.AwsUserArn, ws.AwsUniqueID)
			if err != nil {
				return err
			}
			fmt.Println("* Completed configuration of Organization Master Accounts")
		}

		fmt.Println("* The Role ARN:", arn)
		for i := 0; i < 6; i++ {
			time.Sleep(5 * time.Second)

			fmt.Println("* Onboarding into Kaytu")

			req := workspace.NewPostWorkspaceAPIV1BootstrapWorkspaceNameCredentialParams()
			cnf := models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2AWSCredentialV2Config{
				AccountID:           accountID,
				AssumeRoleName:      roleName,
				ExternalID:          ws.AwsUniqueID,
				HealthCheckPolicies: nil,
			}
			req.SetWorkspaceName(workspaceName)
			req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIAddCredentialRequest{
				AwsConfig:     &cnf,
				ConnectorType: models.SourceTypeAWS,
			})
			_, err = kaytuClient.Workspace.PostWorkspaceAPIV1BootstrapWorkspaceNameCredential(req, auth)

			if err != nil {
				if i < 5 {
					continue
				}
				return err
			}
			break
		}
		fmt.Println("* Finished :)")
		return nil
	},
}

func init() {
	gen.OnboardCmd.AddCommand(awsCmd)
	awsCmd.Flags().String("profile", "", "Specifying AWS CLI profile ")
}
