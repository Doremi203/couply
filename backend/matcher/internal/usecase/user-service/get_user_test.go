package user_service

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	mock_user_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/user"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_GetUser(t *testing.T) {
	t.Parallel()

	now := time.Now()

	type mocks struct {
		userStorageFacade *mock_user_service.MockuserStorageFacade
		photoURLGenerator *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    *dto.GetUserV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.GetUserV1Response
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
				m.userStorageFacade.EXPECT().GetUserTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, user.ErrUserNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetUserV1Request{},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, user.ErrUserNotFound)
			},
		},
		{
			name: "generate download urls error",
			setup: func(m mocks) {
				m.userStorageFacade.EXPECT().GetUserTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&user.User{
						ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						Photos: []user.Photo{
							{
								UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
								OrderNumber: 0,
								ObjectKey:   "users/11111111-1111-1111-1111-111111111111/slot/0.jpg",
								MimeType:    ".jpg",
								UploadedAt:  &now,
								UploadURL:   func(s string) *string { return &s }("uploadURL"),
								DownloadURL: nil,
							},
						},
					}, nil)
				m.photoURLGenerator.EXPECT().GenerateDownload(gomock.Any(), "users/11111111-1111-1111-1111-111111111111/slot/0.jpg").
					Return("", errors.Error("some error"))
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetUserV1Request{},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.userStorageFacade.EXPECT().GetUserTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&user.User{
						ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetUserV1Request{},
			},
			tokenErr: false,
			want: &dto.GetUserV1Response{
				User: &user.User{
					ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mocks := mocks{
				userStorageFacade: mock_user_service.NewMockuserStorageFacade(ctrl),
				photoURLGenerator: mock_user.NewMockPhotoURLGenerator(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			usecase := NewUseCase(mocks.photoURLGenerator, mocks.userStorageFacade)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.GetUser(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
