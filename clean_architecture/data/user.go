package data

import (
	"context"
	"time"

	"github.com/adrian-lin-1-0-0/go-study/clean_architecture/biz"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Card struct {
	ID        int64
	UserID    int64
	Money     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

var _ biz.UserRepo = (*userRepo)(nil)
var _ biz.CardRepo = (*cardRepo)(nil)

type userRepo struct {
	data *Data
}

type cardRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, m *biz.User) (int64, error) {

	user := User{Name: m.Name, Email: m.Email}
	result := u.data.DB(ctx).Create(&user)
	return user.ID, result.Error
}

func NewCardRepo(data *Data) biz.CardRepo {
	return &cardRepo{
		data: data,
	}
}

func (c *cardRepo) CreateCard(ctx context.Context, id int64) (int64, error) {
	var card Card
	card.UserID = id
	card.Money = 1000
	result := c.data.DB(ctx).Save(&card)
	return card.ID, result.Error
}
