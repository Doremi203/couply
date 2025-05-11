package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type Music int

const (
	MusicUnspecified Music = iota
	MusicPop
	MusicRock
	MusicHipHop
	MusicRap
	MusicElectronic
	MusicJazz
	MusicClassical
	MusicIndie
	MusicRNB
	MusicMetal
	MusicFolk
	MusicCountry
	MusicAlternative
)

func PBToMusic(m desc.Music) Music {
	switch m {
	case desc.Music_MUSIC_UNSPECIFIED:
		return MusicUnspecified
	case desc.Music_MUSIC_POP:
		return MusicPop
	case desc.Music_MUSIC_ROCK:
		return MusicRock
	case desc.Music_MUSIC_HIPHOP:
		return MusicHipHop
	case desc.Music_MUSIC_RAP:
		return MusicRap
	case desc.Music_MUSIC_ELECTRONIC:
		return MusicElectronic
	case desc.Music_MUSIC_JAZZ:
		return MusicJazz
	case desc.Music_MUSIC_CLASSICAL:
		return MusicClassical
	case desc.Music_MUSIC_INDIE:
		return MusicIndie
	case desc.Music_MUSIC_RNB:
		return MusicRNB
	case desc.Music_MUSIC_METAL:
		return MusicMetal
	case desc.Music_MUSIC_FOLK:
		return MusicFolk
	case desc.Music_MUSIC_COUNTRY:
		return MusicCountry
	case desc.Music_MUSIC_ALTERNATIVE:
		return MusicAlternative
	default:
		return MusicUnspecified
	}
}

func MusicToPB(m Music) desc.Music {
	switch m {
	case MusicUnspecified:
		return desc.Music_MUSIC_UNSPECIFIED
	case MusicPop:
		return desc.Music_MUSIC_POP
	case MusicRock:
		return desc.Music_MUSIC_ROCK
	case MusicHipHop:
		return desc.Music_MUSIC_HIPHOP
	case MusicRap:
		return desc.Music_MUSIC_RAP
	case MusicElectronic:
		return desc.Music_MUSIC_ELECTRONIC
	case MusicJazz:
		return desc.Music_MUSIC_JAZZ
	case MusicClassical:
		return desc.Music_MUSIC_CLASSICAL
	case MusicIndie:
		return desc.Music_MUSIC_INDIE
	case MusicRNB:
		return desc.Music_MUSIC_RNB
	case MusicMetal:
		return desc.Music_MUSIC_METAL
	case MusicFolk:
		return desc.Music_MUSIC_FOLK
	case MusicCountry:
		return desc.Music_MUSIC_COUNTRY
	case MusicAlternative:
		return desc.Music_MUSIC_ALTERNATIVE
	default:
		return desc.Music_MUSIC_UNSPECIFIED
	}
}

func MusicSliceToPB(m []Music) []desc.Music {
	pb := make([]desc.Music, len(m))
	for i, v := range m {
		pb[i] = MusicToPB(v)
	}
	return pb
}

func PBToMusicSlice(pb []desc.Music) []Music {
	m := make([]Music, len(pb))
	for i, v := range pb {
		m[i] = PBToMusic(v)
	}
	return m
}
