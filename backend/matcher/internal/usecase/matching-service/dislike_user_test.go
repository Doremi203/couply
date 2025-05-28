package matching_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	mock_matching_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/matching"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_DislikeUser(t *testing.T) {
	t.Parallel()

	type mocks struct {
		matchingStorageFacade *mock_matching_service.MockmatchingStorageFacade
	}
	type args struct {
		token token.Token
		in    *dto.DislikeUserV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.DislikeUserV1Response
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
				m.matchingStorageFacade.EXPECT().UpdateLikeTx(gomock.Any(), &matching.Like{
					SenderID:   uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					ReceiverID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Message:    "",
					Status:     matching.StatusDeclined,
				}).Return(matching.ErrLikesNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.DislikeUserV1Request{
					TargetUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, matching.ErrLikesNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.matchingStorageFacade.EXPECT().UpdateLikeTx(gomock.Any(), &matching.Like{
					SenderID:   uuid.MustParse("11111111-1111-1111-1111-111111111112"),
					ReceiverID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Message:    "",
					Status:     matching.StatusDeclined,
				}).Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.DislikeUserV1Request{
					TargetUserID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     &dto.DislikeUserV1Response{},
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
			got, err := usecase.DislikeUser(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
