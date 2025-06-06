package subscription_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/payments/gen/api/subscription-service/v1"
	"github.com/google/uuid"
)

type CancelSubscriptionV1Request struct {
	SubscriptionID uuid.UUID
}

func PBToCancelSubscriptionRequest(req *desc.CancelSubscriptionV1Request) (*CancelSubscriptionV1Request, error) {
	subID, err := uuid.Parse(req.GetSubscriptionId())
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}
	return &CancelSubscriptionV1Request{
		SubscriptionID: subID,
	}, nil
}

type CancelSubscriptionV1Response struct{}

func CancelSubscriptionResponseToPB(_ *CancelSubscriptionV1Response) *desc.CancelSubscriptionV1Response {
	return &desc.CancelSubscriptionV1Response{}
}
