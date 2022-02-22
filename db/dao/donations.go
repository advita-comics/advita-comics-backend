package dao

import (
	"context"
	"time"

	"github.com/go-rel/rel"
)

// Donation -  модель пожертвования
type Donation struct {
	ID              int
	Amount          float64
	CompanyID       int
	DonationTypeID  int
	UserID          int
	ComicsID        int
	Personalisation string

	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Override table name to be `donation`
func (b Donation) Table() string {
	return "donation"
}

// DonationType -  модель варианта пожертвования
type DonationType struct {
	ID          int
	Name        string
	MinAmount   float64
	Description string

	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Override table name to be `donation_type`
func (b DonationType) Table() string {
	return "donation_type"
}

type donationDao struct {
	db rel.Repository
}

// DonationDao - методы для взаимодействия с таблицей пожертвований и их типов
type DonationDao interface {
	Create(ctx context.Context, donation *Donation) (*Donation, error)
	Donations(ctx context.Context, companyID int) ([]Donation, error)
	DonationSum(ctx context.Context, companyID int) (float64, error)
	ListDonationTypes(ctx context.Context, filter map[string]interface{}) ([]DonationType, error)
}

// NewDonationDao - конструктор
func NewDonationDao(repository rel.Repository) DonationDao {
	return donationDao{db: repository}
}

// Create - создает новое пожертвование
func (d donationDao) Create(ctx context.Context, donation *Donation) (*Donation, error) {
	d.db.MustInsert(ctx, donation, rel.Reload(true))
	return donation, nil
}

// DonationSum - считает сумму всех пожертвований на определенную компанию
func (d donationDao) DonationSum(ctx context.Context, companyID int) (float64, error) {
	var err error
	var sum struct {
		Sum float64
	}

	sql := rel.SQL(`SELECT sum(amount) as sum
								FROM donation
							WHERE status=? and company_id=?;`, ACTIVE, companyID)
	err = d.db.Find(ctx, &sum, sql)
	if err != nil {
		return sum.Sum, err
	}

	return sum.Sum, nil
}

func (d donationDao) Donations(ctx context.Context, companyID int) ([]Donation, error) {
	var err error
	var donations []Donation

	sql := rel.SQL(`SELECT id, amount, company_id, user_id, comics_id
								FROM donation
							WHERE status=? and company_id=?;`, ACTIVE, companyID)
	err = d.db.FindAll(ctx, &donations, sql)
	if err != nil {
		return donations, err
	}

	return donations, nil
}

// ListDonationTypes - отдает слайс вариантов пожертвований удовлетворяющих фильтру,
// сортирует по убыванию минимальной суммы
func (d donationDao) ListDonationTypes(ctx context.Context, filter map[string]interface{}) ([]DonationType, error) {
	var res []DonationType
	f := []rel.FilterQuery{
		rel.Eq(Status, ACTIVE),
	}

	for k, v := range filter {
		f = append(f, rel.Eq(k, v))
	}

	if err := d.db.FindAll(ctx, &res, rel.Where(f...), rel.SortQuery{
		Field: "min_amount",
		Sort:  -1,
	}); err != nil {
		if err == rel.ErrNotFound {
			return res, nil
		}
		return nil, err
	}

	return res, nil
}
