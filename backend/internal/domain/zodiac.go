package domain

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

type Zodiac int

const (
	ZodiacUnspecified Zodiac = iota
	ZodiacAries
	ZodiacTaurus
	ZodiacGemini
	ZodiacCancer
	ZodiacLeo
	ZodiacVirgo
	ZodiacLibra
	ZodiacScorpio
	ZodiacSagittarius
	ZodiacCapricorn
	ZodiacAquarius
	ZodiacPisces
)

func PBToZodiac(zodiac desc.Zodiac) Zodiac {
	switch zodiac {
	case desc.Zodiac_ZODIAC_UNSPECIFIED:
		return ZodiacUnspecified
	case desc.Zodiac_ZODIAC_ARIES:
		return ZodiacAries
	case desc.Zodiac_ZODIAC_TAURUS:
		return ZodiacTaurus
	case desc.Zodiac_ZODIAC_GEMINI:
		return ZodiacGemini
	case desc.Zodiac_ZODIAC_CANCER:
		return ZodiacCancer
	case desc.Zodiac_ZODIAC_LEO:
		return ZodiacLeo
	case desc.Zodiac_ZODIAC_VIRGO:
		return ZodiacVirgo
	case desc.Zodiac_ZODIAC_LIBRA:
		return ZodiacLibra
	case desc.Zodiac_ZODIAC_SCORPIO:
		return ZodiacScorpio
	case desc.Zodiac_ZODIAC_SAGITTARIUS:
		return ZodiacSagittarius
	case desc.Zodiac_ZODIAC_CAPRICORN:
		return ZodiacCapricorn
	case desc.Zodiac_ZODIAC_AQUARIUS:
		return ZodiacAquarius
	case desc.Zodiac_ZODIAC_PISCES:
		return ZodiacPisces
	default:
		return Zodiac(0)
	}
}

func ZodiacToPB(zodiac Zodiac) desc.Zodiac {
	switch zodiac {
	case ZodiacUnspecified:
		return desc.Zodiac_ZODIAC_UNSPECIFIED
	case ZodiacAries:
		return desc.Zodiac_ZODIAC_ARIES
	case ZodiacTaurus:
		return desc.Zodiac_ZODIAC_TAURUS
	case ZodiacGemini:
		return desc.Zodiac_ZODIAC_GEMINI
	case ZodiacCancer:
		return desc.Zodiac_ZODIAC_CANCER
	case ZodiacLeo:
		return desc.Zodiac_ZODIAC_LEO
	case ZodiacVirgo:
		return desc.Zodiac_ZODIAC_VIRGO
	case ZodiacLibra:
		return desc.Zodiac_ZODIAC_LIBRA
	case ZodiacScorpio:
		return desc.Zodiac_ZODIAC_SCORPIO
	case ZodiacSagittarius:
		return desc.Zodiac_ZODIAC_SAGITTARIUS
	case ZodiacCapricorn:
		return desc.Zodiac_ZODIAC_CAPRICORN
	case ZodiacAquarius:
		return desc.Zodiac_ZODIAC_AQUARIUS
	case ZodiacPisces:
		return desc.Zodiac_ZODIAC_PISCES
	default:
		return desc.Zodiac(0)
	}
}
