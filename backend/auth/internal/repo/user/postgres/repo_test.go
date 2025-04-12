package userpostgres

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_repo_Create(t *testing.T) {
	type args struct {
		u user.User
	}
	tests := []struct {
		name     string
		args     args
		fixtures []string
		setup    func(*testing.T, context.Context, *repo)
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "create new user then success",
			args: args{
				u: user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Password: []byte("password"),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "user with email already exists then already exists error",
			args: args{
				u: user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Password: []byte("password"),
				},
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.Create(ctx, user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Password: []byte("password"),
				})
				require.NoError(t, err)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, user.ErrAlreadyExists)
			},
		},
	}
	for _, tt := range tests {
		tester.Run(tt.name, tt.fixtures, time.Second*10, t, func(t *testing.T, ctx context.Context, db postgres.Client) {
			r := &repo{
				db: db,
			}

			if tt.setup != nil {
				tt.setup(t, ctx, r)
			}

			err := r.Create(ctx, tt.args.u)

			tt.wantErr(t, err)
		})
	}
}

func Test_repo_GetByEmail(t *testing.T) {
	type args struct {
		email user.Email
	}
	tests := []struct {
		name     string
		args     args
		fixtures []string
		setup    func(*testing.T, context.Context, *repo)
		want     user.User
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "no user with given email then user not found error",
			args: args{
				email: "user@example.com",
			},
			want: user.User{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				var nfErr user.NotFoundError
				return assert.ErrorAs(t, err, &nfErr)
			},
		},
		{
			name: "user with given email exists then return user",
			args: args{
				email: "user@example.com",
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.Create(ctx, user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Password: []byte("password"),
				})
				require.NoError(t, err)
			},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		tester.Run(tt.name, tt.fixtures, time.Second*10, t, func(t *testing.T, ctx context.Context, db postgres.Client) {
			r := &repo{
				db: db,
			}

			if tt.setup != nil {
				tt.setup(t, ctx, r)
			}

			got, err := r.GetByEmail(ctx, tt.args.email)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
