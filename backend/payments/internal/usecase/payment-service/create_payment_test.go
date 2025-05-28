package payment_service

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/payment-service"
	mock_payment_service "github.com/Doremi203/couply/backend/payments/internal/mocks/usecase/payment"
	updater2 "github.com/Doremi203/couply/backend/payments/internal/usecase/updater"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_CreatePayment(t *testing.T) {
	t.Parallel()

	type mocks struct {
		paymentStorageFacade *mock_payment_service.MockpaymentStorageFacade
		paymentGatewayFacade *mock_payment_service.MockpaymentGateway
	}
	type args struct {
		token token.Token
		in    *dto.CreatePaymentV1Request
	}
	tests := []struct {
		name     string
		setup    func(mocks)
		args     args
		tokenErr bool
		want     *dto.CreatePaymentV1Response
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
				m.paymentGatewayFacade.EXPECT().CreatePayment(gomock.Any(), int64(199), "RUB").
					Return("11111111-1111-1111-1111-111111111111", nil)

				m.paymentStorageFacade.EXPECT().CreatePaymentTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, p *payment.Payment) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), p.UserID)
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), p.SubscriptionID)
					assert.Equal(t, int64(199), p.Amount)
					assert.Equal(t, "RUB", p.Currency)
					assert.Equal(t, "11111111-1111-1111-1111-111111111111", p.GatewayID)

					assert.NotEmpty(t, p.ID)
					assert.NotZero(t, p.CreatedAt)
					assert.NotZero(t, p.UpdatedAt)

					return nil
				}).Return(payment.ErrSubscriptionDoesntExist)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.CreatePaymentV1Request{
					SubscriptionID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Amount:         199,
					Currency:       "RUB",
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, payment.ErrSubscriptionDoesntExist)
			},
		},
		{
			name: "create payment error",
			setup: func(m mocks) {
				m.paymentGatewayFacade.EXPECT().CreatePayment(gomock.Any(), int64(199), "RUB").
					Return("", errors.Error("some error"))

			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.CreatePaymentV1Request{
					SubscriptionID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Amount:         199,
					Currency:       "RUB",
				},
			},
			tokenErr: false,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.paymentGatewayFacade.EXPECT().CreatePayment(gomock.Any(), int64(199), "RUB").
					Return("11111111-1111-1111-1111-111111111111", nil)

				m.paymentStorageFacade.EXPECT().CreatePaymentTx(
					gomock.Any(),
					gomock.Any(),
				).Do(func(_ context.Context, p *payment.Payment) error {
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), p.UserID)
					assert.Equal(t, uuid.MustParse("11111111-1111-1111-1111-111111111111"), p.SubscriptionID)
					assert.Equal(t, int64(199), p.Amount)
					assert.Equal(t, "RUB", p.Currency)
					assert.Equal(t, "11111111-1111-1111-1111-111111111111", p.GatewayID)

					assert.NotEmpty(t, p.ID)
					assert.NotZero(t, p.CreatedAt)
					assert.NotZero(t, p.UpdatedAt)

					return nil
				}).Return(nil)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.CreatePaymentV1Request{
					SubscriptionID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Amount:         199,
					Currency:       "RUB",
				},
			},
			tokenErr: false,
			want: &dto.CreatePaymentV1Response{
				PaymentID: "11111111-1111-1111-1111-111111111111",
				Status:    payment.PaymentStatusPending,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mocks := mocks{
				paymentStorageFacade: mock_payment_service.NewMockpaymentStorageFacade(ctrl),
				paymentGatewayFacade: mock_payment_service.NewMockpaymentGateway(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			updater := &updater2.Updater{}

			usecase := NewUseCase(mocks.paymentStorageFacade, mocks.paymentGatewayFacade, updater)

			ctx := token.ContextWithToken(context.Background(), tt.args.token)
			if tt.tokenErr {
				ctx = context.Background()
			}
			got, err := usecase.CreatePayment(ctx, tt.args.in)

			tt.wantErr(t, err)
			if !tt.tokenErr && tt.want != nil {
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}
