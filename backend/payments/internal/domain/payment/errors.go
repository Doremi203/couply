package payment

import "github.com/Doremi203/couply/backend/auth/pkg/errors"

var (
	ErrDuplicatePayment = errors.Error("payment already exists")
	ErrPaymentNotFound  = errors.Error("payment not found")
	ErrPaymentsNotFound = errors.Error("payments not found")
)
