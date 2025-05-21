package blocker

import desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"

type BlockStatus int

const (
	BlockStatusUnspecified BlockStatus = iota
	BlockStatusPending
	BlockStatusAccepted
	BlockStatusDeclined
)

func PBToBlockStatus(blockStatus desc.BlockStatus) BlockStatus {
	switch blockStatus {
	case desc.BlockStatus_BLOCK_STATUS_UNSPECIFIED:
		return BlockStatusUnspecified
	case desc.BlockStatus_BLOCK_STATUS_PENDING:
		return BlockStatusPending
	case desc.BlockStatus_BLOCK_STATUS_ACCEPTED:
		return BlockStatusAccepted
	case desc.BlockStatus_BLOCK_STATUS_DECLINED:
		return BlockStatusDeclined
	default:
		return BlockStatus(0)
	}
}

func BlockStatusToPB(blockStatus BlockStatus) desc.BlockStatus {
	switch blockStatus {
	case BlockStatusUnspecified:
		return desc.BlockStatus_BLOCK_STATUS_UNSPECIFIED
	case BlockStatusPending:
		return desc.BlockStatus_BLOCK_STATUS_PENDING
	case BlockStatusAccepted:
		return desc.BlockStatus_BLOCK_STATUS_ACCEPTED
	case BlockStatusDeclined:
		return desc.BlockStatus_BLOCK_STATUS_DECLINED
	default:
		return desc.BlockStatus(0)
	}
}
