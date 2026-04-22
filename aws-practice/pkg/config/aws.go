package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// AWSConfig holds the AWS configuration
type AWSConfig struct {
	Region string
}

// LoadAWSConfig loads the default AWS configuration
func LoadAWSConfig(ctx context.Context, region string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

// GetRegion returns the AWS region from environment or default
func GetRegion() string {
	region := "us-east-1"
	return region
}