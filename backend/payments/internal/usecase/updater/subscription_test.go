package updater

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	mock_updater "github.com/Doremi203/couply/backend/payments/internal/mocks/usecase/updater"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdater_processExpiredSubscription(t *testing.T) {
	t.Parallel()

	type mocks struct {
		subscriptionFacade *mock_updater.MocksubscriptionStorageFacade
		paymentGateway     *mock_updater.MockpaymentGateway
		paymentStorage     *mock_updater.MockpaymentStorageFacade
		userClient         *mock_updater.MockuserClient
	}

	tests := []struct {
		name         string
		autoRenew    bool
		setupMocks   func(*mocks, *subscription.Subscription)
		expectedLogs []error
	}{
		{
			name:      "auto renew success",
			autoRenew: true,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.paymentGateway.EXPECT().CreatePayment(gomock.Any(), int64(199), payment.MainCurrency).
					Return("new_gateway_id", nil)
				m.paymentStorage.EXPECT().CreatePaymentTx(gomock.Any(), gomock.Any()).Return(nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(
					gomock.Any(),
					sub.ID,
					subscription.SubscriptionStatusPendingPayment,
				).Return(nil)
			},
			expectedLogs: nil,
		},
		{
			name:      "expire without renew",
			autoRenew: false,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(
					gomock.Any(),
					sub.ID,
					subscription.SubscriptionStatusExpired,
				).Return(nil)
				m.userClient.EXPECT().GetUserByIDV1(gomock.Any(), sub.UserID.String()).
					Return(&userservicegrpc.User{}, nil)
				m.userClient.EXPECT().UpdateUserByIDV1(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedLogs: nil,
		},
		{
			name:      "auto renew fails on payment gateway error",
			autoRenew: true,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.paymentGateway.EXPECT().CreatePayment(gomock.Any(), int64(199), payment.MainCurrency).
					Return("", errors.Error("payment gateway error"))
			},
			expectedLogs: []error{errors.Error("payment gateway error")},
		},
		{
			name:      "auto renew fails on payment creation error",
			autoRenew: true,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.paymentGateway.EXPECT().CreatePayment(gomock.Any(), int64(199), payment.MainCurrency).
					Return("gateway_id", nil)
				m.paymentStorage.EXPECT().CreatePaymentTx(gomock.Any(), gomock.Any()).
					Return(errors.Error("db error"))
			},
			expectedLogs: []error{errors.Error("db error")},
		},
		{
			name:      "auto renew fails on subscription update error",
			autoRenew: true,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.paymentGateway.EXPECT().CreatePayment(gomock.Any(), int64(199), payment.MainCurrency).
					Return("gateway_id", nil)
				m.paymentStorage.EXPECT().CreatePaymentTx(gomock.Any(), gomock.Any()).Return(nil)
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(
					gomock.Any(),
					sub.ID,
					subscription.SubscriptionStatusPendingPayment,
				).Return(errors.Error("update error"))
			},
			expectedLogs: []error{errors.Error("update error")},
		},
		{
			name:      "expire subscription fails on update error",
			autoRenew: false,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(
					gomock.Any(),
					sub.ID,
					subscription.SubscriptionStatusExpired,
				).Return(errors.Error("update error"))
			},
			expectedLogs: []error{errors.Error("update error")},
		},
		{
			name:      "expire subscription user update fails",
			autoRenew: false,
			setupMocks: func(m *mocks, sub *subscription.Subscription) {
				m.subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(
					gomock.Any(),
					sub.ID,
					subscription.SubscriptionStatusExpired,
				).Return(nil)
				m.userClient.EXPECT().GetUserByIDV1(gomock.Any(), sub.UserID.String()).
					Return(&userservicegrpc.User{}, nil)
				m.userClient.EXPECT().UpdateUserByIDV1(gomock.Any(), gomock.Any()).
					Return(errors.Error("user update error"))
			},
			expectedLogs: []error{errors.Error("user update error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			m := &mocks{
				subscriptionFacade: mock_updater.NewMocksubscriptionStorageFacade(ctrl),
				paymentGateway:     mock_updater.NewMockpaymentGateway(ctrl),
				paymentStorage:     mock_updater.NewMockpaymentStorageFacade(ctrl),
				userClient:         mock_updater.NewMockuserClient(ctrl),
			}

			sub := &subscription.Subscription{
				ID:        uuid.New(),
				UserID:    uuid.New(),
				Plan:      subscription.SubscriptionPlanMonthly,
				AutoRenew: tt.autoRenew,
				EndDate:   time.Now().Add(-time.Hour),
			}

			logger := &loggerStub{}

			u := &Updater{
				subscriptionStorageFacade: m.subscriptionFacade,
				paymentGateway:            m.paymentGateway,
				paymentStorageFacade:      m.paymentStorage,
				userClient:                m.userClient,
				logger:                    logger,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(m, sub)
			}

			u.processExpiredSubscription(context.Background(), sub)

			assert.Equal(t, len(tt.expectedLogs), len(logger.errors), "количество ошибок в логе не совпадает")
			for i, expectedErr := range tt.expectedLogs {
				if i < len(logger.errors) {
					assert.EqualError(t, logger.errors[i], expectedErr.Error(), "ошибка в логе не совпадает")
				}
			}
		})
	}
}

func TestUpdater_isSubscriptionExpired(t *testing.T) {
	u := &Updater{}

	now := time.Now()
	tests := []struct {
		name     string
		endDate  time.Time
		expected bool
	}{
		{
			name:     "subscription expired",
			endDate:  now.Add(-time.Hour),
			expected: true,
		},
		{
			name:     "subscription active",
			endDate:  now.Add(time.Hour),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sub := &subscription.Subscription{EndDate: tt.endDate}
			assert.Equal(t, tt.expected, u.isSubscriptionExpired(sub, now))
		})
	}
}

func TestUpdater_processSubscriptionUpdates(t *testing.T) {
	t.Parallel()

	type mocks struct {
		subscriptionFacade *mock_updater.MocksubscriptionStorageFacade
	}

	tests := []struct {
		name         string
		setupMocks   func(*mocks)
		expectedLogs []error
	}{
		{
			name: "error getting active subscriptions",
			setupMocks: func(m *mocks) {
				m.subscriptionFacade.EXPECT().
					GetSubscriptionsByStatusTx(gomock.Any(), subscription.SubscriptionStatusActive).
					Return(nil, errors.Error("database error"))
			},
			expectedLogs: []error{errors.Error("database error")},
		},
		{
			name: "no active subscriptions",
			setupMocks: func(m *mocks) {
				m.subscriptionFacade.EXPECT().
					GetSubscriptionsByStatusTx(gomock.Any(), subscription.SubscriptionStatusActive).
					Return([]*subscription.Subscription{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			m := &mocks{
				subscriptionFacade: mock_updater.NewMocksubscriptionStorageFacade(ctrl),
			}

			logger := &loggerStub{}
			u := &Updater{
				subscriptionStorageFacade: m.subscriptionFacade,
				logger:                    logger,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(m)
			}

			u.processSubscriptionUpdates(context.Background())

			assert.Equal(t, len(tt.expectedLogs), len(logger.errors), "unexpected error count")
			for i, expectedErr := range tt.expectedLogs {
				assert.EqualError(t, logger.errors[i], expectedErr.Error(), "error mismatch")
			}
		})
	}
}

func TestUpdater_getActiveSubscriptions(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	subscriptionFacade := mock_updater.NewMocksubscriptionStorageFacade(ctrl)

	expectedSubs := []*subscription.Subscription{
		{ID: uuid.New(), Status: subscription.SubscriptionStatusActive},
		{ID: uuid.New(), Status: subscription.SubscriptionStatusActive},
	}

	subscriptionFacade.EXPECT().
		GetSubscriptionsByStatusTx(gomock.Any(), subscription.SubscriptionStatusActive).
		Return(expectedSubs, nil)

	u := &Updater{subscriptionStorageFacade: subscriptionFacade}

	subs, err := u.getActiveSubscriptions(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedSubs, subs)
}

func TestUpdater_checkAndUpdateExpiredSubscriptions_EmptyList(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	subscriptionFacade := mock_updater.NewMocksubscriptionStorageFacade(ctrl)
	logger := &loggerStub{}

	u := &Updater{
		subscriptionStorageFacade: subscriptionFacade,
		logger:                    logger,
	}

	subscriptionFacade.EXPECT().UpdateSubscriptionStatusTx(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	u.checkAndUpdateExpiredSubscriptions(context.Background(), []*subscription.Subscription{})

	assert.Empty(t, logger.errors)
}
