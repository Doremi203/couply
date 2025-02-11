package domain

import (
	"github.com/Doremi203/Couply/backend/internal/domain/interest"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

type User struct {
	ID        int64              `db:"id"`
	Name      string             `db:"name"`
	Age       int32              `db:"age"`
	Gender    Gender             `db:"gender"`
	Location  string             `db:"location"`
	BIO       string             `db:"bio"`
	Goal      Goal               `db:"goal"`
	Interest  *interest.Interest `db:"interest"`
	Zodiac    Zodiac             `db:"zodiac"`
	Height    int32              `db:"height"`
	Education Education          `db:"education"`
	Children  Children           `db:"children"`
	Alcohol   Alcohol            `db:"alcohol"`
	Smoking   Smoking            `db:"smoking"`
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

func (b *UserBuilder) SetGoal(goal Goal) *UserBuilder {
	b.user.Goal = goal
	return b
}

func (b *UserBuilder) SetInterest(interest *interest.Interest) *UserBuilder {
	b.user.Interest = interest
	return b
}

func (b *UserBuilder) SetZodiac(zodiac Zodiac) *UserBuilder {
	b.user.Zodiac = zodiac
	return b
}

func (b *UserBuilder) SetHeight(height int32) *UserBuilder {
	b.user.Height = height
	return b
}

func (b *UserBuilder) SetEducation(education Education) *UserBuilder {
	b.user.Education = education
	return b
}

func (b *UserBuilder) SetChildren(children Children) *UserBuilder {
	b.user.Children = children
	return b
}

func (b *UserBuilder) SetAlcohol(alcohol Alcohol) *UserBuilder {
	b.user.Alcohol = alcohol
	return b
}

func (b *UserBuilder) SetSmoking(smoking Smoking) *UserBuilder {
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
		Id:        strconv.FormatInt(user.ID, 10),
		Name:      user.Name,
		Age:       user.Age,
		Gender:    GenderToPB(user.Gender),
		Location:  user.Location,
		Bio:       user.BIO,
		Goal:      GoalToPB(user.Goal),
		Interest:  interest.InterestToPB(user.Interest),
		Zodiac:    ZodiacToPB(user.Zodiac),
		Height:    user.Height,
		Education: EducationToPB(user.Education),
		Children:  ChildrenToPB(user.Children),
		Alcohol:   AlcoholToPB(user.Alcohol),
		Smoking:   SmokingToPB(user.Smoking),
		Hidden:    user.Hidden,
		Verified:  user.Verified,
		Photos:    PhotoSliceToPB(user.Photos),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func PBToUser(user *desc.User) User {
	id, _ := strconv.Atoi(user.Id)

	return User{
		ID:        int64(id),
		Name:      user.Name,
		Age:       user.Age,
		Gender:    PBToGender(user.Gender),
		Location:  user.Location,
		BIO:       user.Bio,
		Goal:      PBToGoal(user.Goal),
		Interest:  interest.PBToInterest(user.Interest),
		Zodiac:    PBToZodiac(user.Zodiac),
		Height:    user.Height,
		Education: PBToEducation(user.Education),
		Children:  PBToChildren(user.Children),
		Alcohol:   PBToAlcohol(user.Alcohol),
		Smoking:   PBToSmoking(user.Smoking),
		Hidden:    user.Hidden,
		Verified:  user.Verified,
		Photos:    PBToPhotoSlice(user.Photos),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
	}
}
