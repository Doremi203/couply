package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type SelfDevelopment int

const (
	SelfDevelopmentUnspecified SelfDevelopment = iota
	SelfDevelopmentLanguages
	SelfDevelopmentLectures
	SelfDevelopmentOnlineCourses
	SelfDevelopmentSelfEducation
	SelfDevelopmentMeditation
	SelfDevelopmentPsychology
	SelfDevelopmentPhilosophy
	SelfDevelopmentHistory
	SelfDevelopmentTechnology
	SelfDevelopmentReading
)

func PBToSelfDevelopment(sd desc.SelfDevelopment) SelfDevelopment {
	switch sd {
	case desc.SelfDevelopment_SELFDEVELOPMENT_UNSPECIFIED:
		return SelfDevelopmentUnspecified
	case desc.SelfDevelopment_SELFDEVELOPMENT_LANGUAGES:
		return SelfDevelopmentLanguages
	case desc.SelfDevelopment_SELFDEVELOPMENT_LECTURES:
		return SelfDevelopmentLectures
	case desc.SelfDevelopment_SELFDEVELOPMENT_ONLINE_COURSES:
		return SelfDevelopmentOnlineCourses
	case desc.SelfDevelopment_SELFDEVELOPMENT_SELF_EDUCATION:
		return SelfDevelopmentSelfEducation
	case desc.SelfDevelopment_SELFDEVELOPMENT_MEDITATION:
		return SelfDevelopmentMeditation
	case desc.SelfDevelopment_SELFDEVELOPMENT_PSYCHOLOGY:
		return SelfDevelopmentPsychology
	case desc.SelfDevelopment_SELFDEVELOPMENT_PHILOSOPHY:
		return SelfDevelopmentPhilosophy
	case desc.SelfDevelopment_SELFDEVELOPMENT_HISTORY:
		return SelfDevelopmentHistory
	case desc.SelfDevelopment_SELFDEVELOPMENT_TECHNOLOGY:
		return SelfDevelopmentTechnology
	case desc.SelfDevelopment_SELFDEVELOPMENT_READING:
		return SelfDevelopmentReading
	default:
		return SelfDevelopmentUnspecified
	}
}

func SelfDevelopmentToPB(sd SelfDevelopment) desc.SelfDevelopment {
	switch sd {
	case SelfDevelopmentUnspecified:
		return desc.SelfDevelopment_SELFDEVELOPMENT_UNSPECIFIED
	case SelfDevelopmentLanguages:
		return desc.SelfDevelopment_SELFDEVELOPMENT_LANGUAGES
	case SelfDevelopmentLectures:
		return desc.SelfDevelopment_SELFDEVELOPMENT_LECTURES
	case SelfDevelopmentOnlineCourses:
		return desc.SelfDevelopment_SELFDEVELOPMENT_ONLINE_COURSES
	case SelfDevelopmentSelfEducation:
		return desc.SelfDevelopment_SELFDEVELOPMENT_SELF_EDUCATION
	case SelfDevelopmentMeditation:
		return desc.SelfDevelopment_SELFDEVELOPMENT_MEDITATION
	case SelfDevelopmentPsychology:
		return desc.SelfDevelopment_SELFDEVELOPMENT_PSYCHOLOGY
	case SelfDevelopmentPhilosophy:
		return desc.SelfDevelopment_SELFDEVELOPMENT_PHILOSOPHY
	case SelfDevelopmentHistory:
		return desc.SelfDevelopment_SELFDEVELOPMENT_HISTORY
	case SelfDevelopmentTechnology:
		return desc.SelfDevelopment_SELFDEVELOPMENT_TECHNOLOGY
	case SelfDevelopmentReading:
		return desc.SelfDevelopment_SELFDEVELOPMENT_READING
	default:
		return desc.SelfDevelopment_SELFDEVELOPMENT_UNSPECIFIED
	}
}

func SelfDevelopmentSliceToPB(s []SelfDevelopment) []desc.SelfDevelopment {
	pb := make([]desc.SelfDevelopment, len(s))
	for i, v := range s {
		pb[i] = SelfDevelopmentToPB(v)
	}
	return pb
}

func PBToSelfDevelopmentSlice(pb []desc.SelfDevelopment) []SelfDevelopment {
	s := make([]SelfDevelopment, len(pb))
	for i, v := range pb {
		s[i] = PBToSelfDevelopment(v)
	}
	return s
}
