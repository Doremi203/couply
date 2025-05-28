package search_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	mock_search_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/search"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_GetFilter(t *testing.T) {
	t.Parallel()

	type mocks struct {
		searchStorageFacade *mock_search_service.MocksearchStorageFacade
		photoURLGenerator   *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    *dto.GetFilterV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.GetFilterV1Response
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
				m.searchStorageFacade.EXPECT().GetFilterTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&search.Filter{
						UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, search.ErrFilterNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetFilterV1Request{},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, search.ErrFilterNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.searchStorageFacade.EXPECT().GetFilterTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&search.Filter{
						UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					}, nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetFilterV1Request{},
			},
			tokenErr: false,
			want: &dto.GetFilterV1Response{
				Filter: &search.Filter{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
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
			got, err := usecase.GetFilter(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
