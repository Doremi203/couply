package matching

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"

type Status int

const (
	StatusUnspecified Status = iota
	StatusWaiting
	StatusAccepted
	StatusDeclined
)

func PBToStatus(status desc.Status) Status {
	switch status {
	case desc.Status_STATUS_UNSPECIFIED:
		return StatusUnspecified
	case desc.Status_STATUS_WAITING:
		return StatusWaiting
	case desc.Status_STATUS_ACCEPTED:
		return StatusAccepted
	case desc.Status_STATUS_DECLINED:
		return StatusDeclined
	default:
		return Status(0)
	}
}

func StatusToPB(status Status) desc.Status {
	switch status {
	case StatusUnspecified:
		return desc.Status_STATUS_UNSPECIFIED
	case StatusWaiting:
		return desc.Status_STATUS_WAITING
	case StatusAccepted:
		return desc.Status_STATUS_DECLINED
	case StatusDeclined:
		return desc.Status_STATUS_DECLINED
	default:
		return desc.Status(0)
	}
}
