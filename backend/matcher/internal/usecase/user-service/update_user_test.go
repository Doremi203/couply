package user_service

import (
	"context"
	"testing"
	"time"

	mock_user_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/user"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_UpdateUser(t *testing.T) {
	now := time.Now()

	type mocks struct {
		userStorageFacade *mock_user_service.MockuserStorageFacade
		photoURLGenerator *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    *dto.UpdateUserV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.UpdateUserV1Response
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
			name: "create photos error",
			setup: func(m mocks) {
				m.photoURLGenerator.EXPECT().GenerateUpload(gomock.Any(), "users/11111111-1111-1111-1111-111111111111/slot/0.jpg", ".jpg").Return("", errors.Error("error"))
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.UpdateUserV1Request{
					Name:       "user",
					Age:        18,
					Gender:     user.GenderMale,
					Latitude:   32,
					Longitude:  32,
					Bio:        "",
					Goal:       0,
					Interest:   nil,
					Zodiac:     0,
					Height:     198,
					Education:  0,
					Children:   0,
					Alcohol:    0,
					Smoking:    0,
					IsHidden:   false,
					IsVerified: false,
					IsPremium:  false,
					IsBlocked:  false,
					PhotoUploadRequests: []user.PhotoUploadRequest{
						{
							OrderNumber: 0,
							MimeType:    ".jpg",
						},
					},
					UpdatedAt: now,
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "tx error",
			setup: func(m mocks) {
				m.userStorageFacade.EXPECT().UpdateUserTx(gomock.Any(), &user.User{
					ID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:       "user",
					Age:        18,
					Gender:     user.GenderMale,
					Latitude:   32,
					Longitude:  32,
					BIO:        "",
					Goal:       0,
					Interest:   nil,
					Zodiac:     0,
					Height:     198,
					Education:  0,
					Children:   0,
					Alcohol:    0,
					Smoking:    0,
					IsHidden:   false,
					IsVerified: false,
					IsPremium:  false,
					IsBlocked:  false,
					Photos:     []user.Photo{},
					CreatedAt:  time.Time{},
					UpdatedAt:  now,
				}).Return(user.ErrUserNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.UpdateUserV1Request{
					Name:                "user",
					Age:                 18,
					Gender:              user.GenderMale,
					Latitude:            32,
					Longitude:           32,
					Bio:                 "",
					Goal:                0,
					Interest:            nil,
					Zodiac:              0,
					Height:              198,
					Education:           0,
					Children:            0,
					Alcohol:             0,
					Smoking:             0,
					IsHidden:            false,
					IsVerified:          false,
					IsPremium:           false,
					IsBlocked:           false,
					PhotoUploadRequests: nil,
					UpdatedAt:           now,
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, user.ErrUserNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.photoURLGenerator.EXPECT().GenerateUpload(gomock.Any(), "users/11111111-1111-1111-1111-111111111111/slot/0.jpg", ".jpg").Return("uploadURL", nil)
				m.userStorageFacade.EXPECT().UpdateUserTx(gomock.Any(), &user.User{
					ID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:       "user",
					Age:        18,
					Gender:     user.GenderMale,
					Latitude:   32,
					Longitude:  32,
					BIO:        "",
					Goal:       0,
					Interest:   nil,
					Zodiac:     0,
					Height:     198,
					Education:  0,
					Children:   0,
					Alcohol:    0,
					Smoking:    0,
					IsHidden:   false,
					IsVerified: false,
					IsPremium:  false,
					IsBlocked:  false,
					Photos: []user.Photo{
						{
							UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
							OrderNumber: 0,
							ObjectKey:   "users/11111111-1111-1111-1111-111111111111/slot/0.jpg",
							MimeType:    ".jpg",
							UploadedAt:  nil,
							UploadURL:   func(s string) *string { return &s }("uploadURL"),
							DownloadURL: nil,
						},
					},
					CreatedAt: time.Time{},
					UpdatedAt: now,
				}).Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.UpdateUserV1Request{
					Name:       "user",
					Age:        18,
					Gender:     user.GenderMale,
					Latitude:   32,
					Longitude:  32,
					Bio:        "",
					Goal:       0,
					Interest:   nil,
					Zodiac:     0,
					Height:     198,
					Education:  0,
					Children:   0,
					Alcohol:    0,
					Smoking:    0,
					IsHidden:   false,
					IsVerified: false,
					IsPremium:  false,
					IsBlocked:  false,
					PhotoUploadRequests: []user.PhotoUploadRequest{
						{
							OrderNumber: 0,
							MimeType:    ".jpg",
						},
					},
					UpdatedAt: now,
				},
			},
			want: &dto.UpdateUserV1Response{
				User: &user.User{
					ID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:       "user",
					Age:        18,
					Gender:     user.GenderMale,
					Latitude:   32,
					Longitude:  32,
					BIO:        "",
					Goal:       0,
					Interest:   nil,
					Zodiac:     0,
					Height:     198,
					Education:  0,
					Children:   0,
					Alcohol:    0,
					Smoking:    0,
					IsHidden:   false,
					IsVerified: false,
					IsPremium:  false,
					IsBlocked:  false,
					Photos: []user.Photo{
						{
							UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
							OrderNumber: 0,
							ObjectKey:   "users/11111111-1111-1111-1111-111111111111/slot/0.jpg",
							MimeType:    ".jpg",
							UploadedAt:  nil,
							UploadURL:   func(s string) *string { return &s }("uploadURL"),
							DownloadURL: nil,
						},
					},
					CreatedAt: time.Time{},
					UpdatedAt: now,
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			got, err := usecase.UpdateUser(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
