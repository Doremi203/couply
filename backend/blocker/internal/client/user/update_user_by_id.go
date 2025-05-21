package user

import (
	"context"
	"time"

	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

func (c *Client) UpdateUserByIDV1(ctx context.Context, user *userservicegrpc.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := c.client.UpdateUserByIDV1(timeoutCtx, &userservicegrpc.UpdateUserByIDV1Request{
		Id:         user.GetId(),
		Name:       user.GetName(),
		Age:        user.GetAge(),
		Gender:     user.GetGender(),
		Latitude:   user.GetLatitude(),
		Longitude:  user.GetLongitude(),
		Bio:        user.GetBio(),
		Goal:       user.GetGoal(),
		Interest:   user.GetInterest(),
		Zodiac:     user.GetZodiac(),
		Height:     user.GetHeight(),
		Education:  user.GetEducation(),
		Children:   user.GetChildren(),
		Alcohol:    user.GetAlcohol(),
		Smoking:    user.GetSmoking(),
		IsHidden:   user.GetIsHidden(),
		IsVerified: user.GetIsVerified(),
		IsPremium:  user.GetIsPremium(),
		IsBlocked:  user.GetIsBlocked(),
	})
	return err
}
