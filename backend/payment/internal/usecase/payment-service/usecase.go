package payment_service

import (
	"context"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	"github.com/google/uuid"
	"time"
)

type paymentStorageFacade interface {
	CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) (*payment.Payment, error)
	GetPaymentStatusTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error)
	GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error)
	UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error
}

type paymentGateway interface {
	CreatePayment(ctx context.Context, amount int64, currency string) (string, error)
	GetPaymentStatus(ctx context.Context, gatewayID string) (payment.PaymentStatus, error)
}

type UseCase struct {
	paymentStorageFacade paymentStorageFacade
	paymentGateway       paymentGateway
}

func NewUseCase(paymentStorageFacade paymentStorageFacade, paymentGateway paymentGateway) *UseCase {
	uc := &UseCase{
		paymentStorageFacade: paymentStorageFacade,
		paymentGateway:       paymentGateway,
	}

	go uc.startPaymentStatusUpdater(context.Background(), 30*time.Second)

	return uc
}

func (c *UseCase) startPaymentStatusUpdater(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.updatePendingPayments(ctx)
		}
	}
}

func (c *UseCase) updatePendingPayments(ctx context.Context) {
	pendingPayments, err := c.paymentStorageFacade.GetPendingPaymentsTx(ctx)
	if err != nil {
		return
	}

	for _, p := range pendingPayments {
		go c.checkAndUpdatePaymentStatusWithRetry(ctx, p.GetID(), p.GetGatewayID())
	}
}

func (c *UseCase) checkAndUpdatePaymentStatusWithRetry(ctx context.Context, paymentID uuid.UUID, gatewayID string) {
	const (
		initialDelay = 2500 * time.Millisecond
		maxRetries   = 3
		factor       = 2
	)

	delay := initialDelay
	for i := 0; i < maxRetries; i++ {
		time.Sleep(delay)

		status, err := c.paymentGateway.GetPaymentStatus(ctx, gatewayID)
		if err != nil {
			// TODO: add logger
			delay *= time.Duration(factor)
			continue
		}

		currentPayment, err := c.paymentStorageFacade.GetPaymentStatusTx(ctx, paymentID)
		if err != nil {
			// TODO: add logger
			return
		}

		if status != currentPayment.GetStatus() {
			if err := c.paymentStorageFacade.UpdatePaymentStatusTx(ctx, paymentID, status); err != nil {
				// TODO: add logger
			}
			return
		}

		if status == payment.PaymentStatusPending {
			delay *= time.Duration(factor)
			continue
		}

		return
	}
}
