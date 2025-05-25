package user

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/Doremi203/couply/backend/common/libs/slices"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrDuplicateUser   = errors.Error("user already exists")
	ErrUserDoesntExist = errors.Error("user does not exist")
	ErrUserNotFound    = errors.Error("user not found")
	ErrUsersNotFound   = errors.Error("users not found")
)

type User struct {
	ID         uuid.UUID          `db:"id"`
	Name       string             `db:"name"`
	Age        int32              `db:"age"`
	Gender     Gender             `db:"gender"`
	Latitude   float64            `db:"latitude"`
	Longitude  float64            `db:"longitude"`
	BIO        string             `db:"bio"`
	Goal       common.Goal        `db:"goal"`
	Interest   *interest.Interest `db:"-"`
	Zodiac     common.Zodiac      `db:"zodiac"`
	Height     int32              `db:"height"`
	Education  common.Education   `db:"education"`
	Children   common.Children    `db:"children"`
	Alcohol    common.Alcohol     `db:"alcohol"`
	Smoking    common.Smoking     `db:"smoking"`
	IsHidden   bool               `db:"is_hidden"`
	IsVerified bool               `db:"is_verified"`
	IsPremium  bool               `db:"is_premium"`
	IsBlocked  bool               `db:"is_blocked"`
	Photos     []Photo            `db:"-"`
	CreatedAt  time.Time          `db:"created_at"`
	UpdatedAt  time.Time          `db:"updated_at"`
}

func (x *User) GenerateDownloadPhotoURLS(ctx context.Context, gen PhotoURLGenerator) error {
	var errs []error
	for i := range x.Photos {
		if err := x.Photos[i].GetDownloadURL(ctx, gen); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: &User{}}
}

func (b *UserBuilder) SetID(id uuid.UUID) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) SetAge(age int32) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) SetGender(gender Gender) *UserBuilder {
	b.user.Gender = gender
	return b
}

func (b *UserBuilder) SetLatitude(latitude float64) *UserBuilder {
	b.user.Latitude = latitude
	return b
}

func (b *UserBuilder) SetLongitude(longitude float64) *UserBuilder {
	b.user.Longitude = longitude
	return b
}

func (b *UserBuilder) SetBIO(bio string) *UserBuilder {
	b.user.BIO = bio
	return b
}

func (b *UserBuilder) SetGoal(goal common.Goal) *UserBuilder {
	b.user.Goal = goal
	return b
}

func (b *UserBuilder) SetInterest(interest *interest.Interest) *UserBuilder {
	b.user.Interest = interest
	return b
}

func (b *UserBuilder) SetZodiac(zodiac common.Zodiac) *UserBuilder {
	b.user.Zodiac = zodiac
	return b
}

func (b *UserBuilder) SetHeight(height int32) *UserBuilder {
	b.user.Height = height
	return b
}

func (b *UserBuilder) SetEducation(education common.Education) *UserBuilder {
	b.user.Education = education
	return b
}

func (b *UserBuilder) SetChildren(children common.Children) *UserBuilder {
	b.user.Children = children
	return b
}

func (b *UserBuilder) SetAlcohol(alcohol common.Alcohol) *UserBuilder {
	b.user.Alcohol = alcohol
	return b
}

func (b *UserBuilder) SetSmoking(smoking common.Smoking) *UserBuilder {
	b.user.Smoking = smoking
	return b
}

func (b *UserBuilder) SetIsHidden(isHidden bool) *UserBuilder {
	b.user.IsHidden = isHidden
	return b
}

func (b *UserBuilder) SetIsVerified(isVerified bool) *UserBuilder {
	b.user.IsVerified = isVerified
	return b
}

func (b *UserBuilder) SetIsPremium(isPremium bool) *UserBuilder {
	b.user.IsPremium = isPremium
	return b
}

func (b *UserBuilder) SetIsBlocked(isBlocked bool) *UserBuilder {
	b.user.IsBlocked = isBlocked
	return b
}

func (b *UserBuilder) SetPhotos(photos []Photo) *UserBuilder {
	b.user.Photos = photos
	return b
}

func (b *UserBuilder) SetCreatedAt(createdAt time.Time) *UserBuilder {
	b.user.CreatedAt = createdAt
	return b
}

func (b *UserBuilder) SetUpdatedAt(updatedAt time.Time) *UserBuilder {
	b.user.UpdatedAt = updatedAt
	return b
}

func (b *UserBuilder) Build() *User {
	return b.user
}

func UserToPB(user *User) *desc.User {
	return &desc.User{
		Id:         user.ID.String(),
		Name:       user.Name,
		Age:        user.Age,
		Gender:     GenderToPB(user.Gender),
		Latitude:   user.Latitude,
		Longitude:  user.Longitude,
		Bio:        user.BIO,
		Goal:       common.GoalToPB(user.Goal),
		Interest:   interest.InterestToPB(user.Interest),
		Zodiac:     common.ZodiacToPB(user.Zodiac),
		Height:     user.Height,
		Education:  common.EducationToPB(user.Education),
		Children:   common.ChildrenToPB(user.Children),
		Alcohol:    common.AlcoholToPB(user.Alcohol),
		Smoking:    common.SmokingToPB(user.Smoking),
		IsHidden:   user.IsHidden,
		IsVerified: user.IsVerified,
		IsPremium:  user.IsPremium,
		IsBlocked:  user.IsBlocked,
		Photos: slices.Map(user.Photos, func(from Photo) *desc.Photo {
			return &desc.Photo{
				OrderNumber: from.OrderNumber,
				Url:         from.DownloadURL,
			}
		}),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
