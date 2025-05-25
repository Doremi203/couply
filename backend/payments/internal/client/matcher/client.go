package matcher

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client userservicegrpc.UserServiceClient
}

func NewClient(addr string) (*Client, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.Wrap(err, "grpc.NewClient")
	}
	return &Client{
		conn:   conn,
		client: userservicegrpc.NewUserServiceClient(conn),
	}, conn, nil
}
