package internalaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
  awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

type Config struct {
  URL string
  Region string
}

type customEndpointResolver struct{
  config *Config
}

func (r customEndpointResolver) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
  if service == sns.ServiceID && region == r.config.Region {
    return aws.Endpoint{
      URL: r.config.URL,
      SigningRegion: r.config.Region,
    }, nil
  }
  return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
}

func New(config *Config) (*sns.Client, error) {
  cfg, err := awsconfig.LoadDefaultConfig(
    context.TODO(),
    awsconfig.WithRegion(config.Region),
    awsconfig.WithEndpointResolverWithOptions(customEndpointResolver{
      config: config,
    }),
  )

  if err != nil {
    return nil, err
  }

  return sns.NewFromConfig(cfg), nil
}
