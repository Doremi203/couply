package hash

import (
	"testing"

	mock_argon "github.com/Doremi203/couply/backend/auth/pkg/argon/mock"
	mock_salt "github.com/Doremi203/couply/backend/auth/pkg/salt/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDefaultProvider_Hash(t *testing.T) {
	type mocks struct {
		saltProvider  *mock_salt.MockProvider
		argonProvider *mock_argon.MockProvider
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "salt generate error then error",
			setup: func(m mocks) {
				m.saltProvider.EXPECT().Generate(saltLength).Return(nil, assert.AnError)
			},
			args:    args{},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.saltProvider.EXPECT().Generate(saltLength).Return([]byte("salt"), nil)
				m.argonProvider.EXPECT().Hash([]byte("value"), []byte("salt")).Return([]byte("argon-hash"))
			},
			args: args{
				value: "value",
			},
			want:    []byte("YXJnb24taGFzaA$c2FsdA"),
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

			h := defaultProvider{
				saltProvider:  mocks.saltProvider,
				argonProvider: mocks.argonProvider,
			}
			got, err := h.Hash(tt.args.value)

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
		value      string
		saltedHash []byte
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "incorrectly hashed value then error",
			args: args{
				value:      "pass",
				saltedHash: []byte("part1$part2$part3"),
			},
			wantErr: assert.Error,
		},
		{
			name: "value does not match hashed value then invalid value error",
			setup: func(m mocks) {
				m.argonProvider.EXPECT().Hash([]byte("incorrect"), []byte("salt")).Return([]byte("incorrect-hash"))
			},
			args: args{
				value:      "incorrect",
				saltedHash: []byte("YXJnb24taGFzaA$c2FsdA"),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrNoMatch)
			},
		},
		{
			name: "value match hashed value then no error",
			setup: func(m mocks) {
				m.argonProvider.EXPECT().Hash([]byte("value"), []byte("salt")).Return([]byte("argon-hash"))
			},
			args: args{
				value:      "value",
				saltedHash: []byte("YXJnb24taGFzaA$c2FsdA"),
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

			h := defaultProvider{
				saltProvider:  mocks.saltProvider,
				argonProvider: mocks.argonProvider,
			}

			err := h.Verify(tt.args.value, tt.args.saltedHash)
			tt.wantErr(t, err)
		})
	}
}
