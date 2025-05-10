package user

import (
	"time"

	"github.com/Doremi203/couply/backend/common/libs/slices"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	Interest   *interest.Interest `db:"interest"`
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
	Photos     []Photo            `db:"photos"`
	CreatedAt  time.Time          `db:"created_at"`
	UpdatedAt  time.Time          `db:"updated_at"`
}

func (x *User) GetID() uuid.UUID {
	if x != nil {
		return x.ID
	}
	return uuid.Nil
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *User) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *User) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *User) GetBIO() string {
	if x != nil {
		return x.BIO
	}
	return ""
}

func (x *User) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return 0
}

func (x *User) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *User) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return 0
}

func (x *User) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *User) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return 0
}

func (x *User) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return 0
}

func (x *User) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return 0
}

func (x *User) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return 0
}

func (x *User) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *User) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *User) GetIsPremium() bool {
	if x != nil {
		return x.IsPremium
	}
	return false
}

func (x *User) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *User) GetPhotos() []Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *User) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}

func (x *User) GetUpdatedAt() time.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return time.Time{}
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
		Id:         user.GetID().String(),
		Name:       user.GetName(),
		Age:        user.GetAge(),
		Gender:     GenderToPB(user.GetGender()),
		Latitude:   user.GetLatitude(),
		Longitude:  user.GetLongitude(),
		Bio:        user.GetBIO(),
		Goal:       common.GoalToPB(user.GetGoal()),
		Interest:   interest.InterestToPB(user.GetInterest()),
		Zodiac:     common.ZodiacToPB(user.GetZodiac()),
		Height:     user.GetHeight(),
		Education:  common.EducationToPB(user.GetEducation()),
		Children:   common.ChildrenToPB(user.GetChildren()),
		Alcohol:    common.AlcoholToPB(user.GetAlcohol()),
		Smoking:    common.SmokingToPB(user.GetSmoking()),
		IsHidden:   user.GetIsHidden(),
		IsVerified: user.GetIsVerified(),
		IsPremium:  user.GetIsPremium(),
		IsBlocked:  user.GetIsBlocked(),
		Photos: slices.Map(user.GetPhotos(), func(from Photo) *desc.Photo {
			return &desc.Photo{
				OrderNumber: from.GetOrderNumber(),
				Url:         from.DownloadURL,
			}
		}),
		CreatedAt: timestamppb.New(user.GetCreatedAt()),
		UpdatedAt: timestamppb.New(user.GetUpdatedAt()),
	}
}

func UsersToPB(users []*User) []*desc.User {
	pbUsers := make([]*desc.User, 0, len(users))
	for _, u := range users {
		if u != nil {
			pbUsers = append(pbUsers, UserToPB(u))
		}
	}
	return pbUsers
}
