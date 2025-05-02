package user

import (
	"time"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	ID        int64              `db:"id"`
	Name      string             `db:"name"`
	Age       int32              `db:"age"`
	Gender    Gender             `db:"gender"`
	Location  string             `db:"location"`
	BIO       string             `db:"bio"`
	Goal      common.Goal        `db:"goal"`
	Interest  *interest.Interest `db:"interest"`
	Zodiac    common.Zodiac      `db:"zodiac"`
	Height    int32              `db:"height"`
	Education common.Education   `db:"education"`
	Children  common.Children    `db:"children"`
	Alcohol   common.Alcohol     `db:"alcohol"`
	Smoking   common.Smoking     `db:"smoking"`
	Hidden    bool               `db:"hidden"`
	Verified  bool               `db:"verified"`
	Photos    []*Photo           `db:"photos"`
	CreatedAt time.Time          `db:"created_at"`
	UpdatedAt time.Time          `db:"updated_at"`
}

type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: &User{}}
}

func (b *UserBuilder) SetID(id int64) *UserBuilder {
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

func (b *UserBuilder) SetLocation(location string) *UserBuilder {
	b.user.Location = location
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

func (b *UserBuilder) SetHidden(hidden bool) *UserBuilder {
	b.user.Hidden = hidden
	return b
}

func (b *UserBuilder) SetVerified(verified bool) *UserBuilder {
	b.user.Verified = verified
	return b
}

func (b *UserBuilder) SetPhotos(photos []*Photo) *UserBuilder {
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
		Id:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		Gender:    GenderToPB(user.Gender),
		Location:  user.Location,
		Bio:       user.BIO,
		Goal:      common.GoalToPB(user.Goal),
		Interest:  interest.InterestToPB(user.Interest),
		Zodiac:    common.ZodiacToPB(user.Zodiac),
		Height:    user.Height,
		Education: common.EducationToPB(user.Education),
		Children:  common.ChildrenToPB(user.Children),
		Alcohol:   common.AlcoholToPB(user.Alcohol),
		Smoking:   common.SmokingToPB(user.Smoking),
		Hidden:    user.Hidden,
		Verified:  user.Verified,
		Photos:    PhotoSliceToPB(user.Photos),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func PBToUser(user *desc.User) *User {
	return &User{
		ID:        user.GetId(),
		Name:      user.GetName(),
		Age:       user.GetAge(),
		Gender:    PBToGender(user.GetGender()),
		Location:  user.GetLocation(),
		BIO:       user.GetBio(),
		Goal:      common.PBToGoal(user.GetGoal()),
		Interest:  interest.PBToInterest(user.GetInterest()),
		Zodiac:    common.PBToZodiac(user.GetZodiac()),
		Height:    user.GetHeight(),
		Education: common.PBToEducation(user.GetEducation()),
		Children:  common.PBToChildren(user.GetChildren()),
		Alcohol:   common.PBToAlcohol(user.GetAlcohol()),
		Smoking:   common.PBToSmoking(user.GetSmoking()),
		Hidden:    user.GetHidden(),
		Verified:  user.GetVerified(),
		Photos:    PBToPhotoSlice(user.GetPhotos()),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: user.GetUpdatedAt().AsTime(),
	}
}
