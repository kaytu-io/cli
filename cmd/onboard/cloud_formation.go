package onboard

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"strings"
	"time"
)

//go:embed template.yaml
var templateBody string

func CreateStackForMaster(cfg aws.Config) error {
	templateBody = strings.ReplaceAll(templateBody, "${.ManagementAccountID}", "517592840862")
	templateBody = strings.ReplaceAll(templateBody, "${.WorkspaceID}", "ws-test7")

	client := cloudformation.NewFromConfig(cfg)
	stacks, err := client.DescribeStackSet(context.Background(), &cloudformation.DescribeStackSetInput{
		StackSetName: aws.String("Kaytu-Org-Deploy"),
	})
	if err != nil {
		if !strings.Contains(err.Error(), "StackSetNotFoundException") {
			return fmt.Errorf("failure during describing stack set: %v", err)
		}
		stacks = &cloudformation.DescribeStackSetOutput{}
	}
	if stacks.StackSet != nil && stacks.StackSet.Status == types.StackSetStatusActive {
		return nil
	}

	rootID, err := GetRootID(cfg)
	if err != nil {
		return fmt.Errorf("failure during getting root id: %v", err)
	}

	//TODO-Saleh get the template from backend

	_, err = client.CreateStackSet(context.Background(), &cloudformation.CreateStackSetInput{
		StackSetName: aws.String("Kaytu-Org-Deploy"),
		Capabilities: []types.Capability{types.CapabilityCapabilityNamedIam},
		Parameters: []types.Parameter{
			{
				ParameterKey:   aws.String("OrganizationUnitList"),
				ParameterValue: aws.String(rootID),
			},
		},
		TemplateBody: aws.String(templateBody),
	})
	if err != nil {
		return fmt.Errorf("failure during creating stack set: %v", err)
	}

	fmt.Print("Waiting for StackSet status to be active .")
	for i := 0; i < 5; i++ {
		stacks, err := client.DescribeStackSet(context.Background(), &cloudformation.DescribeStackSetInput{
			StackSetName: aws.String("Kaytu-Org-Deploy"),
		})
		if err != nil {
			return fmt.Errorf("failure during checking stack set status: %v", err)
		}
		if stacks.StackSet != nil && stacks.StackSet.Status == types.StackSetStatusActive {
			return nil
		}
		fmt.Print(".", stacks.StackSet.Status)
		time.Sleep(5 * time.Second)
	}
	fmt.Println()
	return nil
}

func CreateStackForSingleAccount(cfg aws.Config) error {
	client := cloudformation.NewFromConfig(cfg)
	stacks, err := client.DescribeStacks(context.Background(), &cloudformation.DescribeStacksInput{
		StackName: aws.String("Kaytu-Single-Deploy"),
	})
	if err != nil {
		if !strings.Contains(err.Error(), "does not exist") {
			return fmt.Errorf("failure during describing stack: %v", err)
		}
		stacks = &cloudformation.DescribeStacksOutput{}
	}

	if len(stacks.Stacks) > 0 {
		return nil
	}

	//TODO-Saleh get the template from backend

	_, err = client.CreateStack(context.Background(), &cloudformation.CreateStackInput{
		StackName:    aws.String("Kaytu-Single-Deploy"),
		Capabilities: []types.Capability{types.CapabilityCapabilityNamedIam},
		TemplateBody: aws.String(templateBody),
	})
	if err != nil {
		return fmt.Errorf("failure during creating stack: %v", err)
	}

	fmt.Print("Waiting for Stack status to be active .")
	for i := 0; i < 5; i++ {
		stacks, err := client.DescribeStacks(context.Background(), &cloudformation.DescribeStacksInput{
			StackName: aws.String("Kaytu-Single-Deploy"),
		})
		if err != nil {
			return fmt.Errorf("failure during checking stack status: %v", err)
		}

		for _, stack := range stacks.Stacks {
			fmt.Print(".", stack.StackStatus)
			if stack.StackStatus == types.StackStatusCreateComplete {
				return nil
			}
		}

		time.Sleep(5 * time.Second)
	}
	fmt.Println()
	return nil
}
