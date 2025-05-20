package updater

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
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
		u.logger.Error(err)
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
			u.logger.Error(err)
			delay *= time.Duration(factor)
			continue
		}

		currentPayment, err := u.paymentStorageFacade.GetPaymentStatusTx(ctx, paymentID)
		if err != nil {
			u.logger.Error(err)
			return
		}

		if status != currentPayment.GetStatus() {
			if err := u.paymentStorageFacade.UpdatePaymentStatusTx(ctx, paymentID, status); err != nil {
				u.logger.Error(err)
				return
			}

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
		u.logger.Error(err)
		return
	}

	var newStatus subscription.SubscriptionStatus

	switch paymentStatus {
	case payment.PaymentStatusSuccess:
		newStatus = subscription.SubscriptionStatusActive
		if sub.GetStatus() == subscription.SubscriptionStatusPendingPayment {
			now := time.Now()
			endDate := u.calculateEndDate(now, sub.GetPlan())
			err := u.subscriptionStorageFacade.UpdateSubscriptionDatesTx(ctx, sub.GetID(), now, endDate)
			if err != nil {
				u.logger.Error(err)
				return
			}
		}
	case payment.PaymentStatusFailed:
		newStatus = subscription.SubscriptionStatusPendingPayment
	default:
		return
	}

	err = u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(ctx, sub.GetID(), newStatus)
	if err != nil {
		u.logger.Error(err)
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
