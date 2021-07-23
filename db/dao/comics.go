package dao

import (
	"context"
	"time"

	"github.com/go-rel/rel"
)

// Comics -  модель комикса
type Comics struct {
	ID          int
	Name        string
	Description string
	Path        string

	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type comicsDao struct {
	db rel.Repository
}

// ComicsDao - методы для взаимодействия с таблицей комиксов
type ComicsDao interface {
	List(ctx context.Context, filter map[string]interface{}) ([]Comics, error)
}

// NewComicsDao - конструктор
func NewComicsDao(repository rel.Repository) ComicsDao {
	return comicsDao{db: repository}
}

// List - отдает слайс комиксов удовлетворяющих фильтру
func (d comicsDao) List(ctx context.Context, filter map[string]interface{}) ([]Comics, error) {
	var res []Comics
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
