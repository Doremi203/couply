package updater

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (u *Updater) processSubscriptionUpdates(ctx context.Context) {
	activeSubs, err := u.getActiveSubscriptions(ctx)
	if err != nil {
		u.logger.Error(err)
		return
	}

	u.checkAndUpdateExpiredSubscriptions(ctx, activeSubs)
}

func (u *Updater) getActiveSubscriptions(ctx context.Context) ([]*subscription.Subscription, error) {
	return u.subscriptionStorageFacade.GetSubscriptionsByStatusTx(ctx, subscription.SubscriptionStatusActive)
}

func (u *Updater) checkAndUpdateExpiredSubscriptions(ctx context.Context, subs []*subscription.Subscription) {
	now := time.Now()
	for _, sub := range subs {
		if u.isSubscriptionExpired(sub, now) {
			u.processExpiredSubscription(ctx, sub)
		}
	}
}

func (u *Updater) isSubscriptionExpired(sub *subscription.Subscription, now time.Time) bool {
	return sub.EndDate.Before(now)
}

func (u *Updater) processExpiredSubscription(ctx context.Context, sub *subscription.Subscription) {
	if sub.AutoRenew {
		u.handleAutoRenewal(ctx, sub)
	} else {
		u.expireSubscription(ctx, sub)
	}
}

func (u *Updater) expireSubscription(ctx context.Context, sub *subscription.Subscription) {
	err := u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(
		ctx,
		sub.ID,
		subscription.SubscriptionStatusExpired,
	)
	if err != nil {
		u.logger.Error(err)
		return
	}

	if err = u.updateUserPremiumStatus(ctx, sub.UserID, false); err != nil {
		u.logger.Error(err)
	}
}

func (u *Updater) handleAutoRenewal(ctx context.Context, sub *subscription.Subscription) {
	paymentID, err := uuid.NewV7()
	if err != nil {
		u.logger.Error(err)
		return
	}

	amount := subscription.GetPlanPrice(sub.Plan)
	gatewayID, err := u.paymentGateway.CreatePayment(ctx, amount, payment.MainCurrency)
	if err != nil {
		u.logger.Error(err)
		return
	}

	newPayment := u.createPaymentObject(paymentID, sub, amount, gatewayID)
	if err = u.savePaymentAndUpdateSubscription(ctx, newPayment, sub); err != nil {
		u.logger.Error(err)
	}
}

func (u *Updater) createPaymentObject(
	paymentID uuid.UUID,
	sub *subscription.Subscription,
	amount int64,
	gatewayID string,
) *payment.Payment {
	now := time.Now()
	return &payment.Payment{
		ID:             paymentID,
		UserID:         sub.UserID,
		SubscriptionID: sub.ID,
		Amount:         amount,
		Currency:       payment.MainCurrency,
		Status:         payment.PaymentStatusPending,
		GatewayID:      gatewayID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func (u *Updater) savePaymentAndUpdateSubscription(
	ctx context.Context,
	payment *payment.Payment,
	sub *subscription.Subscription,
) error {
	if _, err := u.paymentStorageFacade.CreatePaymentTx(ctx, payment); err != nil {
		return err
	}

	return u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(
		ctx,
		sub.ID,
		subscription.SubscriptionStatusPendingPayment,
	)
}
