package dao

import (
	"github.com/go-rel/rel"
)

const (
	ACTIVE  = 10
	DELETED = 0
	Status  = "status"
	ID      = "id"
)

// Models - queries of all models
type Models interface {
	ComicsDao() ComicsDao
	DonationDao() DonationDao
	UserDao() UserDao
	CompanyDao() CompanyDao
}

type models struct {
	comics    ComicsDao
	donations DonationDao
	users     UserDao
	companies CompanyDao
}

// NewModels - return new manager instance of all models
func NewModels(repository rel.Repository) *models {
	return &models{
		comics:    NewComicsDao(repository),
		donations: NewDonationDao(repository),
		users:     NewUserDao(repository),
		companies: NewCompanyDao(repository),
	}
}

// ComicsDao - дао комиксов
func (m *models) ComicsDao() ComicsDao {
	return m.comics
}

// DonationDao - дао пожертвований
func (m *models) DonationDao() DonationDao {
	return m.donations
}

// UserDao - дао юзеров
func (m *models) UserDao() UserDao {
	return m.users
}

// CompanyDao - дао компаний
func (m *models) CompanyDao() CompanyDao {
	return m.companies
}
