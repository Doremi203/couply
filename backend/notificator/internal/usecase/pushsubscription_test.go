package usecase

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	mock_push "github.com/Doremi203/couply/backend/notificator/internal/mocks/push"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPushSubscription_Subscribe(t *testing.T) {
	type mocks struct {
		pushRepo *mock_push.MockRepo
	}
	type args struct {
		subscription push.Subscription
	}
	tests := []struct {
		name    string
		args    args
		setup   func(mocks)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "upsert success then no error",
			args: args{
				subscription: push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
			},
			setup: func(m mocks) {
				m.pushRepo.EXPECT().UpsertSubscription(gomock.Any(), push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				}).Return(nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "upsert error then error",
			args: args{
				subscription: push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
			},
			setup: func(m mocks) {
				m.pushRepo.EXPECT().UpsertSubscription(gomock.Any(), push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				}).Return(assert.AnError)
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mocks := mocks{
				pushRepo: mock_push.NewMockRepo(ctrl),
			}
			if tt.setup != nil {
				tt.setup(mocks)
			}

			s := PushSubscription{
				pushRepo: mocks.pushRepo,
			}

			err := s.Subscribe(context.Background(), tt.args.subscription)

			tt.wantErr(t, err)
		})
	}
}

func TestPushSubscription_Unsubscribe(t *testing.T) {
	type mocks struct {
		pushRepo *mock_push.MockRepo
	}
	type args struct {
		subscription push.Subscription
	}
	tests := []struct {
		name    string
		args    args
		setup   func(mocks)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "delete success then no error",
			args: args{
				subscription: push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
			},
			setup: func(m mocks) {
				m.pushRepo.EXPECT().DeleteSubscription(gomock.Any(), push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				}).Return(nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "delete error then error",
			args: args{
				subscription: push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				},
			},
			setup: func(m mocks) {
				m.pushRepo.EXPECT().DeleteSubscription(gomock.Any(), push.Subscription{
					Endpoint: "endpoint",
					Credentials: push.Credentials{
						P256dh:  "public_key",
						AuthKey: "auth_key",
					},
					RecipientID: push.RecipientID(uuid.MustParse("11111111-1111-1111-1111-111111111111")),
				}).Return(assert.AnError)
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mocks := mocks{
				pushRepo: mock_push.NewMockRepo(ctrl),
			}
			if tt.setup != nil {
				tt.setup(mocks)
			}

			s := PushSubscription{
				pushRepo: mocks.pushRepo,
			}

			err := s.Unsubscribe(context.Background(), tt.args.subscription)

			tt.wantErr(t, err)
		})
	}
}
