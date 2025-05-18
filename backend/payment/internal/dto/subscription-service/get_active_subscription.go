package subscription_service

import (
	desc "github.com/Doremi203/couply/backend/payment/gen/api/subscription-service/v1"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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

func (x *GetActiveSubscriptionV1Response) GetSubscriptionID() uuid.UUID {
	if x != nil {
		return x.SubscriptionID
	}
	return uuid.Nil
}

func (x *GetActiveSubscriptionV1Response) GetSubscriptionPlan() subscription.SubscriptionPlan {
	if x != nil {
		return x.SubscriptionPlan
	}
	return subscription.SubscriptionPlan(0)
}

func (x *GetActiveSubscriptionV1Response) GetSubscriptionStatus() subscription.SubscriptionStatus {
	if x != nil {
		return x.SubscriptionStatus
	}
	return subscription.SubscriptionStatus(0)
}

func (x *GetActiveSubscriptionV1Response) GetAutoRenew() bool {
	if x != nil {
		return x.AutoRenew
	}
	return false
}

func (x *GetActiveSubscriptionV1Response) GetStartDate() time.Time {
	if x != nil {
		return x.StartDate
	}
	return time.Time{}
}

func (x *GetActiveSubscriptionV1Response) GetEndDate() time.Time {
	if x != nil {
		return x.EndDate
	}
	return time.Time{}
}

func (x *GetActiveSubscriptionV1Response) GetPaymentIDs() []uuid.UUID {
	if x != nil {
		return x.PaymentIDs
	}
	return nil
}

func GetActiveSubscriptionResponseToPB(resp *GetActiveSubscriptionV1Response) *desc.GetActiveSubscriptionV1Response {
	paymentIds := make([]string, 0, len(resp.GetPaymentIDs()))
	for _, id := range resp.GetPaymentIDs() {
		paymentIds = append(paymentIds, id.String())
	}
	return &desc.GetActiveSubscriptionV1Response{
		SubscriptionId: resp.GetSubscriptionID().String(),
		Plan:           subscription.SubscriptionPlanToPB(resp.GetSubscriptionPlan()),
		Status:         subscription.SubscriptionStatusToPB(resp.GetSubscriptionStatus()),
		AutoRenew:      resp.GetAutoRenew(),
		StartDate:      timestamppb.New(resp.GetStartDate()),
		EndDate:        timestamppb.New(resp.GetEndDate()),
		PaymentIds:     paymentIds,
	}
}
