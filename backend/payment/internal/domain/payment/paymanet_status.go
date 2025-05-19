package payment

import desc "github.com/Doremi203/couply/backend/payment/gen/api/payment-service/v1"

type PaymentStatus int

const (
	PaymentStatusUnspecified PaymentStatus = iota
	PaymentStatusPending
	PaymentStatusSuccess
	PaymentStatusFailed
	PaymentStatusRefunded
)

func PBToSPaymentStatus(paymentStatus desc.PaymentStatus) PaymentStatus {
	switch paymentStatus {
	case desc.PaymentStatus_PAYMENT_STATUS_UNSPECIFIED:
		return PaymentStatusUnspecified
	case desc.PaymentStatus_PAYMENT_STATUS_PENDING:
		return PaymentStatusPending
	case desc.PaymentStatus_PAYMENT_STATUS_SUCCESS:
		return PaymentStatusSuccess
	case desc.PaymentStatus_PAYMENT_STATUS_FAILED:
		return PaymentStatusFailed
	case desc.PaymentStatus_PAYMENT_STATUS_REFUNDED:
		return PaymentStatusRefunded
	default:
		return PaymentStatus(0)
	}
}

func PaymentStatusToPB(paymentStatus PaymentStatus) desc.PaymentStatus {
	switch paymentStatus {
	case PaymentStatusUnspecified:
		return desc.PaymentStatus_PAYMENT_STATUS_UNSPECIFIED
	case PaymentStatusPending:
		return desc.PaymentStatus_PAYMENT_STATUS_PENDING
	case PaymentStatusSuccess:
		return desc.PaymentStatus_PAYMENT_STATUS_SUCCESS
	case PaymentStatusFailed:
		return desc.PaymentStatus_PAYMENT_STATUS_FAILED
	case PaymentStatusRefunded:
		return desc.PaymentStatus_PAYMENT_STATUS_REFUNDED
	default:
		return desc.PaymentStatus(0)
	}
}
