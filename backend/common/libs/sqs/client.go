package sqs

import (
	"context"
	"encoding/json"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type ClientWriter[T any] interface {
	SendMessage(ctx context.Context, data T) error
}

type ClientReader[T any] interface {
	ReadMessages(ctx context.Context, logger log.Logger, maxCount int) ([]Message[T], error)
}

func New[T any](cfg Config) (*client[T], error) {
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

	return &client[T]{
		client: sqs.New(sess),
		config: cfg,
	}, nil
}

type client[T any] struct {
	client *sqs.SQS
	config Config
}

func (c *client[T]) SendMessage(ctx context.Context, data T) error {
	messageBody, err := json.Marshal(data)
	if err != nil {
		return errors.WrapFail(err, "marshal message")
	}

	input := &sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    aws.String(c.config.QueueURL),
	}

	_, err = c.client.SendMessageWithContext(ctx, input)
	if err != nil {
		return errors.WrapFailf(err, "send message to %v", errors.Token("queue_url", c.config.QueueURL))
	}

	return nil
}

func (c *client[T]) ReadMessages(ctx context.Context, logger log.Logger, maxCount int) ([]Message[T], error) {
	ret := make([]Message[T], 0, maxCount)

	resp, err := c.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(int64(maxCount)),
		QueueUrl:            aws.String(c.config.QueueURL),
	})
	if err != nil {
		return nil, errors.WrapFail(err, "recieve message from queue")
	}

	for _, msg := range resp.Messages {
		var body T
		err = json.Unmarshal([]byte(*msg.Body), &body)
		if err != nil {
			logger.Error(errors.WrapFailf(err, "unmarshal message body %v", errors.Token("message_body", *msg.Body)))
			continue
		}

		ret = append(ret, Message[T]{
			ID:            *msg.MessageId,
			ReceiptHandle: *msg.ReceiptHandle,
			Data:          body,
		})
	}

	return ret, nil
}

type Message[T any] struct {
	ID            string
	ReceiptHandle string
	Data          T
}
