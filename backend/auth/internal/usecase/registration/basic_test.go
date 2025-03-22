package registration

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	mock_password "github.com/Doremi203/couply/backend/auth/internal/mocks/password"
	mock_user "github.com/Doremi203/couply/backend/auth/internal/mocks/user"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUseCase_BasicRegister(t *testing.T) {
	type mocks struct {
		UserRepository *mock_user.MockRepo
		Hasher         *mock_password.MockHasher
		UIDGenerator   *mock_user.MockUIDGenerator
	}
	type args struct {
		email    user.Email
		password user.Password
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "hashing password error then error",
			setup: func(m mocks) {
				m.Hasher.EXPECT().Hash(user.Password("password")).Return(user.HashedPassword{}, assert.AnError)
			},
			args: args{
				email:    "email",
				password: "password",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mocks := mocks{
				UserRepository: mock_user.NewMockRepo(ctrl),
				Hasher:         mock_password.NewMockHasher(ctrl),
				UIDGenerator:   mock_user.NewMockUIDGenerator(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			r := UseCase{
				userRepository: mocks.UserRepository,
				hasher:         mocks.Hasher,
				uidGenerator:   mocks.UIDGenerator,
			}
			err := r.BasicRegister(context.Background(), tt.args.email, tt.args.password)
			tt.wantErr(t, err)
		})
	}
}
