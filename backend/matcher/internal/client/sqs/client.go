package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Client struct {
	client *sqs.SQS
	config SQSConfig
}

func New(cfg SQSConfig) (*Client, error) {
	awsConfig := &aws.Config{
		Region:   aws.String(cfg.Region),
		Endpoint: aws.String(cfg.Endpoint),
		Credentials: credentials.NewStaticCredentials(
			cfg.AccessKey,
			cfg.SecretKey,
			"",
		),
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: sqs.New(sess),
		config: cfg,
	}, nil
}
