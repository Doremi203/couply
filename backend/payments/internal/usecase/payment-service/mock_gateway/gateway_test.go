package mock_gateway

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockGateway_CreatePayment(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		gateway := NewMockGateway()

		id, err := gateway.CreatePayment(context.Background(), 199, "RUB")
		assert.NoError(t, err)
		assert.NotEmpty(t, id)

		_, err = uuid.Parse(id)
		assert.NoError(t, err)
	})
}

func TestMockGateway_GetPaymentStatus(t *testing.T) {
	gateway := NewMockGateway()

	t.Run("invalid gateway ID", func(t *testing.T) {
		_, err := gateway.GetPaymentStatus(context.Background(), "invalid-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "uuid.Parse")
	})

	t.Run("pending status - recently created", func(t *testing.T) {
		id, _ := uuid.NewV7()
		status, err := gateway.GetPaymentStatus(context.Background(), id.String())
		require.NoError(t, err)
		assert.Equal(t, payment.PaymentStatusPending, status)
	})

	t.Run("success status - after delay", func(t *testing.T) {
		id := uuid.Must(uuid.Parse("017f3e42-7b27-7000-8000-000000000000"))
		status, err := gateway.GetPaymentStatus(context.Background(), id.String())
		require.NoError(t, err)
		assert.Equal(t, payment.PaymentStatusSuccess, status)
	})
}

func Test_extractTimeFromUUIDv7(t *testing.T) {
	now := time.Now()
	id, _ := uuid.NewV7()

	t.Run("extract time from UUIDv7", func(t *testing.T) {
		extracted := extractTimeFromUUIDv7(id)
		assert.WithinDuration(t, now, extracted, time.Millisecond)
	})
}
