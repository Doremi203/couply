package sqs

type Config struct {
	Endpoint  string `secret:"sqs-endpoint"`
	AccessKey string `secret:"sqs-access-key"`
	SecretKey string `secret:"sqs-secret-key"`
	Region    string `secret:"sqs-region"`
	QueueURL  string `secret:"sqs-matching-queue-url"`
}
