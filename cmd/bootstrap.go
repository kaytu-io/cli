package cmd

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/kaytu-io/cli-program/cmd/onboard"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use: "bootstrap",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var awsCmd = &cobra.Command{
	Use: "aws",
	RunE: func(cmd *cobra.Command, args []string) error {
		//config, err := pkg.GetConfig(cmd, false)
		//if err != nil {
		//	return err
		//}
		//kaytuClient, auth := kaytu.GetKaytuAuthClientWithConfig("kaytu", config.AccessToken)
		//if err != nil {
		//	if errors.Is(err, pkg.ExpiredSession) {
		//		fmt.Println(err.Error())
		//		return nil
		//	}
		//	return err
		//}
		//
		//resp, err := kaytuClient.Workspace.GetWorkspaceAPIV1Workspaces(workspace.NewGetWorkspaceAPIV1WorkspacesParams(), auth)
		//if err != nil {
		//	return fmt.Errorf("[bootstrap-aws]: %v", err)
		//}
		//response := resp.GetPayload()
		//var workspaceName string
		//if len(response) > 1 {
		//	var items []string
		//	for _, r := range response {
		//		items = append(items, r.Name)
		//	}
		//	prompt := promptui.Select{
		//		Label: "Please select the bootstrapping workspace",
		//		Items: items,
		//	}
		//	_, result, err := prompt.Run()
		//	if err != nil {
		//		return fmt.Errorf("[bootstrap-aws]: %v", err)
		//	}
		//	workspaceName = result
		//} else if len(response) == 0 {
		//	return fmt.Errorf("no workspace found")
		//} else {
		//	workspaceName = response[0].Name
		//}
		//_ = workspaceName

		var cfg aws.Config
		cfg, err := onboard.GetConfig(context.Background(), "", "", "", "", nil)
		if err != nil {
			return err
		}

		isMaster, err := onboard.AmITheMaster(cfg)
		if err != nil {
			return err
		}

		if isMaster {
			err = onboard.CreateStackForMaster(cfg)
			if err != nil {
				return err
			}
		} else {
			err = onboard.CreateStackForSingleAccount(cfg)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var kaytuCmd = &cobra.Command{
	Use: "kaytu",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := CreateManagement("517592840862", "ws-test6")
		if err != nil {
			return err
		}
		return nil
	},
}

func CreateManagement(accountID, workspaceID string) error {
	userName := fmt.Sprintf("jump-%s", workspaceID)
	policyName := fmt.Sprintf("policy-jump-%s", workspaceID)
	managementPolicyName := fmt.Sprintf("management-policy-jump-%s", workspaceID)
	roleName := fmt.Sprintf("management-role-%s", workspaceID)

	var cfg aws.Config
	cfg, err := onboard.GetConfig(context.Background(), "", "", "", "", nil)
	if err != nil {
		return err
	}

	ctx := context.Background()
	iamClient := iam.NewFromConfig(cfg)

	user, err := iamClient.CreateUser(ctx, &iam.CreateUserInput{UserName: aws.String(userName)})
	if err != nil {
		return err
	}
	_ = user

	assumePolicy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "sts:AssumeRole",
                "iam:GetUser",
                "iam:ListAttachedUserPolicies",
                "iam:ListAccessKeys"
            ],
            "Resource": "*"
        }
    ]
}`

	var v interface{}
	err = json.Unmarshal([]byte(assumePolicy), &v)
	if err != nil {
		return err
	}
	fmt.Println("===")
	policy, err := iamClient.CreatePolicy(ctx, &iam.CreatePolicyInput{
		PolicyDocument: aws.String(assumePolicy),
		PolicyName:     aws.String(policyName),
	})
	if err != nil {
		return err
	}

	attachment, err := iamClient.AttachUserPolicy(ctx, &iam.AttachUserPolicyInput{PolicyArn: policy.Policy.Arn, UserName: aws.String(userName)})
	if err != nil {
		return err
	}
	fmt.Println(attachment.ResultMetadata)

	managementPolicyStr := `{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": "sts:AssumeRole",
    "Resource": "arn:aws:iam::*:role/kaytu-*"
  }
}`
	fmt.Println("===")

	managementPolicy, err := iamClient.CreatePolicy(ctx, &iam.CreatePolicyInput{
		PolicyDocument: aws.String(managementPolicyStr),
		PolicyName:     aws.String(managementPolicyName),
	})
	if err != nil {
		return err
	}

	trustPolicy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::%s:root"
            },
            "Action": [
                "sts:AssumeRole",
                "sts:TagSession"
            ]
        }
    ]
}`
	trustPolicy = fmt.Sprintf(trustPolicy, accountID)
	role, err := iamClient.CreateRole(ctx, &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(trustPolicy),
		RoleName:                 aws.String(roleName),
	})
	if err != nil {
		return err
	}
	_ = role

	_, err = iamClient.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{PolicyArn: managementPolicy.Policy.Arn, RoleName: aws.String(roleName)})
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
	bootstrapCmd.AddCommand(awsCmd)
	bootstrapCmd.AddCommand(kaytuCmd)
}
