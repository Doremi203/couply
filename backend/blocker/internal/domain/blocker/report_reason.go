package blocker

import desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"

type ReportReason int

const (
	ReportReasonUnspecified ReportReason = iota
	ReportReasonFakeProfile
	ReportReasonSpam
	ReportReasonAbuse
	ReportReasonInappropriateContent
	ReportReasonAge
	ReportReasonOther
)

func (r ReportReason) String() string {
	switch r {
	case ReportReasonUnspecified:
		return "REASON_UNSPECIFIED"
	case ReportReasonFakeProfile:
		return "REASON_FAKE_PROFILE"
	case ReportReasonSpam:
		return "REASON_SPAM"
	case ReportReasonAbuse:
		return "REASON_ABUSE"
	case ReportReasonInappropriateContent:
		return "REASON_INAPPROPRIATE_CONTENT"
	case ReportReasonAge:
		return "REASON_AGE"
	case ReportReasonOther:
		return "REASON_OTHER"
	default:
		return "UNKNOWN_REASON"
	}
}

func PBToReportReason(reportReason desc.ReportReason) ReportReason {
	switch reportReason {
	case desc.ReportReason_REASON_UNSPECIFIED:
		return ReportReasonUnspecified
	case desc.ReportReason_REASON_FAKE_PROFILE:
		return ReportReasonFakeProfile
	case desc.ReportReason_REASON_SPAM:
		return ReportReasonSpam
	case desc.ReportReason_REASON_ABUSE:
		return ReportReasonAbuse
	case desc.ReportReason_REASON_INAPPROPRIATE_CONTENT:
		return ReportReasonInappropriateContent
	case desc.ReportReason_REASON_AGE:
		return ReportReasonAge
	case desc.ReportReason_REASON_OTHER:
		return ReportReasonOther
	default:
		return ReportReason(0)
	}
}

func ReportReasonToPB(reportReason ReportReason) desc.ReportReason {
	switch reportReason {
	case ReportReasonUnspecified:
		return desc.ReportReason_REASON_UNSPECIFIED
	case ReportReasonFakeProfile:
		return desc.ReportReason_REASON_FAKE_PROFILE
	case ReportReasonSpam:
		return desc.ReportReason_REASON_SPAM
	case ReportReasonAbuse:
		return desc.ReportReason_REASON_ABUSE
	case ReportReasonInappropriateContent:
		return desc.ReportReason_REASON_INAPPROPRIATE_CONTENT
	case ReportReasonAge:
		return desc.ReportReason_REASON_AGE
	case ReportReasonOther:
		return desc.ReportReason_REASON_OTHER
	default:
		return desc.ReportReason(0)
	}
}
