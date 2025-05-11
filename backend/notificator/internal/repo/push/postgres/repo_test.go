package pushpostgres

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_repo_UpsertSubscription(t *testing.T) {
	type args struct {
		subscription push.Subscription
	}
	tests := []struct {
		name     string
		args     args
		fixtures []string
		setup    func(*testing.T, context.Context, *repo)
		want     push.Subscription
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "no subscriptions exist then insert",
			args: args{
				subscription: push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
			},
			want: push.Subscription{
				Endpoint: "endpoint",
				Credentials: push.Credentials{
					P256dh:  "public_key",
					AuthKey: "auth_key",
				},
				UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
			},
			wantErr: assert.NoError,
		},
		{
			name: "subscriptions exist then update",
			args: args{
				subscription: push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key_updated",
						AuthKey: "auth_key_updated",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-2222-111111111111")),
				},
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.UpsertSubscription(ctx, push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				})
				require.NoError(t, err)
			},
			want: push.Subscription{
				Endpoint: "endpoint",
				Credentials: push.Credentials{
					P256dh:  "public_key_updated",
					AuthKey: "auth_key_updated",
				},
				UserID: user.ID(uuid.MustParse("11111111-1111-1111-2222-111111111111")),
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

			err := r.UpsertSubscription(ctx, tt.args.subscription)
			tt.wantErr(t, err)

			subscriptions, err := r.GetSubscriptionsByUserID(ctx, tt.args.subscription.UserID)
			require.NoError(t, err)
			require.Len(t, subscriptions, 1)

			assert.Equal(t, tt.want, subscriptions[0])
		})
	}
}

func Test_repo_GetSubscriptionsByUserID(t *testing.T) {
	type args struct {
		userID user.ID
	}
	tests := []struct {
		name     string
		args     args
		fixtures []string
		setup    func(*testing.T, context.Context, *repo)
		want     []push.Subscription
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "no subscriptions then nil",
			args: args{
				userID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "all subscriptions for one user",
			args: args{
				userID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.UpsertSubscription(ctx, push.Subscription{
					Endpoint: "endpoint_1",
					Credentials: push.Credentials{
						P256dh:  "public_key_1",
						AuthKey: "auth_key_1",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				})
				require.NoError(t, err)
				err = r.UpsertSubscription(ctx, push.Subscription{
					Endpoint: "endpoint_2",
					Credentials: push.Credentials{
						P256dh:  "public_key_2",
						AuthKey: "auth_key_2",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				})
				require.NoError(t, err)
			},
			want: []push.Subscription{
				{
					Endpoint: "endpoint_1",
					Credentials: push.Credentials{
						P256dh:  "public_key_1",
						AuthKey: "auth_key_1",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
				{
					Endpoint: "endpoint_2",
					Credentials: push.Credentials{
						P256dh:  "public_key_2",
						AuthKey: "auth_key_2",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "different users with subscriptions",
			args: args{
				userID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
			},
			setup: func(t *testing.T, ctx context.Context, r *repo) {
				err := r.UpsertSubscription(ctx, push.Subscription{
					Endpoint: "endpoint_1",
					Credentials: push.Credentials{
						P256dh:  "public_key_1",
						AuthKey: "auth_key_1",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				})
				require.NoError(t, err)
				err = r.UpsertSubscription(ctx, push.Subscription{
					Endpoint: "endpoint_2",
					Credentials: push.Credentials{
						P256dh:  "public_key_2",
						AuthKey: "auth_key_2",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-2222-111111111111")),
				})
				require.NoError(t, err)
			},
			want: []push.Subscription{
				{
					Endpoint: "endpoint_1",
					Credentials: push.Credentials{
						P256dh:  "public_key_1",
						AuthKey: "auth_key_1",
					},
					UserID: user.ID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
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

			got, err := r.GetSubscriptionsByUserID(ctx, tt.args.userID)

			tt.wantErr(t, err)
			assert.ElementsMatch(t, tt.want, got, "GetSubscriptionsByUserID(ctx, %v)", tt.args.userID)
		})
	}
}
