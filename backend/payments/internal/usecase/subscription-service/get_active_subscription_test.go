package subscription_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
	mock_subscription_service "github.com/Doremi203/couply/backend/payments/internal/mocks/usecase/subscription"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_GetActiveSubscription(t *testing.T) {
	t.Parallel()

	type mocks struct {
		subscriptionStorageFacade *mock_subscription_service.MocksubscriptionStorageFacade
	}
	type args struct {
		token token.Token
		in    *dto.GetActiveSubscriptionV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.GetActiveSubscriptionV1Response
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
				m.subscriptionStorageFacade.EXPECT().GetActiveSubscriptionTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, subscription.ErrActiveSubscriptionDoesntExist)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetActiveSubscriptionV1Request{},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, subscription.ErrActiveSubscriptionDoesntExist)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.subscriptionStorageFacade.EXPECT().GetActiveSubscriptionTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(
						&subscription.Subscription{
							ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						}, nil,
					)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetActiveSubscriptionV1Request{},
			},
			tokenErr: false,
			want: &dto.GetActiveSubscriptionV1Response{
				SubscriptionID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mocks := mocks{
				subscriptionStorageFacade: mock_subscription_service.NewMocksubscriptionStorageFacade(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			usecase := NewUseCase(mocks.subscriptionStorageFacade)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.GetActiveSubscription(ctx, tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
