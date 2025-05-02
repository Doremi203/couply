package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type Art int

const (
	ArtUnspecified Art = iota
	ArtPainting
	ArtPhotograph
	ArtMusic
	ArtSinging
	ArtWriting
	ArtSculpture
	ArtTheater
	ArtCinema
	ArtNeedlework
)

func PBToArt(art desc.Art) Art {
	switch art {
	case desc.Art_ART_UNSPECIFIED:
		return ArtUnspecified
	case desc.Art_ART_PAINTING:
		return ArtPainting
	case desc.Art_ART_PHOTOGRAPH:
		return ArtPhotograph
	case desc.Art_ART_MUSIC:
		return ArtMusic
	case desc.Art_ART_SINGING:
		return ArtSinging
	case desc.Art_ART_WRITING:
		return ArtWriting
	case desc.Art_ART_SCULPTURE:
		return ArtSculpture
	case desc.Art_ART_THEATER:
		return ArtTheater
	case desc.Art_ART_CINEMA:
		return ArtCinema
	case desc.Art_ART_NEEDLEWORK:
		return ArtNeedlework
	default:
		return Art(0)
	}
}

func ArtToPB(art Art) desc.Art {
	switch art {
	case ArtUnspecified:
		return desc.Art_ART_UNSPECIFIED
	case ArtPainting:
		return desc.Art_ART_PAINTING
	case ArtPhotograph:
		return desc.Art_ART_PHOTOGRAPH
	case ArtMusic:
		return desc.Art_ART_MUSIC
	case ArtSinging:
		return desc.Art_ART_SINGING
	case ArtWriting:
		return desc.Art_ART_WRITING
	case ArtSculpture:
		return desc.Art_ART_SCULPTURE
	case ArtTheater:
		return desc.Art_ART_THEATER
	case ArtCinema:
		return desc.Art_ART_CINEMA
	case ArtNeedlework:
		return desc.Art_ART_NEEDLEWORK
	default:
		return desc.Art(0)
	}
}

func ArtSliceToPB(arts []Art) []desc.Art {
	artsPB := make([]desc.Art, 0, len(arts))

	for _, art := range arts {
		artsPB = append(artsPB, ArtToPB(art))
	}

	return artsPB
}

func PBToArtSlice(arts []desc.Art) []Art {
	artsDomain := make([]Art, 0, len(arts))

	for _, art := range arts {
		artsDomain = append(artsDomain, PBToArt(art))
	}

	return artsDomain
}
