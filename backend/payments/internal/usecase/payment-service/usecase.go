package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/usecase/updater"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
)

type paymentStorageFacade interface {
	paymentStorageSetterFacade
	paymentStorageGetterFacade
}

type paymentStorageSetterFacade interface {
	CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) (*payment.Payment, error)
	UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error
}

type paymentStorageGetterFacade interface {
	GetPaymentStatusTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error)
	GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error)
}

type paymentGateway interface {
	CreatePayment(ctx context.Context, amount int64, currency string) (string, error)
	GetPaymentStatus(ctx context.Context, gatewayID string) (payment.PaymentStatus, error)
}

type UseCase struct {
	paymentStorageFacade paymentStorageFacade
	paymentGateway       paymentGateway
	updater              *updater.Updater
}

func NewUseCase(paymentStorageFacade paymentStorageFacade, paymentGateway paymentGateway, updater *updater.Updater) *UseCase {
	return &UseCase{
		paymentStorageFacade: paymentStorageFacade,
		paymentGateway:       paymentGateway,
		updater:              updater,
	}
}
