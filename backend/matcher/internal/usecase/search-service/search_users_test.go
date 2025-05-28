package search_service

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	mock_search_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/search"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_SearchUsers(t *testing.T) {
	t.Parallel()

	now := time.Now()

	type mocks struct {
		searchStorageFacade *mock_search_service.MocksearchStorageFacade
		photoURLGenerator   *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    *dto.SearchUsersV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.SearchUsersV1Response
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
				m.searchStorageFacade.EXPECT().SearchUsersTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uint64(0),
					uint64(1)).
					Return(nil, nil, search.ErrFilterNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.SearchUsersV1Request{
					Offset: 0,
					Limit:  1,
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, search.ErrFilterNotFound)
			},
		},
		{
			name: "generate download urls error",
			setup: func(m mocks) {
				m.searchStorageFacade.EXPECT().SearchUsersTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uint64(0),
					uint64(1)).
					Return([]*user.User{
						{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
							Photos: []user.Photo{
								{
									UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111112"),
									OrderNumber: 0,
									ObjectKey:   "users/11111111-1111-1111-1111-111111111112/slot/0.jpg",
									MimeType:    ".jpg",
									UploadedAt:  &now,
									UploadURL:   func(s string) *string { return &s }("uploadURL"),
									DownloadURL: nil,
								},
							},
						},
					}, map[uuid.UUID]float64{
						uuid.MustParse("11111111-1111-1111-1111-111111111112"): 1,
					}, nil)

				m.photoURLGenerator.EXPECT().GenerateDownload(gomock.Any(), "users/11111111-1111-1111-1111-111111111112/slot/0.jpg").
					Return("", errors.Error("some error"))
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.SearchUsersV1Request{
					Offset: 0,
					Limit:  1,
				},
			},
			tokenErr: false,
			want: &dto.SearchUsersV1Response{
				UsersSearchInfo: []*search.UserSearchInfo{
					{
						User: &user.User{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
							Photos: []user.Photo{
								{
									UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111112"),
									OrderNumber: 0,
									ObjectKey:   "users/11111111-1111-1111-1111-111111111112/slot/0.jpg",
									MimeType:    ".jpg",
									UploadedAt:  &now,
									UploadURL:   func(s string) *string { return &s }("uploadURL"),
									DownloadURL: nil,
								},
							},
						},
						DistanceToUser: 1,
					},
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "no distance to user",
			setup: func(m mocks) {
				m.searchStorageFacade.EXPECT().SearchUsersTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uint64(0),
					uint64(1)).
					Return([]*user.User{
						{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
							Photos: []user.Photo{
								{
									UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111112"),
									OrderNumber: 0,
									ObjectKey:   "users/11111111-1111-1111-1111-111111111112/slot/0.jpg",
									MimeType:    ".jpg",
									UploadedAt:  &now,
									UploadURL:   func(s string) *string { return &s }("uploadURL"),
									DownloadURL: nil,
								},
							},
						},
					}, nil, nil)

				m.photoURLGenerator.EXPECT().GenerateDownload(gomock.Any(), "users/11111111-1111-1111-1111-111111111112/slot/0.jpg").
					Return("downloadURL", nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.SearchUsersV1Request{
					Offset: 0,
					Limit:  1,
				},
			},
			tokenErr: false,
			want: &dto.SearchUsersV1Response{
				UsersSearchInfo: []*search.UserSearchInfo{
					{
						User: &user.User{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
							Photos: []user.Photo{
								{
									UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111112"),
									OrderNumber: 0,
									ObjectKey:   "users/11111111-1111-1111-1111-111111111112/slot/0.jpg",
									MimeType:    ".jpg",
									UploadedAt:  &now,
									UploadURL:   func(s string) *string { return &s }("uploadURL"),
									DownloadURL: func(s string) *string { return &s }("downloadURL"),
								},
							},
						},
						DistanceToUser: 0,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.searchStorageFacade.EXPECT().SearchUsersTx(gomock.Any(),
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uint64(0),
					uint64(1)).
					Return([]*user.User{
						{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
							Photos: []user.Photo{
								{
									UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111112"),
									OrderNumber: 0,
									ObjectKey:   "users/11111111-1111-1111-1111-111111111112/slot/0.jpg",
									MimeType:    ".jpg",
									UploadedAt:  &now,
									UploadURL:   func(s string) *string { return &s }("uploadURL"),
									DownloadURL: nil,
								},
							},
						},
					}, map[uuid.UUID]float64{
						uuid.MustParse("11111111-1111-1111-1111-111111111112"): 1,
					}, nil)

				m.photoURLGenerator.EXPECT().GenerateDownload(gomock.Any(), "users/11111111-1111-1111-1111-111111111112/slot/0.jpg").
					Return("downloadURL", nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.SearchUsersV1Request{
					Offset: 0,
					Limit:  1,
				},
			},
			tokenErr: false,
			want: &dto.SearchUsersV1Response{
				UsersSearchInfo: []*search.UserSearchInfo{
					{
						User: &user.User{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
							Photos: []user.Photo{
								{
									UserID:      uuid.MustParse("11111111-1111-1111-1111-111111111112"),
									OrderNumber: 0,
									ObjectKey:   "users/11111111-1111-1111-1111-111111111112/slot/0.jpg",
									MimeType:    ".jpg",
									UploadedAt:  &now,
									UploadURL:   func(s string) *string { return &s }("uploadURL"),
									DownloadURL: func(s string) *string { return &s }("downloadURL"),
								},
							},
						},
						DistanceToUser: 1,
					},
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
				searchStorageFacade: mock_search_service.NewMocksearchStorageFacade(ctrl),
				photoURLGenerator:   mock_user.NewMockPhotoURLGenerator(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			logger := &loggerStub{}

			usecase := NewUseCase(mocks.searchStorageFacade, mocks.photoURLGenerator, logger)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.SearchUsers(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
