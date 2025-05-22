package updater

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

const (
	initialDelay = 2500 * time.Millisecond
	maxRetries   = 3
	factor       = 2
)

var (
	errInvalidPaymentStatus = errors.Error("invalid payment status")
)

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
	delay := initialDelay
	for i := 0; i < maxRetries; i++ {
		time.Sleep(delay)

		status, err := u.paymentGateway.GetPaymentStatus(ctx, gatewayID)
		if err != nil {
			u.logger.Error(err)
			delay *= time.Duration(factor)
			continue
		}

		if err = u.processPaymentStatusUpdate(ctx, paymentID, status); err != nil {
			u.logger.Error(err)
			return
		}

		if status == payment.PaymentStatusPending {
			delay *= time.Duration(factor)
			continue
		}

		return
	}
}

func (u *Updater) processPaymentStatusUpdate(ctx context.Context, paymentID uuid.UUID, status payment.PaymentStatus) error {
	currentPayment, err := u.paymentStorageFacade.GetPaymentStatusTx(ctx, paymentID)
	if err != nil {
		return err
	}

	if status != currentPayment.GetStatus() {
		if err = u.paymentStorageFacade.UpdatePaymentStatusTx(ctx, paymentID, status); err != nil {
			return err
		}

		u.updateRelatedSubscription(ctx, currentPayment.GetSubscriptionID(), status)
	}
	return nil
}

func (u *Updater) updateRelatedSubscription(ctx context.Context, subID uuid.UUID, paymentStatus payment.PaymentStatus) {
	sub, err := u.subscriptionStorageFacade.GetSubscriptionTx(ctx, subID)
	if err != nil {
		u.logger.Error(err)
		return
	}

	newStatus := determineSubscriptionStatus(paymentStatus)
	if newStatus == subscription.SubscriptionStatusUnspecified {
		u.logger.Error(errInvalidPaymentStatus)
		return
	}

	if paymentStatus == payment.PaymentStatusSuccess && sub.GetStatus() == subscription.SubscriptionStatusPendingPayment {
		if err = u.activateSubscription(ctx, sub); err != nil {
			u.logger.Error(err)
			return
		}
	}

	if err = u.subscriptionStorageFacade.UpdateSubscriptionStatusTx(ctx, sub.GetID(), newStatus); err != nil {
		u.logger.Error(err)
	}
}

func determineSubscriptionStatus(paymentStatus payment.PaymentStatus) subscription.SubscriptionStatus {
	switch paymentStatus {
	case payment.PaymentStatusSuccess:
		return subscription.SubscriptionStatusActive
	case payment.PaymentStatusFailed:
		return subscription.SubscriptionStatusPendingPayment
	default:
		return subscription.SubscriptionStatusUnspecified
	}
}

func (u *Updater) activateSubscription(ctx context.Context, sub *subscription.Subscription) error {
	now := time.Now()
	endDate := subscription.CalculateEndDate(now, sub.GetPlan())
	if err := u.subscriptionStorageFacade.UpdateSubscriptionDatesTx(ctx, sub.GetID(), now, endDate); err != nil {
		return err
	}

	return u.updateUserPremiumStatus(ctx, sub.GetUserID(), true)
}
