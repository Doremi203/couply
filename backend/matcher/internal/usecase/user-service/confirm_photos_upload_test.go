package user_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	mock_user_service "github.com/Doremi203/couply/backend/matcher/internal/mocks/usecase/user"
	mock_user "github.com/Doremi203/couply/backend/matcher/internal/mocks/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_ConfirmPhotosUpload(t *testing.T) {
	t.Parallel()

	type mocks struct {
		userStorageFacade *mock_user_service.MockuserStorageFacade
		photoURLGenerator *mock_user.MockPhotoURLGenerator
	}
	type args struct {
		token token.Token
		in    []int32
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
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
				m.userStorageFacade.EXPECT().UpdatePhotosUploadedAtTx(gomock.Any(), []int32{0}, uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(user.ErrPhotoNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: []int32{0},
			},
			tokenErr: false,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, user.ErrPhotoNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.userStorageFacade.EXPECT().UpdatePhotosUploadedAtTx(gomock.Any(), []int32{0}, uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: []int32{0},
			},
			tokenErr: false,
			wantErr:  assert.NoError,
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

			uuidProvider := &providerStub{}

			usecase := NewUseCase(mocks.photoURLGenerator, mocks.userStorageFacade, uuidProvider)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			err := usecase.ConfirmPhotosUpload(ctx, tt.args.in)

			tt.wantErr(t, err)
		})
	}
}
