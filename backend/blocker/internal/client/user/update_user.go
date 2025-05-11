package user

import (
	"context"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/metadata"
)

func (c *Client) UpdateUserV1(ctx context.Context, isBlocked bool) error {
	incomingMD, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "missing metadata in context")
	}

	tokenValues := incomingMD.Get("user-token")
	if len(tokenValues) == 0 {
		return status.Error(codes.Unauthenticated, "user-token header is required")
	}
	userToken := tokenValues[0]

	outgoingMD := metadata.New(map[string]string{
		"user-token": userToken,
	})
	ctx = metadata.NewOutgoingContext(ctx, outgoingMD)

	_, err := c.client.UpdateUserV1(ctx, &userservicegrpc.UpdateUserV1Request{
		IsBlocked: &isBlocked,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"is_blocked"},
		},
	})
	return err
}
