package dao

import (
	"context"
	"time"

	"github.com/go-rel/rel"
)

const (
	Email = "email"
	USER  = "USER"
)

// User -  модель пользователя
type User struct {
	ID            int
	Name          string
	Email         string
	Role          string
	GetReport     bool
	FollowProcess bool

	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Override table name to be `user`
func (b User) Table() string {
	return "user"
}

type userDao struct {
	db rel.Repository
}

// UserDao - методы для взаимодействия с таблицей пользователей
type UserDao interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	List(ctx context.Context, filter map[string]interface{}) ([]User, error)
}

// NewUserDao - конструктор
func NewUserDao(repository rel.Repository) UserDao {
	return userDao{db: repository}
}

// Create - создает нового пользователя
func (d userDao) Create(ctx context.Context, user *User) (*User, error) {
	d.db.MustInsert(ctx, user, rel.Reload(true))
	return user, nil
}

// Update - обновляет существующего пользователя
func (d userDao) Update(ctx context.Context, user *User) (*User, error) {
	d.db.MustUpdate(ctx, user, rel.Reload(true))
	return user, nil
}

// List - отдает слайс комиксов удовлетворяющих фильтру
func (d userDao) List(ctx context.Context, filter map[string]interface{}) ([]User, error) {
	var res []User
	f := []rel.FilterQuery{
		rel.Eq(Status, ACTIVE),
	}

	for k, v := range filter {
		f = append(f, rel.Eq(k, v))
	}

	if err := d.db.FindAll(ctx, &res, rel.Where(f...)); err != nil {
		if err == rel.ErrNotFound {
			return res, nil
		}
		return nil, err
	}

	return res, nil
}
