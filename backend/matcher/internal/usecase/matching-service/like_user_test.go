package matching_service

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	mock_matching_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/matching"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_LikeUser(t *testing.T) {
	t.Parallel()

	now := time.Now()

	type mocks struct {
		matchingStorageFacade *mock_matching_service.MockmatchingStorageFacade
	}
	type args struct {
		token token.Token
		in    *dto.LikeUserV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		sqsErr   bool
		want     *dto.LikeUserV1Response
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
			name: "error sqs send like",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, nil)

				m.matchingStorageFacade.EXPECT().LikeUserTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, like *matching.Like) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), like.SenderID)
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111112"), like.ReceiverID)
					assert.NotZero(t, like.CreatedAt)
					return nil
				}).Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			sqsErr:   true,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "error like user",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, nil)

				m.matchingStorageFacade.EXPECT().LikeUserTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, like *matching.Like) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), like.SenderID)
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111112"), like.ReceiverID)
					assert.NotZero(t, like.CreatedAt)
					return nil
				}).Return(errors.Error("some error"))
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "error get waiting like",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, errors.Error("some error"))
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "error sqs mutual like",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&matching.Like{
						SenderID:   uuid.MustParse("11111111-1111-1111-1111-111111111112"),
						ReceiverID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)

				m.matchingStorageFacade.EXPECT().HandleMutualLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					"").Return(&matching.Match{
					FirstUserID:  uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					SecondUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					CreatedAt:    now,
				}, nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			sqsErr:   true,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "error mutual like",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&matching.Like{
						SenderID:   uuid.MustParse("11111111-1111-1111-1111-111111111112"),
						ReceiverID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)

				m.matchingStorageFacade.EXPECT().HandleMutualLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					"").Return(nil, matching.ErrLikeNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, matching.ErrLikeNotFound)
			},
		},
		{
			name: "success mutual like",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&matching.Like{
						SenderID:   uuid.MustParse("11111111-1111-1111-1111-111111111112"),
						ReceiverID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)

				m.matchingStorageFacade.EXPECT().HandleMutualLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					"").Return(&matching.Match{
					FirstUserID:  uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					SecondUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					CreatedAt:    now,
				}, nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want: &dto.LikeUserV1Response{
				IsMatch: true,
				Match: &matching.Match{
					FirstUserID:  uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					SecondUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					CreatedAt:    now,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success new like",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().GetWaitingLikeTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, nil)

				m.matchingStorageFacade.EXPECT().LikeUserTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, like *matching.Like) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), like.SenderID)
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111112"), like.ReceiverID)
					assert.NotZero(t, like.CreatedAt)
					return nil
				}).Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.LikeUserV1Request{
					TargetUserId: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want: &dto.LikeUserV1Response{
				IsMatch: false,
				Match:   nil,
			},
			wantErr: assert.NoError,
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
			if tt.sqsErr {
				sqsClientWriterStub.WithSendMessageError(errors.Error("some error"))
			}

			usecase := NewUseCase(mocks.matchingStorageFacade, sqsClientWriterStub)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.LikeUser(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
