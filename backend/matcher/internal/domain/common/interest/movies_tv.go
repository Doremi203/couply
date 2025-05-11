package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type MoviesTV int

const (
	MoviesTVUnspecified MoviesTV = iota
	MoviesTVAction
	MoviesTVComedy
	MoviesTVDrama
	MoviesTVSciFi
	MoviesTVAnime
	MoviesTVDocumentaries
	MoviesTVHorror
	MoviesTVFantasy
	MoviesTVThriller
	MoviesTVRomance
	MoviesTVHistorical
)

func PBToMoviesTV(m desc.MoviesTV) MoviesTV {
	switch m {
	case desc.MoviesTV_MOVIESTV_UNSPECIFIED:
		return MoviesTVUnspecified
	case desc.MoviesTV_MOVIESTV_ACTION:
		return MoviesTVAction
	case desc.MoviesTV_MOVIESTV_COMEDY:
		return MoviesTVComedy
	case desc.MoviesTV_MOVIESTV_DRAMA:
		return MoviesTVDrama
	case desc.MoviesTV_MOVIESTV_SCIFI:
		return MoviesTVSciFi
	case desc.MoviesTV_MOVIESTV_ANIME:
		return MoviesTVAnime
	case desc.MoviesTV_MOVIESTV_DOCUMENTARIES:
		return MoviesTVDocumentaries
	case desc.MoviesTV_MOVIESTV_HORROR:
		return MoviesTVHorror
	case desc.MoviesTV_MOVIESTV_FANTASY:
		return MoviesTVFantasy
	case desc.MoviesTV_MOVIESTV_THRILLER:
		return MoviesTVThriller
	case desc.MoviesTV_MOVIESTV_ROMANCE:
		return MoviesTVRomance
	case desc.MoviesTV_MOVIESTV_HISTORICAL:
		return MoviesTVHistorical
	default:
		return MoviesTVUnspecified
	}
}

func MoviesTVToPB(m MoviesTV) desc.MoviesTV {
	switch m {
	case MoviesTVUnspecified:
		return desc.MoviesTV_MOVIESTV_UNSPECIFIED
	case MoviesTVAction:
		return desc.MoviesTV_MOVIESTV_ACTION
	case MoviesTVComedy:
		return desc.MoviesTV_MOVIESTV_COMEDY
	case MoviesTVDrama:
		return desc.MoviesTV_MOVIESTV_DRAMA
	case MoviesTVSciFi:
		return desc.MoviesTV_MOVIESTV_SCIFI
	case MoviesTVAnime:
		return desc.MoviesTV_MOVIESTV_ANIME
	case MoviesTVDocumentaries:
		return desc.MoviesTV_MOVIESTV_DOCUMENTARIES
	case MoviesTVHorror:
		return desc.MoviesTV_MOVIESTV_HORROR
	case MoviesTVFantasy:
		return desc.MoviesTV_MOVIESTV_FANTASY
	case MoviesTVThriller:
		return desc.MoviesTV_MOVIESTV_THRILLER
	case MoviesTVRomance:
		return desc.MoviesTV_MOVIESTV_ROMANCE
	case MoviesTVHistorical:
		return desc.MoviesTV_MOVIESTV_HISTORICAL
	default:
		return desc.MoviesTV_MOVIESTV_UNSPECIFIED
	}
}

func MoviesTVSliceToPB(m []MoviesTV) []desc.MoviesTV {
	pb := make([]desc.MoviesTV, len(m))
	for i, v := range m {
		pb[i] = MoviesTVToPB(v)
	}
	return pb
}

func PBToMoviesTVSlice(pb []desc.MoviesTV) []MoviesTV {
	m := make([]MoviesTV, len(pb))
	for i, v := range pb {
		m[i] = PBToMoviesTV(v)
	}
	return m
}
