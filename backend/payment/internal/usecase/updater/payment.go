package updater

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	"github.com/google/uuid"
)

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

func (u *Updater) updatePendingPayments(ctx context.Context) {
	pendingPayments, err := u.paymentStorageFacade.GetPendingPaymentsTx(ctx)
	if err != nil {
		return
	}

	for _, p := range pendingPayments {
		go u.CheckAndUpdatePaymentStatusWithRetry(ctx, p.GetID(), p.GetGatewayID())
	}
}

func (u *Updater) CheckAndUpdatePaymentStatusWithRetry(ctx context.Context, paymentID uuid.UUID, gatewayID string) {
	const (
		initialDelay = 2500 * time.Millisecond
		maxRetries   = 3
		factor       = 2
	)

	delay := initialDelay
	for i := 0; i < maxRetries; i++ {
		time.Sleep(delay)

		status, err := u.paymentGateway.GetPaymentStatus(ctx, gatewayID)
		if err != nil {
			// TODO: add logger
			delay *= time.Duration(factor)
			continue
		}

		currentPayment, err := u.paymentStorageFacade.GetPaymentStatusTx(ctx, paymentID)
		if err != nil {
			// TODO: add logger
			return
		}

		if status != currentPayment.GetStatus() {
			if err := u.paymentStorageFacade.UpdatePaymentStatusTx(ctx, paymentID, status); err != nil {
				// TODO: add logger
				return
			}

			// Update related subscriptions
			u.updateRelatedSubscriptions(ctx, currentPayment.GetSubscriptionID(), status)
		}

		if status == payment.PaymentStatusPending {
			delay *= time.Duration(factor)
			continue
		}

		return
	}
}

func (u *Updater) updateRelatedSubscriptions(ctx context.Context, subID uuid.UUID, paymentStatus payment.PaymentStatus) {
	sub, err := u.subscriptionStorageFacade.GetSubscriptionTx(ctx, subID)
	if err != nil {
		// TODO: add logger
		return
	}

	var newStatus subscription.SubscriptionStatus

	switch paymentStatus {
	case payment.PaymentStatusSuccess:
		newStatus = subscription.SubscriptionStatusActive
		// For new subscriptions, we might want to update the dates
		if sub.GetStatus() == subscription.SubscriptionStatusPendingPayment {
			now := time.Now()
			endDate := u.calculateEndDate(now, sub.GetPlan())
			// Need to extend the storage facade to handle date updates
			err := u.subscriptionStorageFacade.UpdateSubscriptionDatesTx(ctx, sub.GetID(), now, endDate)
			if err != nil {
				// TODO: add logger
				return
			}
		}
	case payment.PaymentStatusFailed:
		newStatus = subscription.SubscriptionStatusPendingPayment
	case payment.PaymentStatusRefunded:
		newStatus = subscription.SubscriptionStatusCanceled
	default:
		return
	}

	err = u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(ctx, sub.GetID(), newStatus)
	if err != nil {
		// TODO: add logger
		return
	}
}

func (u *Updater) calculateEndDate(now time.Time, plan subscription.SubscriptionPlan) time.Time {
	durationMap := map[subscription.SubscriptionPlan]time.Duration{
		subscription.SubscriptionPlanMonthly:    30 * 24 * time.Hour,
		subscription.SubscriptionPlanSemiAnnual: 180 * 24 * time.Hour,
		subscription.SubscriptionPlanAnnual:     365 * 24 * time.Hour,
	}

	return now.Add(durationMap[plan])
}
