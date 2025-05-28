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

func TestUseCase_CreateSubscription(t *testing.T) {
	t.Parallel()

	type mocks struct {
		subscriptionStorageFacade *mock_subscription_service.MocksubscriptionStorageFacade
	}
	type args struct {
		token token.Token
		in    *dto.CreateSubscriptionV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.CreateSubscriptionV1Response
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
				m.subscriptionStorageFacade.EXPECT().CreateSubscriptionTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, s *subscription.Subscription) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), s.UserID)
					assert.Equal(t, subscription.SubscriptionPlanMonthly, s.Plan)
					assert.Equal(t, subscription.SubscriptionStatusPendingPayment, s.Status)
					assert.Equal(t, true, s.AutoRenew)

					assert.NotEmpty(t, s.ID)
					assert.NotZero(t, s.StartDate)
					assert.NotZero(t, s.EndDate)

					return nil
				}).Return(subscription.ErrSubscriptionHasAlreadyBeenPaid)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.CreateSubscriptionV1Request{
					SubscriptionPlan: subscription.SubscriptionPlanMonthly,
					AutoRenew:        true,
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, subscription.ErrSubscriptionHasAlreadyBeenPaid)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.subscriptionStorageFacade.EXPECT().CreateSubscriptionTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, s *subscription.Subscription) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), s.UserID)
					assert.Equal(t, subscription.SubscriptionPlanMonthly, s.Plan)
					assert.Equal(t, subscription.SubscriptionStatusPendingPayment, s.Status)
					assert.Equal(t, true, s.AutoRenew)

					assert.NotEmpty(t, s.ID)
					assert.NotZero(t, s.StartDate)
					assert.NotZero(t, s.EndDate)

					return nil
				}).Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.CreateSubscriptionV1Request{
					SubscriptionPlan: subscription.SubscriptionPlanMonthly,
					AutoRenew:        true,
				},
			},
			tokenErr: false,
			want: &dto.CreateSubscriptionV1Response{
				SubscriptionPlan:   subscription.SubscriptionPlanMonthly,
				SubscriptionStatus: subscription.SubscriptionStatusPendingPayment,
				AutoRenew:          true,
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
			got, err := usecase.CreateSubscription(ctx, tt.args.in)

			tt.wantErr(t, err)
			if tt.want != nil {
				assert.Equal(t, tt.want.SubscriptionPlan, got.SubscriptionPlan)
				assert.Equal(t, tt.want.SubscriptionStatus, got.SubscriptionStatus)
				assert.Equal(t, tt.want.AutoRenew, got.AutoRenew)
			}
		})
	}
}
