package onboard

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

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
