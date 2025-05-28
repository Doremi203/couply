package search_service

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	mock_search_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/search"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_UpdateFilter(t *testing.T) {
	t.Parallel()

	now := time.Now()

	type mocks struct {
		searchStorageFacade *mock_search_service.MocksearchStorageFacade
		photoURLGenerator   *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    *dto.UpdateFilterV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.UpdateFilterV1Response
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
				m.searchStorageFacade.EXPECT().UpdateFilterTx(gomock.Any(), &search.Filter{
					UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Goal:      common.GoalDating,
					UpdatedAt: now,
				}).Return(user.ErrUserDoesntExist)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.UpdateFilterV1Request{
					Goal:      common.GoalDating,
					UpdatedAt: now,
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
				m.searchStorageFacade.EXPECT().UpdateFilterTx(gomock.Any(), &search.Filter{
					UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Goal:      common.GoalDating,
					UpdatedAt: now,
				}).
					Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.UpdateFilterV1Request{
					Goal:      common.GoalDating,
					UpdatedAt: now,
				},
			},
			tokenErr: false,
			want: &dto.UpdateFilterV1Response{
				Filter: &search.Filter{
					UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Goal:      common.GoalDating,
					UpdatedAt: now,
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
			got, err := usecase.UpdateFilter(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
