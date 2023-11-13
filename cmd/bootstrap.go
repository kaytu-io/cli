package cmd

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	types2 "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

//go:embed onboard/template.yaml
var templateBody string

func GetConfig(ctx context.Context, awsAccessKey, awsSecretKey, awsSessionToken, assumeRoleArn string, externalId *string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}

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

func CreateStack(rootID string, cfg aws.Config) (string, string, error) {
	client := cloudformation.NewFromConfig(cfg)
	stacks, err := client.DescribeStacks(context.Background(), &cloudformation.DescribeStacksInput{
		StackName: aws.String("Kaytu-Org-Deploy"),
	})
	if err != nil {
		if !strings.Contains(err.Error(), "does not exist") {
			return "", "", err
		}
	}

	if stacks == nil || len(stacks.Stacks) == 0 {
		_, err := client.CreateStack(context.Background(), &cloudformation.CreateStackInput{
			StackName:                   aws.String("Kaytu-Org-Deploy"),
			Capabilities:                []types.Capability{types.CapabilityCapabilityNamedIam},
			ClientRequestToken:          nil,
			DisableRollback:             nil,
			EnableTerminationProtection: nil,
			NotificationARNs:            nil,
			OnFailure:                   "",
			Parameters: []types.Parameter{
				{
					ParameterKey:   aws.String("OrganizationUnitList"),
					ParameterValue: aws.String(rootID),
				},
			},
			ResourceTypes:         nil,
			RetainExceptOnCreate:  nil,
			RoleARN:               nil,
			RollbackConfiguration: nil,
			StackPolicyBody:       nil,
			StackPolicyURL:        nil,
			Tags:                  nil,
			TemplateBody:          aws.String(templateBody),
			TemplateURL:           nil,
			TimeoutInMinutes:      nil,
		})
		if err != nil {
			return "", "", err
		}
		fmt.Println("* stacks is created")
	}

	fmt.Println("* waiting for stacks to finish")
	for {
		stacks, err := client.DescribeStacks(context.Background(), &cloudformation.DescribeStacksInput{
			StackName: aws.String("Kaytu-Org-Deploy"),
		})
		if err != nil {
			return "", "", err
		}

		if len(stacks.Stacks) > 0 {
			var memberRoleName, adminRoleName string
			for _, output := range stacks.Stacks[0].Outputs {
				fmt.Println(
					*output.OutputKey,
					*output.OutputValue,
				)
				if *output.OutputKey == "MemberRoleName" {
					memberRoleName = *output.OutputValue
				}
				if *output.OutputKey == "AdminRoleArn" {
					adminRoleName = *output.OutputValue
				}
			}
			if memberRoleName != "" && adminRoleName != "" {
				return adminRoleName, memberRoleName, nil
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func CheckAccessToMasterAccount(cfg aws.Config) error {
	orgClient := organizations.NewFromConfig(cfg)
	out, err := orgClient.DescribeOrganization(context.Background(), &organizations.DescribeOrganizationInput{})
	if err != nil {
		return err
	}

	stsClient := sts.NewFromConfig(cfg)
	caller, err := stsClient.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}

	callerAccount := *caller.Account
	masterAccount := *out.Organization.MasterAccountId
	if callerAccount != masterAccount {
		return errors.New("you're not in the master account. please visit the documentations")
	}

	return nil
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

func GetUser(cfg aws.Config) (string, error) {
	iamClient := iam.NewFromConfig(cfg)
	user, err := iamClient.GetUser(context.Background(), &iam.GetUserInput{UserName: aws.String("kaytu-user")})
	if err != nil {
		if !strings.Contains(err.Error(), "cannot be found") {
			return "", err
		}
	}

	var userName string
	if user == nil || user.User == nil {
		iamUser, err := iamClient.CreateUser(context.Background(), &iam.CreateUserInput{
			UserName:            aws.String("kaytu-user"),
			Path:                nil,
			PermissionsBoundary: nil,
			Tags:                nil,
		})
		if err != nil {
			return "", err
		}
		fmt.Println("* user got created:", *iamUser.User.UserName)
		userName = *iamUser.User.UserName
	} else {
		fmt.Println("* user already exists:", *user.User.UserName)
		userName = *user.User.UserName
	}
	return userName, nil
}

func CreateAndAssignPolicies(cfg aws.Config) error {
	iamClient := iam.NewFromConfig(cfg)

	var customPolicyArn string

	policies, err := iamClient.ListPolicies(context.Background(), &iam.ListPoliciesInput{
		Scope: types2.PolicyScopeTypeLocal,
	})
	if err != nil {
		return err
	}

	for _, policy := range policies.Policies {
		if *policy.PolicyName == "AssumePolicy-Kaytu" {
			customPolicyArn = *policy.Arn
		}
	}

	if customPolicyArn == "" {
		newCustomPolicy, err := iamClient.CreatePolicy(context.Background(), &iam.CreatePolicyInput{
			PolicyDocument: aws.String(`{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Action": "sts:AssumeRole",
        "Resource": "*"
    }
}`),
			PolicyName:  aws.String("AssumePolicy-Kaytu"),
			Description: aws.String(""),
		})
		if err != nil {
			return err
		}
		customPolicyArn = *newCustomPolicy.Policy.Arn
	}
	fmt.Println("* kaytu policy created")

	policyArns := []string{
		//"arn:aws:iam::aws:policy/AWSOrganizationsReadOnlyAccess",
		//"arn:aws:iam::aws:policy/AWSBillingReadOnlyAccess",
		//"arn:aws:iam::aws:policy/AWSBudgetsReadOnlyAccess",
		//"arn:aws:iam::aws:policy/SecurityAudit",
		//"arn:aws:iam::aws:policy/ReadOnlyAccess",
		customPolicyArn,
	}

	for _, policy := range policyArns {
		fmt.Println("* attaching policy:", policy)
		_, err = iamClient.AttachUserPolicy(context.Background(), &iam.AttachUserPolicyInput{
			PolicyArn: aws.String(policy),
			UserName:  aws.String("kaytu-user"),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

var bootstrapCmd = &cobra.Command{
	Use: "bootstrap",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
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
		if len(response) > 1 {
			var items []string
			for _, r := range response {
				items = append(items, r.Name)
			}
			fmt.Println("\n")
			prompt := promptui.Select{
				Label: "Please select the bootstrapping workspace",
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

		var cfg aws.Config
		cfg, err = GetConfig(context.Background(), "", "", "", "", nil)
		if err != nil {
			return err
		}

		fmt.Println("* checking to make sure you are on the master account")
		err = CheckAccessToMasterAccount(cfg)
		if err != nil {
			return err
		}

		fmt.Println("* getting root id")
		rootID, err := GetRootID(cfg)
		if err != nil {
			return err
		}

		fmt.Println("* creating cloudformation stacks")
		adminARN, arn, err := CreateStack(rootID, cfg)
		if err != nil {
			return err
		}

		fmt.Println("* finished, got the arn:", arn)
		userName, err := GetUser(cfg)
		if err != nil {
			return err
		}

		err = CreateAndAssignPolicies(cfg)
		if err != nil {
			return err
		}

		iamClient := iam.NewFromConfig(cfg)
		fmt.Println("* creating access key")
		key, err := iamClient.CreateAccessKey(context.Background(), &iam.CreateAccessKeyInput{
			UserName: aws.String(userName),
		})
		if err != nil {
			return err
		}

		for i := 0; i < 6; i++ {
			time.Sleep(5 * time.Second)

			fmt.Println("* onboarding into kaytu")
			req := workspace.NewPostWorkspaceAPIV1BootstrapWorkspaceNameCredentialParams()
			cnf := models.GithubComKaytuIoKaytuEnginePkgOnboardAPIAWSCredentialConfig{
				AccessKey:           key.AccessKey.AccessKeyId,
				SecretKey:           key.AccessKey.SecretAccessKey,
				AssumeAdminRoleName: adminARN,
				AssumeRoleName:      arn,
			}
			req.SetWorkspaceName(workspaceName)
			req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIAddCredentialRequest{
				Config:        cnf,
				ConnectorType: models.SourceTypeAWS,
			})
			_, err = kaytuClient.Workspace.PostWorkspaceAPIV1BootstrapWorkspaceNameCredential(req, auth)

			if err != nil {
				if i < 5 {
					fmt.Printf("Failure while onboarding: %v\nRetrying in a bit ...", err)
					continue
				}
				return err
			}
			break
		}
		fmt.Println("* finished :)")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
	bootstrapCmd.AddCommand(awsCmd)
}
