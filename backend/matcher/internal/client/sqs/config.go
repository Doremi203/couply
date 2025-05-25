package sqs

type SQSConfig struct {
	Endpoint         string `secret:"sqs-endpoint"`
	AccessKey        string `secret:"sqs-access-key"`
	SecretKey        string `secret:"sqs-secret-key"`
	Region           string `secret:"sqs-region"`
	MatchingQueueURL string `secret:"sqs-matching-queue-url"`
}
