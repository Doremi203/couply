package subscription_service

import (
	"time"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/subscription-service/v1"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetActiveSubscriptionV1Request struct{}

func PBToGetActiveSubscriptionRequest(_ *desc.GetActiveSubscriptionV1Request) *GetActiveSubscriptionV1Request {
	return &GetActiveSubscriptionV1Request{}
}

type GetActiveSubscriptionV1Response struct {
	SubscriptionID     uuid.UUID
	SubscriptionPlan   subscription.SubscriptionPlan
	SubscriptionStatus subscription.SubscriptionStatus
	AutoRenew          bool
	StartDate          time.Time
	EndDate            time.Time
	PaymentIDs         []uuid.UUID
}

func GetActiveSubscriptionResponseToPB(resp *GetActiveSubscriptionV1Response) *desc.GetActiveSubscriptionV1Response {
	paymentIds := make([]string, 0, len(resp.PaymentIDs))
	for _, id := range resp.PaymentIDs {
		paymentIds = append(paymentIds, id.String())
	}
	return &desc.GetActiveSubscriptionV1Response{
		SubscriptionId: resp.SubscriptionID.String(),
		Plan:           subscription.SubscriptionPlanToPB(resp.SubscriptionPlan),
		Status:         subscription.SubscriptionStatusToPB(resp.SubscriptionStatus),
		AutoRenew:      resp.AutoRenew,
		StartDate:      timestamppb.New(resp.StartDate),
		EndDate:        timestamppb.New(resp.EndDate),
		PaymentIds:     paymentIds,
	}
}

func SubscriptionToGetActiveSubscriptionResponse(sub *subscription.Subscription) *GetActiveSubscriptionV1Response {
	return &GetActiveSubscriptionV1Response{
		SubscriptionID:     sub.GetID(),
		SubscriptionPlan:   sub.GetPlan(),
		SubscriptionStatus: sub.GetStatus(),
		AutoRenew:          sub.GetAutoRenew(),
		StartDate:          sub.GetStartDate(),
		EndDate:            sub.GetEndDate(),
		PaymentIDs:         sub.GetPaymentIDs(),
	}
}
