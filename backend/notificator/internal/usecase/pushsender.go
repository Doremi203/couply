package usecase

import "context"

type PushSender struct{}

func (u PushSender) SendBatch(ctx context.Context)
