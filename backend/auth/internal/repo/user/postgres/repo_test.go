package userpostgres

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/common/libs/ptr"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		tester.Run(t, tt.name, tt.fixtures, time.Second*10, func(t *testing.T, ctx context.Context, db postgres.Client) {
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

func Test_repo_GetByAny(t *testing.T) {
	type args struct {
		params user.GetByAnyParams
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
				params: user.GetByAnyParams{
					Email: "user@example.com",
				},
			},
			want: user.User{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, user.ErrNotFound)
			},
		},
		{
			name: "user with given email exists then return user",
			args: args{
				params: user.GetByAnyParams{
					Email: "user@example.com",
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
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "user with given phone exists then return user",
			args: args{
				params: user.GetByAnyParams{
					Phone: "+79123456789",
				},
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.Create(ctx, user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Phone:    "+79123456789",
					Password: []byte("password"),
				})
				require.NoError(t, err)
			},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "email or phone exists then return user",
			args: args{
				params: user.GetByAnyParams{
					Email: "user@example.com",
					Phone: "+79123456789",
				},
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.Create(ctx, user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Phone:    "+79123456789",
					Password: []byte("password"),
				})
				require.NoError(t, err)
			},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "user with oauth exists then return user",
			args: args{
				params: user.GetByAnyParams{
					OAuthUserAccount: &oauth.UserAccount{
						Provider:       oauth.YandexProvider,
						ProviderUserID: "123",
					},
				},
			},
			fixtures: []string{"yandex_oauth.sql"},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "user with phone or email or oauth exists then return user",
			args: args{
				params: user.GetByAnyParams{
					Email: "user@example.com",
					Phone: "+79123456789",
					OAuthUserAccount: &oauth.UserAccount{
						Provider:       oauth.YandexProvider,
						ProviderUserID: "123",
					},
				},
			},
			fixtures: []string{"yandex_oauth.sql"},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "user with id exists then return user",
			args: args{
				params: user.GetByAnyParams{
					ID: ptr.New[user.ID](user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111"))),
				},
			},
			fixtures: []string{"yandex_oauth.sql"},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "user with any of params exists then return user",
			args: args{
				params: user.GetByAnyParams{
					ID:    ptr.New[user.ID](user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111"))),
					Email: "user@example.com",
					Phone: "+79123456789",
					OAuthUserAccount: &oauth.UserAccount{
						Provider:       oauth.YandexProvider,
						ProviderUserID: "123",
					},
				},
			},
			fixtures: []string{"yandex_oauth.sql"},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "any of params but user not oauth then return user",
			args: args{
				params: user.GetByAnyParams{
					ID:    ptr.New[user.ID](user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111"))),
					Email: "user@example.com",
					Phone: "+79123456789",
					OAuthUserAccount: &oauth.UserAccount{
						Provider:       oauth.YandexProvider,
						ProviderUserID: "123",
					},
				},
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.Create(ctx, user.User{
					ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
					Email:    "user@example.com",
					Phone:    "+79123456789",
					Password: []byte("password"),
				})
				require.NoError(t, err)
			},
			want: user.User{
				ID:       user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				Email:    "user@example.com",
				Phone:    "+79123456789",
				Password: []byte("password"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		tester.Run(t, tt.name, tt.fixtures, time.Second*10, func(t *testing.T, ctx context.Context, db postgres.Client) {
			r := &repo{
				db: db,
			}

			if tt.setup != nil {
				tt.setup(t, ctx, r)
			}

			got, err := r.GetByAny(ctx, tt.args.params)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
