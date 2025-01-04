package domain

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

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

func PBToSelfDevelopment(selfDevelopment desc.Selfdevelopment) SelfDevelopment {
	switch selfDevelopment {
	case desc.Selfdevelopment_SELFDEVELOPMENT_UNSPECIFIED:
		return SelfDevelopmentUnspecified
	case desc.Selfdevelopment_SELFDEVELOPMENT_LANGUAGES:
		return SelfDevelopmentLanguages
	case desc.Selfdevelopment_SELFDEVELOPMENT_LECTURES:
		return SelfDevelopmentLectures
	case desc.Selfdevelopment_SELFDEVELOPMENT_ONLINE_COURSES:
		return SelfDevelopmentOnlineCourses
	case desc.Selfdevelopment_SELFDEVELOPMENT_SELF_EDUCATION:
		return SelfDevelopmentSelfEducation
	case desc.Selfdevelopment_SELFDEVELOPMENT_MEDITATION:
		return SelfDevelopmentMeditation
	case desc.Selfdevelopment_SELFDEVELOPMENT_PSYCHOLOGY:
		return SelfDevelopmentPsychology
	case desc.Selfdevelopment_SELFDEVELOPMENT_PHILOSOPHY:
		return SelfDevelopmentPhilosophy
	case desc.Selfdevelopment_SELFDEVELOPMENT_HISTORY:
		return SelfDevelopmentHistory
	case desc.Selfdevelopment_SELFDEVELOPMENT_TECHNOLOGY:
		return SelfDevelopmentTechnology
	case desc.Selfdevelopment_SELFDEVELOPMENT_READING:
		return SelfDevelopmentReading
	default:
		return SelfDevelopment(0)
	}
}

func SelfDevelopmentToPB(selfDevelopment SelfDevelopment) desc.Selfdevelopment {
	switch selfDevelopment {
	case SelfDevelopmentUnspecified:
		return desc.Selfdevelopment_SELFDEVELOPMENT_UNSPECIFIED
	case SelfDevelopmentLanguages:
		return desc.Selfdevelopment_SELFDEVELOPMENT_LANGUAGES
	case SelfDevelopmentLectures:
		return desc.Selfdevelopment_SELFDEVELOPMENT_LECTURES
	case SelfDevelopmentOnlineCourses:
		return desc.Selfdevelopment_SELFDEVELOPMENT_ONLINE_COURSES
	case SelfDevelopmentSelfEducation:
		return desc.Selfdevelopment_SELFDEVELOPMENT_SELF_EDUCATION
	case SelfDevelopmentMeditation:
		return desc.Selfdevelopment_SELFDEVELOPMENT_MEDITATION
	case SelfDevelopmentPsychology:
		return desc.Selfdevelopment_SELFDEVELOPMENT_PSYCHOLOGY
	case SelfDevelopmentPhilosophy:
		return desc.Selfdevelopment_SELFDEVELOPMENT_PHILOSOPHY
	case SelfDevelopmentHistory:
		return desc.Selfdevelopment_SELFDEVELOPMENT_HISTORY
	case SelfDevelopmentTechnology:
		return desc.Selfdevelopment_SELFDEVELOPMENT_TECHNOLOGY
	case SelfDevelopmentReading:
		return desc.Selfdevelopment_SELFDEVELOPMENT_READING
	default:
		return desc.Selfdevelopment(0)
	}
}
