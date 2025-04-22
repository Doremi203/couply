package pswrd

import (
	"testing"

	mock_argon "github.com/Doremi203/couply/backend/auth/pkg/argon/mock"
	mock_salt "github.com/Doremi203/couply/backend/auth/pkg/salt/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDefaultHasher_Hash(t *testing.T) {
	type mocks struct {
		saltProvider  *mock_salt.MockProvider
		argonProvider *mock_argon.MockProvider
	}
	type args struct {
		password Password
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		want    HashedPassword
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "salt generate error then error",
			setup: func(m mocks) {
				m.saltProvider.EXPECT().Generate(saltLength).Return(nil, assert.AnError)
			},
			args:    args{},
			want:    HashedPassword{},
			wantErr: assert.Error,
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.saltProvider.EXPECT().Generate(saltLength).Return([]byte("salt"), nil)
				m.argonProvider.EXPECT().Hash([]byte("password"), []byte("salt")).Return([]byte("argon-hash"))
			},
			args: args{
				password: "password",
			},
			want:    HashedPassword("YXJnb24taGFzaA$c2FsdA"),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mocks := mocks{
				saltProvider:  mock_salt.NewMockProvider(ctrl),
				argonProvider: mock_argon.NewMockProvider(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			h := DefaultHasher{
				saltProvider:  mocks.saltProvider,
				argonProvider: mocks.argonProvider,
			}
			got, err := h.Hash(tt.args.password)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDefaultHasher_Verify(t *testing.T) {
	type mocks struct {
		saltProvider  *mock_salt.MockProvider
		argonProvider *mock_argon.MockProvider
	}
	type args struct {
		password       Password
		hashedPassword HashedPassword
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "incorrectly hashed password then error",
			args: args{
				password:       "pass",
				hashedPassword: HashedPassword("part1$part2$part3"),
			},
			wantErr: assert.Error,
		},
		{
			name: "password does not match hashed password then invalid password error",
			setup: func(m mocks) {
				m.argonProvider.EXPECT().Hash([]byte("incorrect"), []byte("salt")).Return([]byte("incorrect-hash"))
			},
			args: args{
				password:       "incorrect",
				hashedPassword: HashedPassword("YXJnb24taGFzaA$c2FsdA"),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrInvalidPassword)
			},
		},
		{
			name: "password match hashed password then no error",
			setup: func(m mocks) {
				m.argonProvider.EXPECT().Hash([]byte("password"), []byte("salt")).Return([]byte("argon-hash"))
			},
			args: args{
				password:       "password",
				hashedPassword: HashedPassword("YXJnb24taGFzaA$c2FsdA"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mocks := mocks{
				saltProvider:  mock_salt.NewMockProvider(ctrl),
				argonProvider: mock_argon.NewMockProvider(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			h := DefaultHasher{
				saltProvider:  mocks.saltProvider,
				argonProvider: mocks.argonProvider,
			}

			err := h.Verify(tt.args.password, tt.args.hashedPassword)
			tt.wantErr(t, err)
		})
	}
}
