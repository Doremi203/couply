package mock_gateway

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	"github.com/google/uuid"
	"time"
)

type MockGateway struct {
}

func NewMockGateway() *MockGateway {
	return &MockGateway{}
}

func (g *MockGateway) CreatePayment(ctx context.Context, amount int64, currency string) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (g *MockGateway) GetPaymentStatus(ctx context.Context, gatewayID string) (payment.PaymentStatus, error) {
	parsedGatewayID, err := uuid.Parse(gatewayID)
	if err != nil {
		return 0, err
	}

	createdAt := extractTimeFromUUIDv7(parsedGatewayID)
	now := time.Now()

	if now.Sub(createdAt) < 2*time.Second {
		return payment.PaymentStatusPending, nil
	}

	if isPaymentSuccessful(parsedGatewayID) {
		return payment.PaymentStatusSuccess, nil
	}
	return payment.PaymentStatusFailed, nil
}

func extractTimeFromUUIDv7(id uuid.UUID) time.Time {
	// UUID v7 contains time mark in first 48 bits
	bytes := id[:6]

	var buf [8]byte
	copy(buf[:], bytes)

	millis := binary.BigEndian.Uint64(buf[:]) >> 16

	return time.UnixMilli(int64(millis))
}

func isPaymentSuccessful(id uuid.UUID) bool {
	h := sha256.New()
	h.Write(id[:])
	hash := h.Sum(nil)

	// 95% chance for success
	return int(hash[0])%100 < 95
}
