package blocker_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	mock_blocker_service "github.com/Doremi203/couply/backend/blocker/internal/mocks/usecase/blocker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_GetBlockInfo(t *testing.T) {
	t.Parallel()

	type mocks struct {
		userServiceClient    *mock_blocker_service.MockuserClient
		bot                  *mock_blocker_service.MockbotClient
		blockerStorageFacade *mock_blocker_service.MockblockerStorageFacade
	}
	type args struct {
		token token.Token
		in    *dto.GetBlockInfoV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.GetBlockInfoV1Response
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
				m.blockerStorageFacade.EXPECT().GetBlockInfoTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, blocker.ErrUserBlockNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetBlockInfoV1Request{},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, blocker.ErrUserBlockNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.blockerStorageFacade.EXPECT().GetBlockInfoTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(
						&blocker.UserBlock{
							BlockedID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						}, nil,
					)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetBlockInfoV1Request{},
			},
			want: &dto.GetBlockInfoV1Response{
				BlockedUserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mocks := mocks{
				userServiceClient:    mock_blocker_service.NewMockuserClient(ctrl),
				bot:                  mock_blocker_service.NewMockbotClient(ctrl),
				blockerStorageFacade: mock_blocker_service.NewMockblockerStorageFacade(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			logger := &loggerStub{}

			usecase := NewUseCase(
				mocks.userServiceClient,
				mocks.bot,
				mocks.blockerStorageFacade,
				logger,
			)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.GetBlockInfo(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
