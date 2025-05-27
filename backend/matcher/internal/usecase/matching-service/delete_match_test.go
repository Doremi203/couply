package matching_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/gen/api/messages"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	mock_matching_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/matching"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type clientWriterStub struct {
	sendMessageFunc func(ctx context.Context, data *messages.MatcherEvent) error
}

func NewClientWriterStub() *clientWriterStub {
	return &clientWriterStub{
		sendMessageFunc: func(ctx context.Context, data *messages.MatcherEvent) error {
			return nil // default behavior - no error
		},
	}
}

func (c *clientWriterStub) WithSendMessageError(err error) *clientWriterStub {
	c.sendMessageFunc = func(ctx context.Context, data *messages.MatcherEvent) error {
		return err
	}
	return c
}

func (c *clientWriterStub) WithSendMessageFunc(fn func(ctx context.Context, data *messages.MatcherEvent) error) *clientWriterStub {
	c.sendMessageFunc = fn
	return c
}

func (c *clientWriterStub) SendMessage(ctx context.Context, data *messages.MatcherEvent) error {
	return c.sendMessageFunc(ctx, data)
}

func TestUseCase_DeleteMatch(t *testing.T) {
	t.Parallel()

	type mocks struct {
		matchingStorageFacade *mock_matching_service.MockmatchingStorageFacade
	}
	type args struct {
		token token.Token
		in    *dto.DeleteMatchV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.DeleteMatchV1Response
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "token error",
			tokenErr: true,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, token.ErrTokenNotFound)
			},
		},
		{
			name: "tx error",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().DeleteMatchTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112")).
					Return(matching.ErrMatchNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.DeleteMatchV1Request{
					TargetUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, matching.ErrMatchNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().DeleteMatchTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112")).
					Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.DeleteMatchV1Request{
					TargetUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     &dto.DeleteMatchV1Response{},
			wantErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mocks := mocks{
				matchingStorageFacade: mock_matching_service.NewMockmatchingStorageFacade(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			sqsClientWriterStub := NewClientWriterStub()

			usecase := NewUseCase(mocks.matchingStorageFacade, sqsClientWriterStub)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.DeleteMatch(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
