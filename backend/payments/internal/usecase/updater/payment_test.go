package updater

import (
	"context"
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	mock_updater "github.com/Doremi203/couply/backend/payments/internal/mocks/usecase/updater"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type loggerStub struct {
	errors []error
}

func (l *loggerStub) Infof(_ string, _ ...any) {}

func (l *loggerStub) Error(err error) {
	l.errors = append(l.errors, err)
}

func (l *loggerStub) Warn(_ error) {}

func TestUpdater_CheckAndUpdatePaymentStatusWithRetry(t *testing.T) {
	t.Parallel()

	type mocks struct {
		paymentGateway       *mock_updater.MockpaymentGateway
		paymentStorageFacade *mock_updater.MockpaymentStorageFacade
		subscriptionFacade   *mock_updater.MocksubscriptionStorageFacade
		userClient           *mock_updater.MockuserClient
	}

	type args struct {
		paymentID uuid.UUID
		gatewayID string
	}

	tests := []struct {
		name       string
		setupMocks func(*mocks, args)
		args       args
		wantErr    bool
	}{
		{
			name: "success after retries",
			setupMocks: func(m *mocks, a args) {
				// first two tries - err, third - success
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Times(2).Return(payment.PaymentStatusPending, status.Error(codes.Unavailable, "gateway error"))
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Return(payment.PaymentStatusSuccess, nil)
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), a.paymentID).
					Return(&payment.Payment{Status: payment.PaymentStatusPending}, nil)
				m.paymentStorageFacade.EXPECT().UpdatePaymentStatusTx(gomock.Any(), a.paymentID, payment.PaymentStatusSuccess).
					Return(nil)
				m.subscriptionFacade.EXPECT().GetSubscriptionTx(gomock.Any(), gomock.Any()).Return(&subscription.Subscription{
					Status: subscription.SubscriptionStatusPendingPayment,
				}, nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(gomock.Any(), gomock.Any(), subscription.SubscriptionStatusActive).
					Return(nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionDatesTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
				m.userClient.EXPECT().GetUserByIDV1(gomock.Any(), gomock.Any()).Return(&userservicegrpc.User{}, nil)
				m.userClient.EXPECT().UpdateUserByIDV1(gomock.Any(), gomock.Any()).Return(nil)
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: false,
		},
		{
			name: "max retries exceeded",
			setupMocks: func(m *mocks, a args) {
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Times(3).Return(payment.PaymentStatusPending, status.Error(codes.Unavailable, "gateway error"))
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: true,
		},
		{
			name: "error getting payment by ID",
			setupMocks: func(m *mocks, a args) {
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Return(payment.PaymentStatusSuccess, nil)
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), a.paymentID).
					Return(nil, errors.Error("storage error"))
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: true,
		},
		{
			name: "error updating subscription dates",
			setupMocks: func(m *mocks, a args) {
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Return(payment.PaymentStatusSuccess, nil)
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), a.paymentID).
					Return(&payment.Payment{
						Status:         payment.PaymentStatusPending,
						SubscriptionID: uuid.New(),
					}, nil)
				m.paymentStorageFacade.EXPECT().UpdatePaymentStatusTx(gomock.Any(), a.paymentID, payment.PaymentStatusSuccess).
					Return(nil)
				m.subscriptionFacade.EXPECT().GetSubscriptionTx(gomock.Any(), gomock.Any()).
					Return(&subscription.Subscription{
						Status: subscription.SubscriptionStatusPendingPayment,
					}, nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionDatesTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(errors.Error("dates update error"))
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: true,
		},
		{
			name: "error updating user premium status",
			setupMocks: func(m *mocks, a args) {
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Return(payment.PaymentStatusSuccess, nil)
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), a.paymentID).
					Return(&payment.Payment{
						Status:         payment.PaymentStatusPending,
						SubscriptionID: uuid.New(),
					}, nil)
				m.paymentStorageFacade.EXPECT().UpdatePaymentStatusTx(gomock.Any(), a.paymentID, payment.PaymentStatusSuccess).
					Return(nil)
				m.subscriptionFacade.EXPECT().GetSubscriptionTx(gomock.Any(), gomock.Any()).
					Return(&subscription.Subscription{
						Status: subscription.SubscriptionStatusPendingPayment,
					}, nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionDatesTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
				m.userClient.EXPECT().GetUserByIDV1(gomock.Any(), gomock.Any()).Return(&userservicegrpc.User{}, nil)
				m.userClient.EXPECT().UpdateUserByIDV1(gomock.Any(), gomock.Any()).Return(errors.Error("user update error"))
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: true,
		},
		{
			name: "subscription not found during update",
			setupMocks: func(m *mocks, a args) {
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Return(payment.PaymentStatusSuccess, nil)
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), a.paymentID).
					Return(&payment.Payment{
						Status:         payment.PaymentStatusPending,
						SubscriptionID: uuid.New(),
					}, nil)
				m.paymentStorageFacade.EXPECT().UpdatePaymentStatusTx(gomock.Any(), a.paymentID, payment.PaymentStatusSuccess).
					Return(nil)
				m.subscriptionFacade.EXPECT().GetSubscriptionTx(gomock.Any(), gomock.Any()).
					Return(nil, errors.Error("subscription not found"))
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: true,
		},
		{
			name: "subscription already active",
			setupMocks: func(m *mocks, a args) {
				m.paymentGateway.EXPECT().GetPaymentStatus(gomock.Any(), a.gatewayID).
					Return(payment.PaymentStatusSuccess, nil)
				m.paymentStorageFacade.EXPECT().GetPaymentByIDTx(gomock.Any(), a.paymentID).
					Return(&payment.Payment{
						Status:         payment.PaymentStatusPending,
						SubscriptionID: uuid.New(),
					}, nil)
				m.paymentStorageFacade.EXPECT().UpdatePaymentStatusTx(gomock.Any(), a.paymentID, payment.PaymentStatusSuccess).
					Return(nil)
				m.subscriptionFacade.EXPECT().GetSubscriptionTx(gomock.Any(), gomock.Any()).
					Return(&subscription.Subscription{
						Status: subscription.SubscriptionStatusActive,
					}, nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(gomock.Any(), gomock.Any(), subscription.SubscriptionStatusActive).
					Return(nil)
			},
			args: args{
				paymentID: uuid.New(),
				gatewayID: "test_gateway",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			m := &mocks{
				paymentGateway:       mock_updater.NewMockpaymentGateway(ctrl),
				paymentStorageFacade: mock_updater.NewMockpaymentStorageFacade(ctrl),
				subscriptionFacade:   mock_updater.NewMocksubscriptionStorageFacade(ctrl),
				userClient:           mock_updater.NewMockuserClient(ctrl),
			}

			logger := &loggerStub{}

			u := &Updater{
				paymentGateway:            m.paymentGateway,
				paymentStorageFacade:      m.paymentStorageFacade,
				subscriptionStorageFacade: m.subscriptionFacade,
				userClient:                m.userClient,
				logger:                    logger,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(m, tt.args)
			}

			u.CheckAndUpdatePaymentStatusWithRetry(context.Background(), tt.args.paymentID, tt.args.gatewayID)
		})
	}
}
