package search

import (
	"github.com/Doremi203/Couply/backend/internal/domain/user"
	search_service "github.com/Doremi203/Couply/backend/pkg/search-service/v1"
)

type Info struct {
	MinHeight int32
	MaxHeight int32
	Zodiac    user.Zodiac
	Education user.Education
	Children  user.Children
	Alcohol   user.Alcohol
	Smoking   user.Smoking
}

func InfoToPB(info *Info) *search_service.Info {
	return &search_service.Info{
		MinHeight: info.MinHeight,
		MaxHeight: info.MaxHeight,
		Zodiac:    user.ZodiacToPB(info.Zodiac),
		Education: user.EducationToPB(info.Education),
		Children:  user.ChildrenToPB(info.Children),
		Alcohol:   user.AlcoholToPB(info.Alcohol),
		Smoking:   user.SmokingToPB(info.Smoking),
	}
}

func PBToInfo(info *search_service.Info) *Info {
	return &Info{
		MinHeight: info.MinHeight,
		MaxHeight: info.MaxHeight,
		Zodiac:    user.PBToZodiac(info.Zodiac),
		Education: user.PBToEducation(info.Education),
		Children:  user.PBToChildren(info.Children),
		Alcohol:   user.PBToAlcohol(info.Alcohol),
		Smoking:   user.PBToSmoking(info.Smoking),
	}
}
