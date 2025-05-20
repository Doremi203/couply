package subscription_service

import (
	desc "github.com/Doremi203/couply/backend/payment/gen/api/subscription-service/v1"
	"github.com/google/uuid"
)

type CancelSubscriptionV1Request struct {
	SubscriptionID uuid.UUID
}

func (x *CancelSubscriptionV1Request) GetSubscriptionID() uuid.UUID {
	if x != nil {
		return x.SubscriptionID
	}
	return uuid.Nil
}

func PBToCancelSubscriptionRequest(req *desc.CancelSubscriptionV1Request) (*CancelSubscriptionV1Request, error) {
	subID, err := uuid.Parse(req.GetSubscriptionId())
	if err != nil {
		return nil, err
	}
	return &CancelSubscriptionV1Request{
		SubscriptionID: subID,
	}, nil
}

type CancelSubscriptionV1Response struct{}

func CancelSubscriptionResponseToPB(_ *CancelSubscriptionV1Response) *desc.CancelSubscriptionV1Response {
	return &desc.CancelSubscriptionV1Response{}
}
