package mock_gateway

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
)

type MockGateway struct {
}

func NewMockGateway() *MockGateway {
	return &MockGateway{}
}

func (g *MockGateway) CreatePayment(_ context.Context, _ int64, _ string) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (g *MockGateway) GetPaymentStatus(_ context.Context, gatewayID string) (payment.PaymentStatus, error) {
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

	return time.UnixMilli(int64(millis)) //nolint:gosec
}

func isPaymentSuccessful(id uuid.UUID) bool {
	h := sha256.New()
	h.Write(id[:])
	hash := h.Sum(nil)

	// 95% chance for success
	return int(hash[0])%100 < 95
}
