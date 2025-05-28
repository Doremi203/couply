package subscription_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
	mock_subscription_service "github.com/Doremi203/couply/backend/payments/internal/mocks/usecase/subscription"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_CancelSubscription(t *testing.T) {
	t.Parallel()

	type mocks struct {
		subscriptionStorageFacade *mock_subscription_service.MocksubscriptionStorageFacade
	}
	type args struct {
		in *dto.CancelSubscriptionV1Request
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		want    *dto.CancelSubscriptionV1Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "tx error",
			setup: func(m mocks) {
				m.subscriptionStorageFacade.EXPECT().CancelSubscriptionTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(subscription.ErrSubscriptionNotFound)
			},
			args: args{
				in: &dto.CancelSubscriptionV1Request{
					SubscriptionID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, subscription.ErrSubscriptionNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.subscriptionStorageFacade.EXPECT().CancelSubscriptionTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil)
			},
			args: args{
				in: &dto.CancelSubscriptionV1Request{
					SubscriptionID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
			},
			want:    &dto.CancelSubscriptionV1Response{},
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

			got, err := usecase.CancelSubscription(context.Background(), tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
