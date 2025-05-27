package search_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	mock_search_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/search"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type loggerStub struct{}

func (l *loggerStub) Infof(_ string, _ ...any) {}

func (l *loggerStub) Error(_ error) {}

func (l *loggerStub) Warn(_ error) {}

func TestUseCase_AddView(t *testing.T) {
	t.Parallel()

	type mocks struct {
		searchStorageFacade *mock_search_service.MocksearchStorageFacade
		photoURLGenerator   *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    *dto.AddViewV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.AddViewV1Response
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
				m.searchStorageFacade.EXPECT().CreateViewTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112")).
					Return(user.ErrUserDoesntExist)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.AddViewV1Request{
					ViewedID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, user.ErrUserDoesntExist)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.searchStorageFacade.EXPECT().CreateViewTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					uuid.MustParse("11111111-1111-1111-1111-111111111112")).
					Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.AddViewV1Request{
					ViewedID: uuid.MustParse("11111111-1111-1111-1111-111111111112"),
				},
			},
			tokenErr: false,
			want:     &dto.AddViewV1Response{},
			wantErr:  assert.NoError,
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
			got, err := usecase.AddView(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
