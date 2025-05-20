package updater

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (u *Updater) StartSubscriptionStatusUpdater(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			u.updateExpiredSubscriptions(ctx)
		}
	}
}

func (u *Updater) updateExpiredSubscriptions(ctx context.Context) {
	activeSubs, err := u.subscriptionStorageFacade.GetSubscriptionsByStatusTx(ctx, subscription.SubscriptionStatusActive)
	if err != nil {
		u.logger.Error(err)
		return
	}

	now := time.Now()
	for _, sub := range activeSubs {
		if sub.GetEndDate().Before(now) {
			if sub.GetAutoRenew() {
				u.handleAutoRenewal(ctx, sub)
			} else {
				err := u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(ctx, sub.GetID(), subscription.SubscriptionStatusExpired)
				if err != nil {
					u.logger.Error(err)
					continue
				}
			}
		}
	}
}

func (u *Updater) handleAutoRenewal(ctx context.Context, sub *subscription.Subscription) {
	paymentID, err := uuid.NewV7()
	if err != nil {
		u.logger.Error(err)
		return
	}

	amount := u.getPlanAmount(sub.GetPlan())

	gatewayID, err := u.paymentGateway.CreatePayment(ctx, amount, "RUB")
	if err != nil {
		u.logger.Error(err)
		return
	}

	now := time.Now()
	newPayment := &payment.Payment{
		ID:             paymentID,
		UserID:         sub.GetUserID(),
		SubscriptionID: sub.GetID(),
		Amount:         amount,
		Currency:       "RUB",
		Status:         payment.PaymentStatusPending,
		GatewayID:      gatewayID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	_, err = u.paymentStorageFacade.CreatePaymentTx(ctx, newPayment)
	if err != nil {
		u.logger.Error(err)
		return
	}

	err = u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(ctx, sub.GetID(), subscription.SubscriptionStatusPendingPayment)
	if err != nil {
		u.logger.Error(err)
		return
	}
}

func (u *Updater) getPlanAmount(plan subscription.SubscriptionPlan) int64 {
	planPrices := map[subscription.SubscriptionPlan]int64{
		subscription.SubscriptionPlanMonthly:    199,
		subscription.SubscriptionPlanSemiAnnual: 799,
		subscription.SubscriptionPlanAnnual:     1999,
	}
	return planPrices[plan]
}
