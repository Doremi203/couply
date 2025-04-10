package search_service

import (
	"github.com/Doremi203/Couply/backend/internal/domain/search"
	"github.com/Doremi203/Couply/backend/internal/domain/user"
	"github.com/Doremi203/Couply/backend/internal/domain/user/interest"
	desc "github.com/Doremi203/Couply/backend/pkg/search-service/v1"
)

type CreateFilterV1Request struct {
	GenderPriority search.GenderPriority
	MaxAge         int32
	Distance       int32
	Goal           user.Goal
	SearchTape     *search.SearchTape
	Interest       *interest.Interest
	Info           *search.Info
}

type CreateFilterV1Response struct {
	Filter *search.Filter
}

func CreateFilterRequestToPB(req *CreateFilterV1Request) *desc.CreateFilterV1Request {
	return &desc.CreateFilterV1Request{
		GenderPriority: search.GenderPriorityToPB(req.GenderPriority),
		MaxAge:         req.MaxAge,
		Distance:       req.Distance,
		Goal:           user.GoalToPB(req.Goal),
		SearchTape:     search.SearchTapeToPB(req.SearchTape),
		Interest:       interest.InterestToPB(req.Interest),
		Info:           search.InfoToPB(req.Info),
	}
}

func PBToCreateFilterRequest(req *desc.CreateFilterV1Request) *CreateFilterV1Request {
	return &CreateFilterV1Request{
		GenderPriority: search.PBToGenderPriority(req.GenderPriority),
		MaxAge:         req.MaxAge,
		Distance:       req.Distance,
		Goal:           user.PBToGoal(req.Goal),
		SearchTape:     search.PBToSearchTape(req.SearchTape),
		Interest:       interest.PBToInterest(req.Interest),
		Info:           search.PBToInfo(req.Info),
	}
}

func CreateFilterResponseToPB(resp *CreateFilterV1Response) *desc.CreateFilterV1Response {
	return &desc.CreateFilterV1Response{
		Filter: search.FilterToPB(resp.Filter),
	}
}

func PBToCreateFilterResponse(resp *desc.CreateFilterV1Response) *CreateFilterV1Response {
	return &CreateFilterV1Response{
		Filter: search.PBToFilter(resp.Filter),
	}
}

func CreateFilterRequestToFilter(req *CreateFilterV1Request) *search.Filter {
	return &search.Filter{
		GenderPriority: req.GenderPriority,
		MaxAge:         req.MaxAge,
		Distance:       req.Distance,
		Goal:           req.Goal,
		SearchTape:     req.SearchTape,
		Interest:       req.Interest,
		Info:           req.Info,
	}
}
