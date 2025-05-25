package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Messenger interface {
	String() string
}

func (c *Client) SendMessageToMatchingQueue(messageBody Messenger) (*sqs.SendMessageOutput, error) {
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody.String()),
		QueueUrl:    aws.String(c.config.MatchingQueueURL),
	}

	return c.client.SendMessage(input)
}
