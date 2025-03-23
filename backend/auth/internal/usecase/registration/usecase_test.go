package registration

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	mock_password "github.com/Doremi203/couply/backend/auth/internal/mocks/password"
	mock_user "github.com/Doremi203/couply/backend/auth/internal/mocks/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/idempotency"
	mock_uuid "github.com/Doremi203/couply/backend/auth/pkg/uuid/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUseCase_BasicRegister(t *testing.T) {
	type mocks struct {
		userRepository *mock_user.MockRepo
		hasher         *mock_password.MockHasher
		uuidProvider   *mock_uuid.MockProvider
	}
	type args struct {
		idempotencyKey idempotency.Key
		email          user.Email
		password       pswrd.Password
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "get existing user by email error then error",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, assert.AnError)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: assert.Error,
		},
		{
			name: "get existing user by email success then already exists error",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, nil)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, usecase.ErrAlreadyRegistered)
			},
		},
		{
			name: "hashing password error then error",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, user.ErrNotFound)
				m.hasher.EXPECT().Hash(pswrd.Password("password")).Return(pswrd.HashedPassword{}, assert.AnError)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: assert.Error,
		},
		{
			name: "create user error then error",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, user.ErrNotFound)
				m.hasher.EXPECT().Hash(pswrd.Password("password")).Return(pswrd.HashedPassword{}, nil)
				m.uuidProvider.EXPECT().GenerateV7().Return(uuid.UUID{}, assert.AnError)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: assert.Error,
		},
		{
			name: "create user error then error",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, user.ErrNotFound)
				hashedPassword := pswrd.HashedPassword("password-hash")
				m.hasher.EXPECT().Hash(pswrd.Password("password")).Return(hashedPassword, nil)
				id := uuid.New()
				m.uuidProvider.EXPECT().GenerateV7().Return(id, nil)

				usr := user.User{
					ID:       user.ID(id),
					Email:    "email",
					Password: hashedPassword,
				}
				m.userRepository.EXPECT().Create(gomock.Any(), usr).Return(assert.AnError)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: assert.Error,
		},
		{
			name: "create user already exists error then success",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, user.ErrNotFound)
				hashedPassword := pswrd.HashedPassword("password-hash")
				m.hasher.EXPECT().Hash(pswrd.Password("password")).Return(hashedPassword, nil)
				id := uuid.New()
				m.uuidProvider.EXPECT().GenerateV7().Return(id, nil)

				usr := user.User{
					ID:       user.ID(id),
					Email:    "email",
					Password: hashedPassword,
				}
				m.userRepository.EXPECT().Create(gomock.Any(), usr).Return(user.ErrAlreadyExists)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: assert.NoError,
		},
		{
			name: "create user success then success",
			setup: func(m mocks) {
				m.userRepository.EXPECT().GetByEmail(gomock.Any(), user.Email("email")).Return(user.User{}, user.ErrNotFound)
				hashedPassword := pswrd.HashedPassword("password-hash")
				m.hasher.EXPECT().Hash(pswrd.Password("password")).Return(hashedPassword, nil)
				id := uuid.New()
				m.uuidProvider.EXPECT().GenerateV7().Return(id, nil)

				usr := user.User{
					ID:       user.ID(id),
					Email:    "email",
					Password: hashedPassword,
				}
				m.userRepository.EXPECT().Create(gomock.Any(), usr).Return(nil)
			},
			args: args{
				idempotencyKey: "key",
				email:          "email",
				password:       "password",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mocks := mocks{
				userRepository: mock_user.NewMockRepo(ctrl),
				hasher:         mock_password.NewMockHasher(ctrl),
				uuidProvider:   mock_uuid.NewMockProvider(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			r := UseCase{
				userRepository: mocks.userRepository,
				hasher:         mocks.hasher,
				uuidProvider:   mocks.uuidProvider,
			}
			err := r.BasicRegister(context.Background(), tt.args.email, tt.args.password)
			tt.wantErr(t, err)
		})
	}
}
