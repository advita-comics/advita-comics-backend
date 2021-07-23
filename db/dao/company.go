package dao

import (
	"context"
	"time"

	"github.com/go-rel/rel"
)

// Company -  модель компании
type Company struct {
	ID                int
	Name              string
	TerminationAmount float64
	ExpirationDate    time.Time

	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Override table name to be `company`
func (b Company) Table() string {
	return "company"
}

type companyDao struct {
	db rel.Repository
}

// CompanyDao - методы для взаимодействия с таблицей компаний
type CompanyDao interface {
	List(ctx context.Context, filter map[string]interface{}) ([]Company, error)
}

// NewCompanyDao - конструктор
func NewCompanyDao(repository rel.Repository) CompanyDao {
	return companyDao{db: repository}
}

// List - отдает слайс компаний удовлетворяющих фильтру
func (d companyDao) List(ctx context.Context, filter map[string]interface{}) ([]Company, error) {
	var res []Company
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
