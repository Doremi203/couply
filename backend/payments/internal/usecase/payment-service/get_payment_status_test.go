package payment_service

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/payment-service"
	mock_payment_service "github.com/Doremi203/couply/backend/payments/internal/mocks/usecase/payment"
	updater2 "github.com/Doremi203/couply/backend/payments/internal/usecase/updater"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCase_GetPaymentStatus(t *testing.T) {
	t.Parallel()

	now := time.Now()

	type mocks struct {
		paymentStorageFacade *mock_payment_service.MockpaymentStorageFacade
		paymentGatewayFacade *mock_payment_service.MockpaymentGateway
	}
	type args struct {
		token token.Token
		in    *dto.GetPaymentStatusV1Request
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		want    *dto.GetPaymentStatusV1Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "tx error",
			setup: func(m mocks) {
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(nil, payment.ErrPaymentNotFound)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetPaymentStatusV1Request{
					PaymentID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, payment.ErrPaymentNotFound)
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(
						&payment.Payment{
							ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
							Status:    payment.PaymentStatusPending,
							UpdatedAt: now,
						}, nil,
					)
			},
			args: args{
				token: token.Token{
					UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
				in: &dto.GetPaymentStatusV1Request{
					PaymentID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				},
			},
			want: &dto.GetPaymentStatusV1Response{
				PaymentID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				PaymentStatus: payment.PaymentStatusPending,
				UpdatedAt:     now,
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

			got, err := usecase.GetPaymentStatus(context.Background(), tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
