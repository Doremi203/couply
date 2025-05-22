package updater

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

type paymentStorageFacade interface {
	CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) (*payment.Payment, error)
	UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error
	GetPaymentStatusTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error)
	GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error)
}

type subscriptionStorageFacade interface {
	UpdateSubscriptionStatusTx(ctx context.Context, subID uuid.UUID, status subscription.SubscriptionStatus) error
	UpdateSubscriptionDatesTx(ctx context.Context, subID uuid.UUID, startDate, endDate time.Time) error
	GetSubscriptionsByStatusTx(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error)
	GetSubscriptionTx(ctx context.Context, subID uuid.UUID) (*subscription.Subscription, error)
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

func (u *Updater) StartPaymentStatusUpdater(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			u.updatePendingPayments(ctx)
		}
	}
}

func (u *Updater) StartSubscriptionStatusUpdater(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			u.processSubscriptionUpdates(ctx)
		}
	}
}
