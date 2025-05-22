package subscription_service

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/subscription-service/v1"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateSubscriptionV1Request struct {
	SubscriptionPlan subscription.SubscriptionPlan
	AutoRenew        bool
}

func PBToCreateSubscriptionRequest(req *desc.CreateSubscriptionV1Request) *CreateSubscriptionV1Request {
	return &CreateSubscriptionV1Request{
		SubscriptionPlan: subscription.PBToSubscriptionPlan(req.GetPlan()),
		AutoRenew:        req.GetAutoRenew(),
	}
}

func CreateSubscriptionRequestToSubscription(req *CreateSubscriptionV1Request, userID uuid.UUID) (*subscription.Subscription, error) {
	subID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "CreateSubscriptionRequestToSubscription")
	}

	now := time.Now()
	plan := req.SubscriptionPlan
	return &subscription.Subscription{
		ID:        subID,
		UserID:    userID,
		Plan:      plan,
		Status:    subscription.SubscriptionStatusPendingPayment,
		AutoRenew: req.AutoRenew,
		StartDate: now,
		EndDate:   subscription.CalculateEndDate(now, plan),
	}, nil
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

func CreateSubscriptionResponseToPB(resp *CreateSubscriptionV1Response) *desc.CreateSubscriptionV1Response {
	paymentIds := make([]string, 0, len(resp.PaymentIDs))
	for _, id := range resp.PaymentIDs {
		paymentIds = append(paymentIds, id.String())
	}
	return &desc.CreateSubscriptionV1Response{
		SubscriptionId: resp.SubscriptionID.String(),
		Plan:           subscription.SubscriptionPlanToPB(resp.SubscriptionPlan),
		Status:         subscription.SubscriptionStatusToPB(resp.SubscriptionStatus),
		AutoRenew:      resp.AutoRenew,
		StartDate:      timestamppb.New(resp.StartDate),
		EndDate:        timestamppb.New(resp.EndDate),
		PaymentIds:     paymentIds,
	}
}

func SubscriptionToCreateSubscriptionResponse(sub *subscription.Subscription) *CreateSubscriptionV1Response {
	return &CreateSubscriptionV1Response{
		SubscriptionID:     sub.GetID(),
		SubscriptionPlan:   sub.GetPlan(),
		SubscriptionStatus: sub.GetStatus(),
		AutoRenew:          sub.GetAutoRenew(),
		StartDate:          sub.GetStartDate(),
		EndDate:            sub.GetEndDate(),
		PaymentIDs:         sub.GetPaymentIDs(),
	}
}
