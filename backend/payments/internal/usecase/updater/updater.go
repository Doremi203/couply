package updater

import (
	"context"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/log"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

type paymentStorageFacade interface {
	CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) (*payment.Payment, error)
	GetPaymentStatusTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error)
	GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error)
	UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error
}

type subscriptionStorageFacade interface {
	GetSubscriptionsByStatusTx(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error)
	UpdateSubscriptionStatusTx(ctx context.Context, subscriptionID uuid.UUID, status subscription.SubscriptionStatus) error
	GetSubscriptionTx(ctx context.Context, subID uuid.UUID) (*subscription.Subscription, error)
	UpdateSubscriptionDatesTx(ctx context.Context, subscriptionID uuid.UUID, startDate, endDate time.Time) error
}

type paymentGateway interface {
	CreatePayment(ctx context.Context, amount int64, currency string) (string, error)
	GetPaymentStatus(ctx context.Context, gatewayID string) (payment.PaymentStatus, error)
}

type userClient interface {
	GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error)
	UpdateUserByIDV1(ctx context.Context, user *userservicegrpc.User) error
}

type Updater struct {
	paymentStorageFacade      paymentStorageFacade
	subscriptionStorageFacade subscriptionStorageFacade
	paymentGateway            paymentGateway
	userClient                userClient
	logger                    log.Logger
}

func NewUpdater(ps paymentStorageFacade, subs subscriptionStorageFacade, gateway paymentGateway, userClient userClient, logger log.Logger) *Updater {
	return &Updater{
		paymentStorageFacade:      ps,
		subscriptionStorageFacade: subs,
		paymentGateway:            gateway,
		userClient:                userClient,
		logger:                    logger,
	}
}
