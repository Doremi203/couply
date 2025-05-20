package subscription_service

import (
	"time"

	desc "github.com/Doremi203/couply/backend/payment/gen/api/subscription-service/v1"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateSubscriptionV1Request struct {
	SubscriptionPlan subscription.SubscriptionPlan
	AutoRenew        bool
}

func (x *CreateSubscriptionV1Request) GetSubscriptionPlan() subscription.SubscriptionPlan {
	if x != nil {
		return x.SubscriptionPlan
	}
	return subscription.SubscriptionPlan(0)
}

func (x *CreateSubscriptionV1Request) GetAutoRenew() bool {
	if x != nil {
		return x.AutoRenew
	}
	return false
}

func PBToCreateSubscriptionRequest(req *desc.CreateSubscriptionV1Request) *CreateSubscriptionV1Request {
	return &CreateSubscriptionV1Request{
		SubscriptionPlan: subscription.PBToSubscriptionPlan(req.GetPlan()),
		AutoRenew:        req.GetAutoRenew(),
	}
}

type CreateSubscriptionV1Response struct {
	SubscriptionID     uuid.UUID
	SubscriptionPlan   subscription.SubscriptionPlan
	SubscriptionStatus subscription.SubscriptionStatus
	AutoRenew          bool
	StartDate          time.Time
	EndDate            time.Time
	PaymentIDs         []uuid.UUID
}

func (x *CreateSubscriptionV1Response) GetSubscriptionID() uuid.UUID {
	if x != nil {
		return x.SubscriptionID
	}
	return uuid.Nil
}

func (x *CreateSubscriptionV1Response) GetSubscriptionPlan() subscription.SubscriptionPlan {
	if x != nil {
		return x.SubscriptionPlan
	}
	return subscription.SubscriptionPlan(0)
}

func (x *CreateSubscriptionV1Response) GetSubscriptionStatus() subscription.SubscriptionStatus {
	if x != nil {
		return x.SubscriptionStatus
	}
	return subscription.SubscriptionStatus(0)
}

func (x *CreateSubscriptionV1Response) GetAutoRenew() bool {
	if x != nil {
		return x.AutoRenew
	}
	return false
}

func (x *CreateSubscriptionV1Response) GetStartDate() time.Time {
	if x != nil {
		return x.StartDate
	}
	return time.Time{}
}

func (x *CreateSubscriptionV1Response) GetEndDate() time.Time {
	if x != nil {
		return x.EndDate
	}
	return time.Time{}
}

func (x *CreateSubscriptionV1Response) GetPaymentIDs() []uuid.UUID {
	if x != nil {
		return x.PaymentIDs
	}
	return nil
}

func CreateSubscriptionResponseToPB(resp *CreateSubscriptionV1Response) *desc.CreateSubscriptionV1Response {
	paymentIds := make([]string, 0, len(resp.GetPaymentIDs()))
	for _, id := range resp.GetPaymentIDs() {
		paymentIds = append(paymentIds, id.String())
	}
	return &desc.CreateSubscriptionV1Response{
		SubscriptionId: resp.GetSubscriptionID().String(),
		Plan:           subscription.SubscriptionPlanToPB(resp.GetSubscriptionPlan()),
		Status:         subscription.SubscriptionStatusToPB(resp.GetSubscriptionStatus()),
		AutoRenew:      resp.GetAutoRenew(),
		StartDate:      timestamppb.New(resp.GetStartDate()),
		EndDate:        timestamppb.New(resp.GetEndDate()),
		PaymentIds:     paymentIds,
	}
}
