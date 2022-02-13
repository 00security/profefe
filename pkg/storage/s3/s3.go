package s3

import (
	"context"
	"flag"

	"github.com/00security/profefe/pkg/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	defaultRegion = "us-east-1"
)

type Config struct {
	EndpointURL string
	DisableSSL  bool
	Region      string
	Bucket      string
}

func (conf *Config) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&conf.EndpointURL, "s3.endpoint-url", "", "override default URL to s3 service")
	f.BoolVar(&conf.DisableSSL, "s3.disable-ssl", false, "disable SSL verification")
	f.StringVar(&conf.Region, "s3.region", defaultRegion, "object storage region")
	f.StringVar(&conf.Bucket, "s3.bucket", "", "s3 bucket profile destination")
}

func (conf *Config) ResolveEndpoint(service string, region string, options ...interface{}) (aws.Endpoint, error) {
	return aws.Endpoint{
		PartitionID:       "aws",
		URL:               conf.EndpointURL, // or where ever you ran minio
		SigningRegion:     conf.Region,
		HostnameImmutable: true,
	}, nil
}

func (conf *Config) CreateStorage(logger *log.Logger) (*Storage, error) {
	// TODO add a retryer
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(conf.Region))
	if err != nil {
		return nil, err
	}

	if conf.EndpointURL != "" {
		// should one use custom object storage service (e.g. Minio), path-style addressing needs to be set
		cfg.EndpointResolverWithOptions = conf
	}

	return NewStorage(logger, s3.NewFromConfig(cfg), conf.Bucket), nil
}
